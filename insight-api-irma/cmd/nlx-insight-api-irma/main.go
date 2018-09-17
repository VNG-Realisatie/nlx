package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	flags "github.com/jessevdk/go-flags"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/insight-api-irma/irma"
)

var options struct {
	logoptions.LogOptions

	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:80" description:"Adress for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	IRMACredentials              []string `long:"irma-credentials" env:"IRMA_CREDENTIALS" description:"List of IRMA credentails that may be validated"`
	IRMAEndpointURL              string   `long:"irma-endpoint-url" env:"IRMA_ENDPOINT_URL" required:"true" description:"URL for the IRMA api server (without the path /api/v2/... etc)"`
	IRMAJWTRSASignPrivateKeyDER  string   `long:"irma-jwt-rsa-sign-private-key-der" env:"IRMA_JWT_RSA_SIGN_PRIVATE_KEY_DER" required:"true" description:"PEM RSA private key to sign requests for irma api server"`
	IRMAJWTRSAVerifyPublicKeyDER string   `long:"irma-jwt-rsa-verify-public-key-der" env:"IRMA_JWT_RSA_VERIFY_PUBLIC_KEY_DER" required:"true" description:"PEM RSA public key to verify results from irma api server"`
}

var irmaClient *irma.Client

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
	rsaSignPrivateKeyDecoder := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(options.IRMAJWTRSASignPrivateKeyDER))
	rsaSignPrivateKeyBytes, err := ioutil.ReadAll(rsaSignPrivateKeyDecoder)
	if err != nil {
		log.Fatalf("failed to decode all bytes from rsa private key string: %v", err)
	}
	rsaSignPrivateKey, err := decodeDEREncodedRSAPrivateKey(rsaSignPrivateKeyBytes)
	if err != nil {
		log.Fatalf("failed to parse rsa private key: %v", err)
	}
	rsaVerifyPubkeyDecoder := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(options.IRMAJWTRSAVerifyPublicKeyDER))
	rsaVerifyPubkeyBytes, err := ioutil.ReadAll(rsaVerifyPubkeyDecoder)
	if err != nil {
		log.Fatalf("failed to decode all bytes from rsa private key string: %v", err)
	}
	rsaVerifyPubkey, err := decodeDEREncodedRSAPublicKey(rsaVerifyPubkeyBytes)
	if err != nil {
		log.Fatalf("failed to parse rsa private key: %v", err)
	}
	irmaClient, err = irma.NewClient(logger, "insight", options.IRMAEndpointURL, rsaSignPrivateKey, rsaVerifyPubkey)

	http.HandleFunc("/verification-start", newVerificationStart(logger, options.IRMACredentials))
	http.HandleFunc("/verification-poll", newVerificationPoll(logger))
	err = http.ListenAndServe(options.ListenAddress, nil)
	if err != nil {
		logger.Fatal("failed to ListenAndServe", zap.Error(err))
	}
}

func newVerificationStart(logger *zap.Logger, irmaCredentials []string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		discloseRequest := irma.DiscloseRequest{
			Content: []irma.DiscloseRequestContent{
				irma.DiscloseRequestContent{
					Label: "Over 18",
					Attributes: []irma.Attribute{
						irma.Attribute("irma-demo.MijnOverheid.ageLower.over18"),
					},
				},
			},
		}
		verificationDiscloseSession, err := irmaClient.StartVerification(&discloseRequest)
		if err != nil {
			logger.Error("failed to send verification request to irma api server", zap.Error(err))
			http.Error(w, "failed to send verification request to irma api server", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(verificationDiscloseSession)
		if err != nil {
			logger.Error("failed to output records", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return

		}
	}
}

func newVerificationPoll(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		verificationID := r.FormValue("verificationID")
		if verificationID == "" {
			logger.Warn("missing verificationID")
			http.Error(w, "missing verificationID", http.StatusBadRequest)
			return
		}

		verificationResultJWT, verificationResultClaims, err := irmaClient.PollVerification(verificationID)
		_, err = io.WriteString(w, verificationResultJWT)
		if err != nil {
			logger.Error("failed to write JWT", zap.Error(err))
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		_ = verificationResultClaims

		io.WriteString(w, verificationResultJWT)
	}
}

func decodeDEREncodedRSAPrivateKey(bts []byte) (*rsa.PrivateKey, error) {
	var err error
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(bts); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(bts); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, errors.New("key is not a private rsa key")
	}

	return pkey, nil
}

func decodeDEREncodedRSAPublicKey(bts []byte) (*rsa.PublicKey, error) {
	var err error
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PublicKey(bts); err != nil {
		if parsedKey, err = x509.ParsePKIXPublicKey(bts); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, errors.New("key is not a private rsa key")
	}

	return pkey, nil
}
