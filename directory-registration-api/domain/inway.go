// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Inway struct {
	name         string
	organization *Organization
	address      string
	nlxVersion   string
	createdAt    time.Time
	updatedAt    time.Time
}

type NewInwayArgs struct {
	Name         string
	Organization *Organization
	Address      string
	NlxVersion   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

const NlxVersionUnknown = "unknown"

var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9-]{1,100}$`)

func NewInway(args *NewInwayArgs) (*Inway, error) {
	err := validation.ValidateStruct(
		args,
		validation.Field(&args.Name, validation.When(len(args.Name) > 0, validation.Match(nameRegex))),
		validation.Field(&args.Organization, validation.NotNil),
		validation.Field(&args.Address, validation.Required,
			validation.When(strings.Contains(args.Address, ":"), is.DialString),
			validation.When(!strings.Contains(args.Address, ":"), is.DNSName),
		),
		validation.Field(&args.NlxVersion, validation.When(args.NlxVersion != NlxVersionUnknown, is.Semver)),
		validation.Field(&args.CreatedAt, validation.Max(time.Now()).Error("must not be in the future")),
		validation.Field(&args.UpdatedAt, validation.Max(time.Now()).Error("must not be in the future")),
	)

	if err != nil {
		return nil, err
	}

	return &Inway{
		name:         args.Name,
		organization: args.Organization,
		address:      args.Address,
		nlxVersion:   args.NlxVersion,
		createdAt:    args.CreatedAt,
		updatedAt:    args.UpdatedAt,
	}, nil
}

func (i *Inway) Name() string {
	return i.name
}

func (i *Inway) Organization() *Organization {
	return i.organization
}

func (i *Inway) Address() string {
	return i.address
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
		"name: %s, organization: %s, address: %s, nlx version: %s, created at: %s, updated at: %s",
		i.Name(), i.organization.Name(), i.Address(), i.NlxVersion(), i.CreatedAt(), i.UpdatedAt(),
	)
}
