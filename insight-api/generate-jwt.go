// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package insightapi

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"

	"go.nlx.io/nlx/insight-api/irma"
)

type GenerateJWTRequest struct {
	DataSubjects []string `json:"dataSubjects"`
}

func (i *InsightAPI) generateJWT(serviceProviderName string, rsaSignPrivateKey *rsa.PrivateKey) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestedDataSubjects := &GenerateJWTRequest{}
		err := json.NewDecoder(r.Body).Decode(requestedDataSubjects)
		defer r.Body.Close()
		if err != nil {
			i.logger.Error("failed to decode requested data subjects", zap.Error(err))
			http.Error(w, "incorrect request data", http.StatusBadRequest)
			return
		}

		discloseRequest := irma.DiscloseRequest{
			Content: []irma.DiscloseRequestContent{},
		}
		for _, k := range requestedDataSubjects.DataSubjects {
			v, ok := i.irmaAttributes[k]
			if !ok {
				i.logger.Error("unknown dataSubject")
				http.Error(w, "incorrect dataSubject requested", http.StatusBadRequest)
				return
			}

			currentDiscloseContent := irma.DiscloseRequestContent{
				Label:      v.Label,
				Attributes: v.IrmaAttributes,
			}
			discloseRequest.Content = append(discloseRequest.Content, currentDiscloseContent)
		}

		signedJWT, err := i.irmaHandler.GenerateAndSignJWT(&discloseRequest, serviceProviderName, rsaSignPrivateKey)
		if err != nil {
			i.logger.Error("failed to generate JWT", zap.Error(err))
			http.Error(w, "failed to generate JWT", http.StatusInternalServerError)
			return
		}

		render.PlainText(w, r, signedJWT)
	}
}
