// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	"fmt"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Inway struct {
	name                      string
	organization              *Organization
	isOrganizationInway       bool
	address                   string
	managementAPIProxyAddress string
	nlxVersion                string
	createdAt                 time.Time
	updatedAt                 time.Time
}

type NewInwayArgs struct {
	Name                      string
	Organization              *Organization
	IsOrganizationInway       bool
	Address                   string
	ManagementAPIProxyAddress string
	NlxVersion                string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

const NlxVersionUnknown = "unknown"

var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9-]{1,100}$`)

func NewInway(args *NewInwayArgs) (*Inway, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.Name, validation.When(len(args.Name) > 0, validation.Match(nameRegex))),
		validation.Field(&args.Organization, validation.NotNil),
		validation.Field(&args.Address, validation.Required, is.DialString),
		validation.Field(&args.ManagementAPIProxyAddress,
			validation.When(len(args.ManagementAPIProxyAddress) > 0, is.DialString)),
		validation.Field(&args.NlxVersion, validation.When(args.NlxVersion != NlxVersionUnknown, is.Semver)),
		validation.Field(&args.CreatedAt, validation.Max(time.Now()).Error("must not be in the future")),
		validation.Field(&args.UpdatedAt, validation.Max(time.Now()).Error("must not be in the future")),
	)

	if err != nil {
		return nil, err
	}

	return &Inway{
		name:                      args.Name,
		organization:              args.Organization,
		address:                   args.Address,
		managementAPIProxyAddress: args.ManagementAPIProxyAddress,
		nlxVersion:                args.NlxVersion,
		createdAt:                 args.CreatedAt,
		updatedAt:                 args.UpdatedAt,
		isOrganizationInway:       args.IsOrganizationInway,
	}, nil
}

func (i *Inway) Name() string {
	return i.name
}

func (i *Inway) Organization() *Organization {
	return i.organization
}

func (i *Inway) IsOrganizationInway() bool {
	return i.isOrganizationInway
}

func (i *Inway) Address() string {
	return i.address
}

func (i *Inway) ManagementAPIProxyAddress() string {
	return i.managementAPIProxyAddress
}

func (i *Inway) NlxVersion() string {
	return i.nlxVersion
}

func (i *Inway) CreatedAt() time.Time {
	return i.createdAt
}

func (i *Inway) UpdatedAt() time.Time {
	return i.updatedAt
}

func (i *Inway) ToString() string {
	return fmt.Sprintf(
		"name: %s, organization serial number: %s, organization name: %s, address: %s, nlx version: %s, created at: %s, updated at: %s",
		i.Name(), i.organization.serialNumber, i.organization.Name(), i.Address(), i.NlxVersion(), i.CreatedAt(), i.UpdatedAt(),
	)
}
