// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package certportal

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/cloudflare/cfssl/info"
	"github.com/cloudflare/cfssl/signer"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// function type to enable mocking of the signer
type createSignerFunc func() (signer.Signer, error)

func requestCertificateHandler(logger *zap.Logger, createSigner createSignerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &certificateRequest{}
		if err := render.Bind(r, data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error("error reading request", zap.Error(err))
			return
		}

		signReq := signer.SignRequest{
			Request: data.Csr,
		}

		s, err := createSigner()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error("error creating cert signer", zap.Error(err))
			return
		}

		cert, err := s.Sign(signReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error("error signing request", zap.Error(err))
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

func rootCertHandler(logger *zap.Logger, createSigner createSignerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		signer, err := createSigner()
		if err != nil {
			logger.Error("error obtaining root.crt from cfssl root CA", zap.Error(err))
			http.Error(w, "failed to create new cfssl signer", http.StatusInternalServerError)
			return
		}
		resp, err := signer.Info(info.Req{})
		if err != nil {
			logger.Error("error obtaining root.crt from cfssl root CA", zap.Error(err))
			http.Error(w, "failed to obtain root.crt from cfssl root CA", http.StatusInternalServerError)
			return
		}
		io.WriteString(w, resp.Certificate)
	}
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) error {
	if strings.ContainsAny(path, "{}*") {
		return fmt.Errorf("FileServer does not permit URL parameters")
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

	return nil
}
