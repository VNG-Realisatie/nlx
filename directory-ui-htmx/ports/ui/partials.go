// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

type ServicesSearchResults []*ServicesSearchResult

type ServicesSearchResult struct {
	ServiceName              string
	OrganizationSerialNumber string
	OrganizationName         string
	IsOnline                 bool
	APISpecificationType     string
}

type ParticipantsSearchResults []*ParticipantsSearchResult

type ParticipantsSearchResult struct {
	OrganizationName string
	ParticipantSince string
	ServicesCount    uint32
	InwaysCount      uint32
	OutwaysCount     uint32
}
