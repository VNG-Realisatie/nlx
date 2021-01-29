// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jessevdk/go-flags"

	"go.nlx.io/nlx/common/logoptions"
)

var options struct {
	ListenAddress    string `long:"listen-address" env:"LISTEN_ADDRESS" default:"127.0.0.1:8080" description:"Address for the UI to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ManagementAPIURL string `long:"management-api-url" env:"MANAGEMENT_API_URL" required:"true" description:"URL to the Management API"`

	logoptions.LogOptions
}

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

	log.Printf("starting http server on %s", options.ListenAddress)

	managementAPIURL, err := url.Parse(options.ManagementAPIURL)
	if err != nil {
		log.Fatalf("invalid format for Management API URL: %s", err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(managementAPIURL)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/*", newSpaHandler("public",
		"index.html",
	))
	r.Handle("/api/*", proxy)
	r.Handle("/oidc/*", proxy)

	log.Fatal(http.ListenAndServe(options.ListenAddress, r))
}

// spaHandler implements the http.Handler interface.
// the implementation prevents users from getting a 404 when refreshing our SPA(Single Page Application)
// The SPA does client side routing which will not be directly available after refreshing the page
// To fix this spaHandler inspects the URL path to locate a file within the static dir
// If a file is found, it will be served. If not, index.html is served
type spaHandler struct {
	staticPath        string
	indexPath         string
	staticFileHandler http.Handler
}

func newSpaHandler(staticPath, indexPath string) *spaHandler {
	s := &spaHandler{
		indexPath:  indexPath,
		staticPath: staticPath,
	}

	s.staticFileHandler = http.FileServer(http.Dir(s.staticPath))

	return s
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Permissions-Policy", "accelerometer=(), ambient-light-sensor=(), autoplay=(), battery=(), camera=(), display-capture=(), document-domain=(), encrypted-media=(), execution-while-not-rendered=(), execution-while-out-of-viewport=(), fullscreen=(), geolocation=(), gyroscope=(), layout-animations=(), legacy-image-formats=(), magnetometer=(), microphone=(), midi=(), navigation-override=(), oversized-images=(), payment=(), picture-in-picture=(), publickey-credentials-get=(), sync-xhr=(), usb=(), vr=(), wake-lock=(), screen-wake-lock=(), web-share=(), xr-spatial-tracking=()")

	fileExists, err := doesFileExistInStaticFolder(r.URL.Path, h.staticPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !fileExists {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	h.staticFileHandler.ServeHTTP(w, r)
}

func doesFileExistInStaticFolder(urlPath, staticFilesPath string) (bool, error) {
	path, err := filepath.Abs(urlPath)
	if err != nil {
		return false, err
	}

	path = filepath.Join(staticFilesPath, path)
	_, err = os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
