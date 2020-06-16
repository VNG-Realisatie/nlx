// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/pkg/directory"
)

func directoryRoutes(a *API) chi.Router {
	r := chi.NewRouter()
	r.Get("/services", a.directoryServicesHandler)

	return r
}

type DirectoryService struct {
	ServiceName          string                 `json:"serviceName"`
	OrganizationName     string                 `json:"organizationName"`
	APISpecificationType string                 `json:"apiSpecificationType"`
	Status               DirectoryServiceStatus `json:"status"`
}

type DirectoryServiceStatus string

const (
	DirectoryServiceStatusUnknown  DirectoryServiceStatus = "unknown"
	DirectoryServiceStatusUp       DirectoryServiceStatus = "up"
	DirectoryServiceStatusDown     DirectoryServiceStatus = "down"
	DirectoryServiceStatusDegraded DirectoryServiceStatus = "degraded"
)

var inwayStateToDirectoryStatus = map[directory.InwayState]DirectoryServiceStatus{
	directory.InwayStateUnknown: DirectoryServiceStatusUnknown,
	directory.InwayStateUp:      DirectoryServiceStatusUp,
	directory.InwayStateDown:    DirectoryServiceStatusDown,
}

func (a *API) directoryServicesHandler(w http.ResponseWriter, r *http.Request) {
	services, err := a.directoryClient.ListServices()

	if err != nil {
		a.logger.Error("fetching services from directory", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	filtered := []*DirectoryService{}

	for _, s := range services {
		if s.OrganizationName == a.environment.OrganizationName {
			continue
		}

		status := DetermineDirectoryServiceStatus(s.Inways)

		ds := &DirectoryService{
			ServiceName:          s.Name,
			OrganizationName:     s.OrganizationName,
			APISpecificationType: s.APISpecificationType,
			Status:               status,
		}

		filtered = append(filtered, ds)
	}

	render.JSON(w, r, filtered)
}

func DetermineDirectoryServiceStatus(inways []*directory.Inway) DirectoryServiceStatus {
	status := DirectoryServiceStatusUnknown

	if len(inways) == 0 {
		return status
	}

	stateMap := map[directory.InwayState]int{}

	for _, i := range inways {
		stateMap[i.State]++
	}

	if len(stateMap) > 1 {
		return DirectoryServiceStatusDegraded
	}

	for state := range stateMap {
		status = inwayStateToDirectoryStatus[state]
	}

	return status
}
