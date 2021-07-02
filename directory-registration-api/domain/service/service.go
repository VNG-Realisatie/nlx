// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package service

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Service struct {
	id                   uint
	name                 string
	organizationName     string
	internal             bool
	documentationURL     string
	apiSpecificationType SpecificationType
	publicSupportContact string
	techSupportContact   string
	oneTimeCosts         uint
	monthlyCosts         uint
	requestCosts         uint
}

type SpecificationType string

const (
	OpenAPI2 SpecificationType = "OpenAPI2"
	OpenAPI3 SpecificationType = "OpenAPI3"
)

func NewService(
	name,
	organizationName,
	documentationURL string,
	apiSpecificationType SpecificationType,
	publicSupportContact,
	techSupportContact string,
	oneTimeCosts,
	monthlyCosts,
	requestCosts uint,
	internal bool,
) (*Service, error) {
	err := validation.Validate(name, validation.Required, validation.Match(regexp.MustCompile(`^[a-zA-Z0-9-.\s]{1,100}$`)))
	if err != nil {
		return nil, fmt.Errorf("name: %s", err)
	}

	return &Service{
		name:                 name,
		organizationName:     organizationName,
		documentationURL:     documentationURL,
		apiSpecificationType: apiSpecificationType,
		publicSupportContact: publicSupportContact,
		techSupportContact:   techSupportContact,
		oneTimeCosts:         oneTimeCosts,
		monthlyCosts:         monthlyCosts,
		requestCosts:         requestCosts,
		internal:             internal,
	}, nil
}

func (i *Service) ID() uint {
	return i.id
}

func (i *Service) SetID(id uint) {
	i.id = id
}

func (i *Service) Name() string {
	return i.name
}

func (i *Service) OrganizationName() string {
	return i.organizationName
}

func (i *Service) DocumentationURL() string {
	return i.documentationURL
}

func (i *Service) APISpecificationType() SpecificationType {
	return i.apiSpecificationType
}

func (i *Service) PublicSupportContact() string {
	return i.publicSupportContact
}

func (i *Service) TechSupportContact() string {
	return i.techSupportContact
}

func (i *Service) OneTimeCosts() uint {
	return i.oneTimeCosts
}

func (i *Service) MonthlyCosts() uint {
	return i.monthlyCosts
}

func (i *Service) RequestCosts() uint {
	return i.requestCosts
}

func (i *Service) Internal() bool {
	return i.internal
}
