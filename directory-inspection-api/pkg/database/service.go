// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"fmt"

	"github.com/lib/pq"
)

type Service struct {
	Name                 string
	Organization         string
	EndpointURL          string
	DocumentationURL     string
	APISpecificationURL  string
	APISpecificationType string
	Internal             bool
	TechSupportContact   string
	PublicSupportContact string
	OneTimeCosts         int
	MonthlyCosts         int
	RequestCosts         int
	Inways               []*Inway
	InwayAddresses       []string
	HealthyStates        []bool
}

type Inway struct {
	Address string
	State   InwayState
}

type InwayState string

const (
	InwayDOWN InwayState = "DOWN"
	InwayUP   InwayState = "UP"
)

func (db PostgreSQLDirectoryDatabase) ListServices(_ context.Context, organizationName string) ([]*Service, error) {
	var result []*Service

	rows, err := db.selectServicesStatement.Queryx(organizationName)
	if err != nil {
		return nil, fmt.Errorf("failed to execute stmtSelectServices: %v", err)
	}

	for rows.Next() {
		var (
			respService     = &Service{}
			inwayAddresses  = pq.StringArray{}
			healthyStatuses = pq.BoolArray{}
		)

		err = rows.Scan(
			&respService.Organization,
			&respService.Name,
			&respService.Internal,
			&respService.OneTimeCosts,
			&respService.MonthlyCosts,
			&respService.RequestCosts,
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
