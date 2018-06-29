// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package certportal

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/sign"
	"github.com/cloudflare/cfssl/signer"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

// ListenAndServe is a blocking function that listens on provided tcp address to handle requests.
func (cp *CertPortal) ListenAndServe(address string) error {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Post("/request_certificate", requestCertificate(cp.caHost))
	})

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "files")
	FileServer(r, "/", http.Dir(filesDir))

	err := http.ListenAndServe(address, r)
	if err != nil {
		return errors.Wrap(err, "failed to run http server")
	}
	return nil
}

func requestCertificate(caHost string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &certificateRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Status(r, http.StatusBadRequest)
			log.Fatal(err)
			return
		}

		signReq := signer.SignRequest{
			Request: data.Csr,
		}

		s, err := sign.SignerFromConfig(cli.Config{
			Remote: caHost,
		})

		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		cert, err := s.Sign(signReq)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		render.Status(r, http.StatusCreated)
		render.SetContentType(render.ContentTypeJSON)
		render.Render(w, r, &certificateResponse{
			Certificate: string(cert),
		})
	}
}

type certificateRequest struct {
	Csr string `json:"csr"`
}

func (a *certificateRequest) Bind(r *http.Request) error {
	return nil
}

type certificateResponse struct {
	Certificate string `json:"certificate"`
}

func (rd *certificateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
