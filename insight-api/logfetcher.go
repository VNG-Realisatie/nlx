// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package insightapi

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/render"

	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx/types"
	"go.nlx.io/nlx/common/transactionlog"
	"go.uber.org/zap"
)

type GetLogsRequest struct {
	Page        int `schema:"page"`
	RowsPerPage int `schema:"rowsPerPage"`
}

type Record struct {
	*transactionlog.Record
	Created  time.Time      `json:"created"`
	DataJSON types.JSONText `json:"-"`
}

type GetLogRecordsResponse struct {
	Records     []*Record `json:"records"`
	Page        int       `json:"page"`
	RowsPerPage int       `json:"rowsPerPage"`
	RowCount    int       `json:"rowCount"`
}

func (i *InsightAPI) newTxlogFetcher(rsaVerifyPublicKey *rsa.PublicKey) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			i.logger.Error("could not read http request body", zap.Error(err))
			http.Error(w, "could not read http request body", http.StatusBadRequest)
			return
		}

		_, claims, err := i.irmaHandler.VerifyIRMAVerificationResult(jwtBytes, rsaVerifyPublicKey)
		if err != nil {
			i.logger.Error("failed to verify irma jwt", zap.Error(err))
			http.Error(w, "invalid irma jwt", http.StatusBadRequest)
			return
		}

		requestParams := &GetLogsRequest{}
		err = schema.NewDecoder().Decode(requestParams, r.URL.Query())
		if err != nil {
			i.logger.Error("error parsing URL values", zap.Error(err))
			http.Error(w, "failed to parse URL values", http.StatusBadRequest)
			return
		}

		out, err := i.logFetcher.GetLogRecords(requestParams.RowsPerPage, requestParams.Page, i.dataSubjectsByIrmaAttribute, claims)
		if err != nil {
			i.logger.Error("error getting log records", zap.Error(err))
			http.Error(w, "failed to retrieve log records", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, out)
	}
}
