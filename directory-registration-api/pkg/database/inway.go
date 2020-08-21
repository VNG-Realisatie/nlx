// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"fmt"
)

type InsertAvailabilityParams struct {
	OrganizationName            string
	ServiceName                 string
	ServiceInternal             bool
	ServiceDocumentationURL     string
	InwayAPISpecificationType   string
	RequestInwayAddress         string
	ServiceInsightAPIURL        string
	ServiceIrmaAPIURL           string
	ServicePublicSupportContact string
	ServiceTechSupportContact   string
	NlxVersion                  string
}

// InsertAvailability updates the availability
// NOTE: what does this method actually do? Is InsertAvailability the correct name?
func (db PostgreSQLDirectoryDatabase) InsertAvailability(params *InsertAvailabilityParams) error {
	_, err := db.insertAvailabilityStatement.Exec(
		params.OrganizationName,
		params.ServiceName,
		params.ServiceInternal,
		params.ServiceDocumentationURL,
		params.InwayAPISpecificationType,
		params.RequestInwayAddress,
		params.ServiceInsightAPIURL,
		params.ServiceIrmaAPIURL,
		params.ServicePublicSupportContact,
		params.ServiceTechSupportContact,
		params.NlxVersion,
	)
	if err != nil {
		return fmt.Errorf("failed to execute the insert availability statement: %v", err)
	}

	return nil
}
