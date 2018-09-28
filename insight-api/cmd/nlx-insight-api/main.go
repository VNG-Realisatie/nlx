package main

import (
	"bytes"
	"crypto/rsa"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/derrsa"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/insight-api/config"
	"go.nlx.io/nlx/insight-api/irma"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	logoptions.LogOptions

	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:80" description:"Adress for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx_logdb?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	IRMAJWTRSASignPrivateKeyDER  string `long:"irma-jwt-rsa-sign-private-key-der" env:"IRMA_JWT_RSA_SIGN_PRIVATE_KEY_DER" required:"true" description:"PEM RSA private key to sign requests for irma api server"`
	IRMAJWTRSAVerifyPublicKeyDER string `long:"irma-jwt-rsa-verify-public-key-der" env:"IRMA_JWT_RSA_VERIFY_PUBLIC_KEY_DER" required:"true" description:"PEM RSA public key to verify results from irma api server"`

	ServiceConfig string `long:"service-config" env:"SERVICE_CONFIG" default:"service-config.toml" description:"Location of the service config toml file"`
}

func main() {

	// Parse options
	args, err := flags.Parse(&options)
	if err != nil {
		if et, ok := err.(*flags.Error); ok {
			if et.Type == flags.ErrHelp {
				return
			}
		}
		log.Fatalf("error parsing flags: %v", err)
	}
	if len(args) > 0 {
		log.Fatalf("unexpected arguments: %v", args)
	}

	// Setup new zap logger
	zapConfig := options.LogOptions.ZapConfig()
	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	defer func() { // TODO(GeertJohan): #205 make this a common/process exitFunc?
		syncErr := logger.Sync()
		if syncErr != nil {
			// notify the user that proper logging has failed
			fmt.Fprintf(os.Stderr, "failed to sync zap logger: %v\n", syncErr)
			// don't exit when we're in a panic
			if p := recover(); p != nil {
				panic(p)
			}
			os.Exit(1)
		}
	}()

	process.Setup(logger)

	serviceConfig := config.LoadServiceConfig(logger, options.ServiceConfig)

	db, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}
	db.MapperFunc(xstrings.ToSnakeCase)

	dbversion.WaitUntilLatestTxlogDBVersion(logger, db.DB)

	rsaSignPrivateKey, err := derrsa.DecodeDEREncodedRSAPrivateKey(bytes.NewBufferString(options.IRMAJWTRSASignPrivateKeyDER))
	if err != nil {
		logger.Fatal("failed to parse rsa private key", zap.Error(err))
	}
	rsaVerifyPublicKey, err := derrsa.DecodeDEREncodedRSAPublicKey(bytes.NewBufferString(options.IRMAJWTRSAVerifyPublicKeyDER))
	if err != nil {
		logger.Fatal("failed to parse rsa public key", zap.Error(err))
	}
	r := chi.NewRouter()
	r.Use(HappyOptionsHandler)
	r.Get("/getDataSubjects", listDataSubjects(logger, serviceConfig.DataSubjects))
	r.Post("/generateJWT", generateJWT(logger, serviceConfig.DataSubjects, "insight", rsaSignPrivateKey))
	r.Post("/fetch", newTxlogFetcher(logger, db, rsaVerifyPublicKey))
	err = http.ListenAndServe(options.ListenAddress, r)
	if err != nil {
		logger.Fatal("failed to ListenAndServe", zap.Error(err))
	}
}

