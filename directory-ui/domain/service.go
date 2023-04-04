// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package domain

type Service struct {
	Name                     string
	OrganizationName         string
	OrganizationSerialNumber string
	IsOnline                 bool
	APISpecificationType     string
	PublicSupportContact     string
	DocumentationURL         string
}
