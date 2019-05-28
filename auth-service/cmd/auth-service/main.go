// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

type authRequest struct {
	Headers      http.Header `json:"headers"`
	Organization string      `json:"organization"`
	Service      string      `json:"service"`
}

type authResponse struct {
	Authorized bool        `json:"authorized"`
	Headers    http.Header `json:"headers"`
	Reason     string      `json:"reason,omitempty"`
}

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8443" description:"Adress for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	CVSFile       string `long:"csv-file" env:"CSV_FILE" description:"absolute path to csv file to expose" required:"true"`
	CertFile      string `long:"tls-cert" env:"TLS_CERT" description:"Absolute or relative path to the Organization cert .pem"`
	KeyFile       string `long:"tls-key" env:"TLS_KEY" description:"Absolute or relative path to the Organization key .pem"`
}

type user struct {
	ID    string `json:"userID"`
	Token string `json:"username"`
}

var users = make(map[string]*user)

func main() {
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

	users, err = loadCSVFile(options.CVSFile)
	if err != nil {
		log.Fatalf("error loading CVS file: %s", err)
	}

	http.HandleFunc("/auth", authenticateHandler)
	log.Printf("starting http server on %s", options.ListenAddress)
	log.Fatal(http.ListenAndServeTLS(options.ListenAddress, options.CertFile, options.KeyFile, nil))
}

func loadCSVFile(filePath string) (map[string]*user, error) {
	data, err := ioutil.ReadFile(options.CVSFile)
	if err != nil {
		return nil, err
	}
	u := make(map[string]*user)
	reader := csv.NewReader(bytes.NewBuffer(data))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		user := &user{
			ID:    line[0],
			Token: line[1],
		}
		u[user.Token] = user
	}

	return u, nil
}

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	authRequest := &authRequest{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(authRequest)
	if err != nil {
		log.Printf("error decoding request %s", err)
		http.Error(w, "error decoding request", http.StatusBadRequest)
		return
	}

	log.Printf("Received auth request organization '%s' service '%s' ", authRequest.Organization, authRequest.Service)
	token := parseToken(authRequest.Headers)
	_, exists := users[token]
	if !exists {
		log.Println("user not found")
		json.NewEncoder(w).Encode(&authResponse{
			Authorized: false,
			Reason:     "invalid credentials",
		})
		return
	}

	json.NewEncoder(w).Encode(&authResponse{
		Authorized: true,
	})
}

// Parses token from the Proxy-Authorization header. The header should be in format <type> <credentials>
func parseToken(h http.Header) string {
	authString := h.Get("Proxy-Authorization")
	if len(authString) == 0 {
		log.Println("empty authorization header")
		return ""
	}

	authValues := strings.Split(authString, " ")
	if len(authValues) != 2 {
		log.Println("invalid authorization header")
		return ""
	}

	authType := authValues[0]
	authCredentails := authValues[1]
	// In this example implementation we only support the Bearer authorization type.
	switch authType {
	case "Bearer":
		return authCredentails
	default:
		return ""
	}
}
