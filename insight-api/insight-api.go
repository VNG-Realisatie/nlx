// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package insightapi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"go.nlx.io/nlx/insight-api/config"
	"go.nlx.io/nlx/insight-api/irma"
)

type InsightAPI struct {
	logger                      *zap.Logger
	router                      *chi.Mux
	logFetcher                  InsightLogFetcher
	irmaHandler                 irma.JWTHandler
	irmaAttributes              map[string]config.DataSubject
	dataSubjectsByIrmaAttribute map[string][]string
}

func NewInsightAPI(logger *zap.Logger, insightConfig *config.InsightConfig, jwtHandler irma.JWTHandler, logFetcher InsightLogFetcher, rsaPrivateKeyFile, rsaPublicKeyFile string) (*InsightAPI, error) {
	rsaSignPrivateKey, err := parseRSAPrivateKeyFile(rsaPrivateKeyFile)
	if err != nil {
		return nil, fmt.Errorf("error decoding private key: %s", err)
	}
	rsaVerifyPublicKey, err := parseRSAPublicKeyFile(rsaPublicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("error decoding public key: %s", err)
	}
	insightAPI := &InsightAPI{
		logger:                      logger,
		irmaAttributes:              insightConfig.DataSubjects,
		irmaHandler:                 jwtHandler,
		logFetcher:                  logFetcher,
		dataSubjectsByIrmaAttribute: make(map[string][]string),
	}
	insightAPI.router = chi.NewRouter()
	insightAPI.router.Use(HappyOptionsHandler)
	insightAPI.router.Get("/getDataSubjects", insightAPI.listDataSubjects())
	insightAPI.router.Post("/generateJWT", insightAPI.generateJWT("insight", rsaSignPrivateKey))
	insightAPI.router.Post("/fetch", insightAPI.newTxlogFetcher(rsaVerifyPublicKey))

	// map irma attributes to a list of datasubjects that can be accessed by it
	for dataSubjectKey, dataSubjectProperties := range insightConfig.DataSubjects {
		for _, irmaAttribute := range dataSubjectProperties.IrmaAttributes {
			insightAPI.dataSubjectsByIrmaAttribute[string(irmaAttribute)] = append(insightAPI.dataSubjectsByIrmaAttribute[string(irmaAttribute)], dataSubjectKey)
		}
	}

	return insightAPI, nil
}

func (i *InsightAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.router.ServeHTTP(w, r)
}