func listDataSubjects(logger *zap.Logger, dataSubjects map[string]config.DataSubject) http.HandlerFunc {
	type JSONDataSubject struct {
		Label string `json:"label"`
	}

	type JSONResult struct {
		DataSubjects map[string]JSONDataSubject `json:"dataSubjects"`
	}

	outputList := JSONResult{
		DataSubjects: map[string]JSONDataSubject{},
	}
	for k, v := range dataSubjects {
		outputList.DataSubjects[k] = JSONDataSubject{
			Label: v.Label,
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(outputList)
		if err != nil {
			logger.Error("failed to output DataSubjects", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}
}

func generateJWT(logger *zap.Logger, dataSubjects map[string]config.DataSubject, serviceProviderName string, rsaSignPrivateKey *rsa.PrivateKey) http.HandlerFunc {
	type JSONRequest struct {
		DataSubjects []string `json:"dataSubjects"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}

		requestedDataSubjects := &JSONRequest{}
		err := json.NewDecoder(r.Body).Decode(requestedDataSubjects)
		defer r.Body.Close()
		if err != nil {
			logger.Error("failed to decode requested data subjects", zap.Error(err))
			//http.Error(w, "incorrect request data", http.StatusBadRequest)
			return
		}

		discloseRequest := irma.DiscloseRequest{
			Content: []irma.DiscloseRequestContent{},
		}
		for _, k := range requestedDataSubjects.DataSubjects {
			v, ok := dataSubjects[k]
			if !ok {
				logger.Error("unknown dataSubject")
				http.Error(w, "incorrect request data", http.StatusBadRequest)
				return
			}
			currentDiscloseContent := irma.DiscloseRequestContent{
				Label:      v.Label,
				Attributes: v.IrmaAttributes,
			}
			discloseRequest.Content = append(discloseRequest.Content, currentDiscloseContent)
		}

		signedJWT, err := irma.GenerateAndSignJWT(&discloseRequest, serviceProviderName, rsaSignPrivateKey)
		if err != nil {
			logger.Error("failed to generate JWT", zap.Error(err))
			http.Error(w, "failed to generate JWT", http.StatusInternalServerError)
			return
		}

		_, err = w.Write([]byte(signedJWT))

		if err != nil {
			logger.Error("failed to output JWT", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}
}

func newTxlogFetcher(logger *zap.Logger, db *sqlx.DB, rsaVerifyPublicKey *rsa.PublicKey) http.HandlerFunc {
	stmtCreateMatchDataSubjects, err := db.Preparex(`
		CREATE TEMPORARY TABLE matchDataSubjects(
			key varchar(100),
			value varchar(100)
		)
		ON COMMIT DROP
	`)
	if err != nil {
		logger.Fatal("failed to prepare stmtCreateMatchDataSubjects", zap.Error(err))
	}

	rawStmtInsertMatchDataSubjects := `INSERT INTO matchDataSubjects (key, value) VALUES ($1, $2)`

	rawStmtFetchLogs := `
		WITH matchedRecords AS (
			SELECT DISTINCT record_id
			FROM transactionlog.datasubjects
				INNER JOIN matchDataSubjects
					ON datasubjects.key = matchDataSubjects.key
						AND datasubjects.value = matchDataSubjects.value
		)
		SELECT
			created,
			src_organization,
			dest_organization,
			service_name,
			logrecord_id,
			data AS data_json
		FROM transactionlog.records
			INNER JOIN matchedRecords
				ON records.id = matchedRecords.record_id
		ORDER BY created
	`

	type Record struct {
		*transactionlog.Record
		Created  time.Time      `json:"created"`
		DataJSON types.JSONText `json:"-"`
	}

	type Out struct {
		Records []*Record `json:"records"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			logger.Error("could not read http request body", zap.Error(err))
			http.Error(w, "could not read http request body", http.StatusBadRequest)
			return
		}
		/*token, claims, err := irma.VerifyIRMAVerificationResult(jwtBytes, rsaVerifyPublicKey)
		_ = token
		if err != nil {
			logger.Error("failed to verify irma jwt", zap.Error(err))
			http.Error(w, "invalid irma jwt", http.StatusBadRequest)
			return
		}*/

		tx, err := db.Beginx()
		if err != nil {
			logger.Error("failed to start transaction", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		defer func() {
			errRollback := tx.Rollback()
			if errRollback == sql.ErrTxDone {
				return // tx was already comitted
			}
			if errRollback != nil {
				logger.Error("error rolling back transaction", zap.Error(errRollback))
			}
		}()

		// TODO: investigate possible other solutions or optimizations for this crazy exercise
		_, err = tx.Stmtx(stmtCreateMatchDataSubjects).Exec()
		if err != nil {
			logger.Error("failed to create temp table", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		stmtInsertMatchDataSubjects, err := tx.Preparex(rawStmtInsertMatchDataSubjects)
		if err != nil {
			logger.Error("failed to prepare statement inserting query attributes into temp table", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		attributes := map[string]string{"burgerservicenummer": "194837487"}

		for key, value := range attributes {
			_, err = stmtInsertMatchDataSubjects.Exec(key, value)
			if err != nil {
				logger.Error("failed to insert query attributes into temp table", zap.Error(err))
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
		}

		stmtFetchLogs, err := tx.Preparex(rawStmtFetchLogs)
		if err != nil {
			logger.Error("failed to prepare statement for fetching transaction logs", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		res, err := stmtFetchLogs.Queryx()
		if err != nil {
			logger.Error("failed to fetch transaction logs", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		var out = Out{
			Records: make([]*Record, 0),
		}
		for res.Next() {
			rec := &Record{}
			err = res.StructScan(rec)
			if err != nil {
				logger.Error("failed to scan transaction log into struct", zap.Error(err))
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			err = rec.DataJSON.Unmarshal(&rec.Record.Data)
			if err != nil {
				logger.Error("failed to unmarshal record data fields", zap.Error(err))
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			out.Records = append(out.Records, rec)
		}

		err = tx.Commit()
		if err != nil {
			logger.Warn("failed to commit transaction after succesful select", zap.Error(err))
			// gracefully handle this situation, we can still serve the consumer.
		}

		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			logger.Error("failed to output records", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return

		}
	}
}
