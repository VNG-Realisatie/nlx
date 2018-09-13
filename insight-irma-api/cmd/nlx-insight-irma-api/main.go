package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	flags "github.com/svent/go-flags"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
)

var options struct {
	logoptions.LogOptions
	ListenAddress   string   `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:80" description:"Adress for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	IRMACredentials []string `long:"irma-credentials" env:"IRMA_CREDENTIALS" description:"List of IRMA credentails that may be validated"`
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

	http.HandleFunc("/start-validation", newStartValidation(logger, options.IRMACredentials))
	http.HandleFunc("/poll-validation", newPollValidation(logger))
	err = http.ListenAndServe(options.ListenAddress, nil)
	if err != nil {
		logger.Fatal("failed to ListenAndServe", zap.Error(err))
	}
}

func newStartValidation(logger *zap.Logger, irmaCredentials []string) http.HandlerFunc {

	type Out struct {
		ValidationID        string `json:"validation_id"`
		ValidationChallenge string `json:"validation_challenge"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		out := &Out{}
		err := json.NewEncoder(w).Encode(out)
		if err != nil {
			logger.Error("failed to output records", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return

		}
	}
}

func newPollValidation(logger *zap.Logger) http.HandlerFunc {

	type Out struct {
		ValidatedJWT string `json:"validated_jwt"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		out := &Out{}
		err := json.NewEncoder(w).Encode(out)
		if err != nil {
			logger.Error("failed to output records", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return

		}
	}
}
