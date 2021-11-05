// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ServiceCosts struct {
	OneTime uint
	Monthly uint
	Request uint
}

type Service struct {
	id                   uint
	name                 string
	organization         *Organization
	internal             bool
	documentationURL     string
	apiSpecificationType SpecificationType
	publicSupportContact string
	techSupportContact   string
	costs                *ServiceCosts
	inways               []*ServiceInway
}

type NewServiceArgs struct {
	Name                 string
	Organization         *Organization
	Internal             bool
	DocumentationURL     string
	APISpecificationType SpecificationType
	PublicSupportContact string
	TechSupportContact   string
	Costs                *ServiceCosts
	Inways               []*ServiceInway
}

type SpecificationType string

const (
	OpenAPI2 SpecificationType = "OpenAPI2"
	OpenAPI3 SpecificationType = "OpenAPI3"
)

// Usage is documented in /docs/docs/reference-information/data-validation.md
var serviceNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-.\s]{1,100}$`)

func NewService(args *NewServiceArgs) (*Service, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.Name, validation.Required, validation.Match(serviceNameRegex)),
		validation.Field(&args.Organization, validation.Required),
		validation.Field(&args.Costs, validation.Required),
	)
	if err != nil {
		return nil, err
	}

	return &Service{
		name:                 args.Name,
		organization:         args.Organization,
		documentationURL:     args.DocumentationURL,
		apiSpecificationType: args.APISpecificationType,
		publicSupportContact: args.PublicSupportContact,
		techSupportContact:   args.TechSupportContact,
		costs:                args.Costs,
		inways:               args.Inways,
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

func (i *Service) Organization() *Organization {
	return i.organization
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

func (i *Service) Costs() *ServiceCosts {
	return i.costs
}

func (i *Service) Internal() bool {
	return i.internal
}

func (i *Service) Inways() []*ServiceInway {
	return i.inways
}

type ServiceInway struct {
	address string
	state   ServiceInwayState
}

type ServiceInwayState string

const (
	InwayDOWN ServiceInwayState = "DOWN"
	InwayUP   ServiceInwayState = "UP"
)

type NewServiceInwayArgs struct {
	Address string
	State   ServiceInwayState
}

func NewServiceInway(args *NewServiceInwayArgs) (*ServiceInway, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.Address, validation.Required),
		validation.Field(&args.State, validation.Required),
	)
	if err != nil {
		return nil, err
	}

	return &ServiceInway{
		address: args.Address,
		state:   args.State,
	}, nil
}

func (i *ServiceInway) Address() string {
	return i.address
}

func (i *ServiceInway) State() ServiceInwayState {
	return i.state
}
