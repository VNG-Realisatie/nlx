// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package insightapi

import (
	"net/http"

	"github.com/go-chi/render"
)

type GetDataSubjectsResponse struct {
	DataSubjects map[string]DataSubject `json:"dataSubjects"`
}

type DataSubject struct {
	Label string `json:"label"`
}

func (i *InsightAPI) listDataSubjects() http.HandlerFunc {
	response := GetDataSubjectsResponse{
		DataSubjects: map[string]DataSubject{},
	}
	for k, v := range i.irmaAttributes {
		response.DataSubjects[k] = DataSubject{
			Label: v.Label,
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, response)
	}
}
