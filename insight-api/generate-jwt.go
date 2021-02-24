// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package insightapi

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"

	"github.com/form3tech-oss/jwt-go"
	irma "github.com/privacybydesign/irmago"
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

		discloseRequest := irma.NewDisclosureRequest()

		for _, k := range requestedDataSubjects.DataSubjects {
			v, ok := i.irmaAttributes[k]
			if !ok {
				i.logger.Error("unknown dataSubject")
				http.Error(w, "incorrect dataSubject requested", http.StatusBadRequest)

				return
			}

			id := irma.NewAttributeTypeIdentifier(string(v.IrmaAttributes[0]))
			label := irma.TranslatedString{
				"nl": v.Label,
			}

			discloseRequest.AddSingle(id, nil, label)
		}

		sp := irma.NewServiceProviderJwt(serviceProviderName, discloseRequest)

		signedJWT, err := sp.Sign(jwt.SigningMethodRS256, rsaSignPrivateKey)
		if err != nil {
			i.logger.Error("failed to generate JWT", zap.Error(err))
			http.Error(w, "failed to generate JWT", http.StatusInternalServerError)

			return
		}

		render.PlainText(w, r, signedJWT)
	}
}
