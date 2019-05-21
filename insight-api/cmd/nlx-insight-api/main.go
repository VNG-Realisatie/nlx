// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"bytes"
	"context"
	"crypto/rsa"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/derrsa"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/insight-api/config"
	"go.nlx.io/nlx/insight-api/irma"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	logoptions.LogOptions

	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8080" description:"Adress for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx_logdb?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	IRMAJWTRSASignPrivateKeyDER  string `long:"irma-jwt-rsa-sign-private-key-der" env:"IRMA_JWT_RSA_SIGN_PRIVATE_KEY_DER" required:"true" description:"PEM RSA private key to sign requests for irma api server"`
	IRMAJWTRSAVerifyPublicKeyDER string `long:"irma-jwt-rsa-verify-public-key-der" env:"IRMA_JWT_RSA_VERIFY_PUBLIC_KEY_DER" required:"true" description:"PEM RSA public key to verify results from irma api server"`

	InsightConfig string `long:"insight-config" env:"INSIGHT_CONFIG" default:"insight-config.toml" description:"Location of the insight config toml file"`

	CertFile string `long:"tls-cert" env:"TLS_CERT" description:"Absolute or relative path to the cert .pem"`
	KeyFile  string `long:"tls-key" env:"TLS_KEY" description:"Absolute or relative path to the key .pem"`
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
	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	process := process.NewProcess(logger)

	insightConfig := config.LoadInsightConfig(logger, options.InsightConfig)

	db, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}
	db.MapperFunc(xstrings.ToSnakeCase)
	process.CloseGracefully(db.Close)

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
	r.Get("/getDataSubjects", listDataSubjects(logger, insightConfig.DataSubjects))
	r.Post("/generateJWT", generateJWT(logger, insightConfig.DataSubjects, "insight", rsaSignPrivateKey))
	r.Post("/fetch", newTxlogFetcher(logger, db, insightConfig.DataSubjects, rsaVerifyPublicKey))

	server := &http.Server{
		Addr:    options.ListenAddress,
		Handler: r,
	}

	process.CloseGracefully(func() error {
		// Context with timeout to terminate server if shutdown operation takes longer than minute
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel() // do not remove. Otherwise it could cause implicit goroutine leak
		return server.Shutdown(localCtx)

	})

	if len(options.CertFile) > 0 {
		err = http.ListenAndServeTLS(options.ListenAddress, options.CertFile, options.KeyFile, r)
	} else {
		err = http.ListenAndServe(options.ListenAddress, r)
	}

	if err != nil {
		if err != http.ErrServerClosed {
			logger.Error("failed to ListenAndServe", zap.Error(err))
			return
		}
	}

	// Listener will return immediately on Shutdown call.
	// So we need to wait until all open connections will be closed gracefully
	<-process.ShutdownComplete
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
			http.Error(w, "incorrect request data", http.StatusBadRequest)
			return
		}

		discloseRequest := irma.DiscloseRequest{
			Content: []irma.DiscloseRequestContent{},
		}
		for _, k := range requestedDataSubjects.DataSubjects {
			v, ok := dataSubjects[k]
			if !ok {
				logger.Error("unknown dataSubject")
				http.Error(w, "incorrect dataSubject requested", http.StatusBadRequest)
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

func newTxlogFetcher(logger *zap.Logger, db *sqlx.DB, dataSubjects map[string]config.DataSubject, rsaVerifyPublicKey *rsa.PublicKey) http.HandlerFunc {
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
		ORDER BY created DESC
		LIMIT CASE WHEN $1>0 THEN $1 ELSE NULL END
		OFFSET $2
	`

	rawStmtGetRowCount := `
	WITH matchedRecords AS (
		SELECT DISTINCT record_id
		FROM transactionlog.datasubjects
			INNER JOIN matchDataSubjects
				ON datasubjects.key = matchDataSubjects.key
					AND datasubjects.value = matchDataSubjects.value
	)
	SELECT
		COUNT(*)
	FROM transactionlog.records
		INNER JOIN matchedRecords
			ON records.id = matchedRecords.record_id
	`

	// map irma attributes to a list of datasubjects that can be accessed by it
	var dataSubjectsByIrmaAttribute = make(map[string][]string)
	for dataSubjectKey, dataSubjectProperties := range dataSubjects {
		for _, irmaAttribute := range dataSubjectProperties.IrmaAttributes {
			dataSubjectsByIrmaAttribute[string(irmaAttribute)] = append(dataSubjectsByIrmaAttribute[string(irmaAttribute)], dataSubjectKey)
		}
	}

	type getLogsRequest struct {
		Page        int `schema:"page"`
		RowsPerPage int `schema:"rowsPerPage"`
	}

	type Record struct {
		*transactionlog.Record
		Created  time.Time      `json:"created"`
		DataJSON types.JSONText `json:"-"`
	}

	type Out struct {
		Records     []*Record `json:"records"`
		Page        int       `json:"page"`
		RowsPerPage int       `json:"rowsPerPage"`
		RowCount    int       `json:"rowCount"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		jwtBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			logger.Error("could not read http request body", zap.Error(err))
			http.Error(w, "could not read http request body", http.StatusBadRequest)
			return
		}
		_, claims, err := irma.VerifyIRMAVerificationResult(jwtBytes, rsaVerifyPublicKey)
		if err != nil {
			logger.Error("failed to verify irma jwt", zap.Error(err))
			http.Error(w, "invalid irma jwt", http.StatusBadRequest)
			return
		}

		requestParams := &getLogsRequest{}
		err = schema.NewDecoder().Decode(requestParams, r.URL.Query())
		if err != nil {
			logger.Error("error parsing URL values", zap.Error(err))
			http.Error(w, "failed to parse URL values", http.StatusBadRequest)
			return
		}

		var tx *sqlx.Tx
		for retry := 0; retry < 3; retry++ {
			tx, err = db.Beginx()
			if err == nil {
				break
			}
		}
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

		// TODO: #345 investigate possible other solutions or optimizations for this crazy exercise
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
		defer stmtInsertMatchDataSubjects.Close()

		for irmaAttributeKey, irmaAttributeValue := range claims.Attributes {
			for _, dataSubjectKey := range dataSubjectsByIrmaAttribute[irmaAttributeKey] {
				_, err = stmtInsertMatchDataSubjects.Exec(dataSubjectKey, irmaAttributeValue)
				if err != nil {
					logger.Error("failed to insert query attributes into temp table", zap.Error(err))
					http.Error(w, "server error", http.StatusInternalServerError)
					return
				}
			}
		}

		stmtFetchLogs, err := tx.Preparex(rawStmtFetchLogs)
		if err != nil {
			logger.Error("failed to prepare statement for fetching transaction logs", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		defer stmtFetchLogs.Close()
		res, err := stmtFetchLogs.Queryx(requestParams.RowsPerPage, requestParams.RowsPerPage*requestParams.Page)
		if err != nil {
			logger.Error("failed to fetch transaction logs", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		var out = Out{
			Records:     make([]*Record, 0),
			RowsPerPage: requestParams.RowsPerPage,
			Page:        requestParams.Page,
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

		stmtGetRowCount, err := tx.Preparex(rawStmtGetRowCount)
		if err != nil {
			logger.Error("failed to prepare statement for fetching transaction log row count", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		resRowCount := stmtGetRowCount.QueryRowx()
		err = resRowCount.Scan(&out.RowCount)
		if err != nil {
			logger.Error("failed to fetch transaction log count", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
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
