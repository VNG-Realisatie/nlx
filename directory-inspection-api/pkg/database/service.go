// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
)

type Service struct {
	Name                  string                        `json:"name,omitempty"`
	EndpointURL           string                        `json:"endpointURL,omitempty"`
	DocumentationURL      string                        `json:"documentationURL,omitempty"`
	APISpecificationURL   string                        `json:"apiSpecificationURL,omitempty"`
	Internal              bool                          `json:"internal,omitempty"`
	TechSupportContact    string                        `json:"techSupportContact,omitempty"`
	PublicSupportContact  string                        `json:"publicSupportContact,omitempty"`
	AuthorizationSettings *ServiceAuthorizationSettings `json:"authorizationSettings,omitempty"`
	Inways                []string                      `json:"inways,omitempty"`
}

type ServiceAuthorizationSettings struct {
	Mode string `json:"mode,omitempty"`
}

// ListServices returns a list of services
func (db PostgreSQLDirectoryDatabase) ListServices(ctx context.Context) ([]*Service, error) {
	result := []*Service{
		{
			Name: "service-name-a",
		},
		{
			Name: "service-name-b",
		},
	}

	return result, nil
}
