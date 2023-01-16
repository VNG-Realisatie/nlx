// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package query

import "time"

type Service struct {
	Name                     string
	OrganizationName         string
	OrganizationSerialNumber string
	IsOnline                 bool
	APISpecificationType     string
	PublicSupportContact     string
	DocumentationURL         string
}

type Participant struct {
	Name          string
	Since         time.Time
	ServicesCount uint
	InwaysCount   uint
	OutwaysCount  uint
}
