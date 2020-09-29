// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package http

import (
	"database/sql"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type apiSpecHandler struct {
	httpClient           *http.Client
	selectInwayStatement *sqlx.Stmt
	logger               *zap.Logger
}

func newAPISpecHandler(httpClient *http.Client, db *sqlx.DB, logger *zap.Logger) (*apiSpecHandler, error) {
	h := &apiSpecHandler{
		httpClient: httpClient,
		logger:     logger.With(zap.String("handler", "api-spec")),
	}

	var err error
	h.selectInwayStatement, err = db.Preparex(`
		SELECT
			i.address AS inway_address
		FROM directory.inways i
			INNER JOIN directory.availabilities a ON a.inway_id = i.id
		WHERE a.service_id = $1
		AND a.healthy = true
		LIMIT 1
	`)

	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtSelectServiceInway")
	}

	return h, nil
}

func (h *apiSpecHandler) handleFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		service, ok := ctx.Value(serviceKey).(Service)

		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var inwayAddress string
		err := h.selectInwayStatement.QueryRowx(service.ID).Scan(&inwayAddress)

		if err != nil {
			switch err {
			case sql.ErrNoRows:
				h.logger.Error("No inways available", zap.Error(err))
			default:
				h.logger.Error("Failed to query inway address", zap.Error(err))
			}

			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)

			return
		}

		resp, err := h.getInwayAPISpec(inwayAddress, service.Name)

		if err != nil {
			h.logger.Info("failed to read api spec doc from remote inway", zap.Error(err))
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.WriteHeader(resp.StatusCode)

		_, err = io.Copy(w, resp.Body)

		if err != nil {
			h.logger.Error("Failed to copy body from inway", zap.Error(err))
		}
	}
}

func (h *apiSpecHandler) getInwayAPISpec(inwayAddress, serviceName string) (*http.Response, error) {
	inwayURL := url.URL{
		Scheme: "https",
		Host:   inwayAddress,
		Path:   path.Join("/.nlx/api-spec-doc/", serviceName),
	}

	r, err := h.httpClient.Get(inwayURL.String())
	if err != nil {
		return nil, err
	}

	return r, nil
}
