// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
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

type NewServiceArgs struct {
	Name                 string
	OrganizationName     string
	Internal             bool
	DocumentationURL     string
	APISpecificationType SpecificationType
	PublicSupportContact string
	TechSupportContact   string
	OneTimeCosts         uint
	MonthlyCosts         uint
	RequestCosts         uint
}

type SpecificationType string

const (
	OpenAPI2 SpecificationType = "OpenAPI2"
	OpenAPI3 SpecificationType = "OpenAPI3"
)

var serviceNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-.\s]{1,100}$`)
var organizationNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-._\s]{1,100}$`)

func NewService(args *NewServiceArgs) (*Service, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.Name, validation.Required, validation.Match(serviceNameRegex)),
		validation.Field(&args.OrganizationName, validation.Required, validation.Match(organizationNameRegex)),
	)
	if err != nil {
		return nil, err
	}

	return &Service{
		name:                 args.Name,
		organizationName:     args.OrganizationName,
		documentationURL:     args.DocumentationURL,
		apiSpecificationType: args.APISpecificationType,
		publicSupportContact: args.PublicSupportContact,
		techSupportContact:   args.TechSupportContact,
		oneTimeCosts:         args.OneTimeCosts,
		monthlyCosts:         args.MonthlyCosts,
		requestCosts:         args.RequestCosts,
		internal:             args.Internal,
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
