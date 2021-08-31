// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package inway

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Inway struct {
	name             string
	organizationName string
	address          string
	nlxVersion       string
	createdAt        time.Time
	updatedAt        time.Time
}

type NewInwayArgs struct {
	Name             string
	OrganizationName string
	Address          string
	NlxVersion       string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

const NlxVersionUnknown = "unknown"

var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9-]{1,100}$`)
var organizationNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-._\s]{1,100}$`)

func NewInway(args *NewInwayArgs) (*Inway, error) {
	err := validation.Validate(args.Name, validation.When(len(args.Name) > 0, validation.Match(nameRegex)))
	if err != nil {
		return nil, fmt.Errorf("name: %s", err)
	}

	err = validation.Validate(args.OrganizationName, validation.Required, validation.Match(organizationNameRegex))
	if err != nil {
		return nil, fmt.Errorf("organization name: %s", err)
	}

	err = validation.Validate(
		args.Address,
		validation.Required,
		validation.When(strings.Contains(args.Address, ":"), is.DialString),
		validation.When(!strings.Contains(args.Address, ":"), is.DNSName),
	)
	if err != nil {
		return nil, fmt.Errorf("address: %s", err)
	}

	err = validation.Validate(args.NlxVersion, validation.When(args.NlxVersion != NlxVersionUnknown, is.Semver))
	if err != nil {
		return nil, fmt.Errorf("nlx version: %s", err)
	}

	err = validation.Validate(args.CreatedAt, validation.Max(time.Now()))
	if err != nil {
		return nil, errors.New("created at: must not be in the future")
	}

	err = validation.Validate(args.UpdatedAt, validation.Max(time.Now()))
	if err != nil {
		return nil, errors.New("updated at: must not be in the future")
	}

	return &Inway{
		name:             args.Name,
		organizationName: args.OrganizationName,
		address:          args.Address,
		nlxVersion:       args.NlxVersion,
		createdAt:        args.CreatedAt,
		updatedAt:        args.UpdatedAt,
	}, nil
}

func (i *Inway) Name() string {
	return i.name
}

func (i *Inway) OrganizationName() string {
	return i.organizationName
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
		i.Name(), i.OrganizationName(), i.Address(), i.NlxVersion(), i.CreatedAt(), i.UpdatedAt(),
	)
}
