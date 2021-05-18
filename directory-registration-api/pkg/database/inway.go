// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"fmt"
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// nolint:gocritic // these are valid regex patterns
var organizationNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-._\s]{1,100}$`)

// nolint:gocritic // these are valid regex patterns
var serviceNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-.\s]{1,100}$`)

type InsertAvailabilityParams struct {
	OrganizationName            string
	ServiceName                 string
	ServiceInternal             bool
	ServiceDocumentationURL     string
	InwayAPISpecificationType   string
	RequestInwayAddress         string
	ServicePublicSupportContact string
	ServiceTechSupportContact   string
	NlxVersion                  string
	OneTimeCosts                int32
	MonthlyCosts                int32
	RequestCosts                int32
}

func (params *InsertAvailabilityParams) Validate() error {
	return validation.ValidateStruct(
		params,
		validation.Field(&params.OrganizationName, validation.Required, validation.Match(organizationNameRegex)),
		validation.Field(&params.ServiceName, validation.Required, validation.Match(serviceNameRegex)),
		validation.Field(
			&params.RequestInwayAddress,
			validation.Required,
			validation.When(strings.Contains(params.RequestInwayAddress, ":"), is.DialString),
			validation.When(!strings.Contains(params.RequestInwayAddress, ":"), is.DNSName),
		),
		validation.Field(&params.NlxVersion, validation.When(params.NlxVersion != "unknown", is.Semver)),
	)
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
		params.ServicePublicSupportContact,
		params.ServiceTechSupportContact,
		params.NlxVersion,
		params.RequestCosts,
		params.MonthlyCosts,
		params.OneTimeCosts,
	)
	if err != nil {
		return fmt.Errorf("failed to execute the insert availability statement: %v", err)
	}

	return nil
}
