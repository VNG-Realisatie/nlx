package main

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/derrsa"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/insight-api-irma/irma"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	logoptions.LogOptions

	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:80" description:"Adress for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx_logdb?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	IRMAJWTRSAVerifyPublicKeyDER string `long:"irma-jwt-rsa-verify-public-key-der" env:"IRMA_JWT_RSA_VERIFY_PUBLIC_KEY_DER" required:"true" description:"PEM RSA public key to verify results from irma api server"`
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
	config := options.LogOptions.ZapConfig()
	logger, err := config.Build()
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

	db, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}
	db.MapperFunc(xstrings.ToSnakeCase)

	dbversion.WaitUntilLatestTxlogDBVersion(logger, db.DB)

	rsaVerifyPublicKey, err := derrsa.DecodeDEREncodedRSAPublicKey(bytes.NewBufferString(options.IRMAJWTRSAVerifyPublicKeyDER))
	if err != nil {
		log.Fatalf("failed to parse rsa private key: %v", err)
	}

	http.HandleFunc("/fetch", newTxlogFetcher(logger, db, rsaVerifyPublicKey))
	err = http.ListenAndServe(options.ListenAddress, nil)
	if err != nil {
		logger.Fatal("failed to ListenAndServe", zap.Error(err))
	}
}

func newTxlogFetcher(logger *zap.Logger, db *sqlx.DB, rsaVerifyPublicKey *rsa.PublicKey) http.HandlerFunc {
	stmtFetchLogs, err := db.Preparex(`
		SELECT
			created,
			src_organization,
			dest_organization,
			service_name,
			logrecord_id,
			data AS data_json
		FROM transactionlog.records
		ORDER BY created
	`)
	if err != nil {
		logger.Fatal("failed to prepare query for fetching logs", zap.Error(err))
	}

	type Record struct {
		*transactionlog.Record
		Created  time.Time      `json:"created"`
		DataJSON types.JSONText `json:"-"`
	}

	type Out struct {
		Records []*Record `json:"records"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		jwtBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			logger.Error("could not read http request body", zap.Error(err))
			http.Error(w, "could not read http request body", http.StatusBadRequest)
			return
		}
		token, claims, err := irma.VerifyIRMAVerificationResult(jwtBytes, rsaVerifyPublicKey)
		_ = claims
		_ = token
		if err != nil {
			logger.Error("failed to verify irma jwt", zap.Error(err))
			http.Error(w, "invalid irma jwt", http.StatusBadRequest)
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

		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			logger.Error("failed to output records", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return

		}
	}
}
