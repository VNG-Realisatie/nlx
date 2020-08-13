// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"fmt"

	"github.com/lib/pq"
)

type Service struct {
	Name                  string                        `json:"name,omitempty"`
	Organization          string                        `json:"organization,omitempty"`
	EndpointURL           string                        `json:"endpointURL,omitempty"`
	DocumentationURL      string                        `json:"documentationURL,omitempty"`
	APISpecificationURL   string                        `json:"apiSpecificationURL,omitempty"`
	APISpecificationType  string                        `json:"apiSpecificationType,omitempty"`
	Internal              bool                          `json:"internal,omitempty"`
	TechSupportContact    string                        `json:"techSupportContact,omitempty"`
	PublicSupportContact  string                        `json:"publicSupportContact,omitempty"`
	AuthorizationSettings *ServiceAuthorizationSettings `json:"authorizationSettings,omitempty"`
	Inways                []*Inway                      `json:"inways,omitempty"`
	InwayAddresses        []string                      `json:"inwayAddresses,omitempty"`
	HealthyStates         []bool                        `json:"healthyStates,omitempty"`
}

type Inway struct {
	Address string     `json:"address,omitempty"`
	State   InwayState `json:"state,omitempty"`
}

type InwayState string

const (
	InwayDOWN InwayState = "DOWN"
	InwayUP   InwayState = "UP"
)

type ServiceAuthorizationSettings struct {
	Mode string `json:"mode,omitempty"`
}

// ListServices returns a list of services
func (db PostgreSQLDirectoryDatabase) ListServices(ctx context.Context, organizationName string) ([]*Service, error) {
	var result []*Service

	rows, err := db.selectServicesStatement.Queryx(organizationName)
	if err != nil {
		return nil, fmt.Errorf("failed to execute stmtSelectServices: %v", err)
	}

	for rows.Next() {
		var respService = &Service{}
		var inwayAddresses = pq.StringArray{}
		var healthyStatuses = pq.BoolArray{}
		err = rows.Scan(
			&respService.Organization,
			&respService.Name,
			&respService.Internal,
			&inwayAddresses,
			&respService.DocumentationURL,
			&respService.APISpecificationType,
			&respService.PublicSupportContact,
			&healthyStatuses,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan into struct: %v", err)
		}

		if len(inwayAddresses) != len(healthyStatuses) {
			db.logger.Error("length of the inwayadresses does not match healthchecks")
		} else {
			respService.InwayAddresses = inwayAddresses
			respService.HealthyStates = healthyStatuses
		}

		var inway *Inway
		for inwayIndex, inwayAddress := range inwayAddresses {
			inway = &Inway{
				Address: inwayAddress,
				State:   InwayDOWN,
			}

			if healthyStatuses[inwayIndex] {
				inway.State = InwayUP
			}

			respService.Inways = append(respService.Inways, inway)
		}

		result = append(result, respService)
	}

	return result, nil
}
