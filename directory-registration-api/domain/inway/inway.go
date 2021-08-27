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

const NlxVersionUnknown = "unknown"

func NewInway(name, organizationName, address, nlxVersion string, createdAt, updatedAt time.Time) (*Inway, error) {
	err := validation.Validate(name, validation.When(len(name) > 0, validation.Match(regexp.MustCompile(`^[a-zA-Z0-9-]{1,100}$`))))
	if err != nil {
		return nil, fmt.Errorf("name: %s", err)
	}

	err = validation.Validate(organizationName, validation.Required, validation.Match(regexp.MustCompile(`^[a-zA-Z0-9-._\s]{1,100}$`)))
	if err != nil {
		return nil, fmt.Errorf("organization name: %s", err)
	}

	err = validation.Validate(
		address,
		validation.Required,
		validation.When(strings.Contains(address, ":"), is.DialString),
		validation.When(!strings.Contains(address, ":"), is.DNSName),
	)
	if err != nil {
		return nil, fmt.Errorf("address: %s", err)
	}

	err = validation.Validate(nlxVersion, validation.When(nlxVersion != NlxVersionUnknown, is.Semver))
	if err != nil {
		return nil, fmt.Errorf("nlx version: %s", err)
	}

	err = validation.Validate(createdAt, validation.Max(time.Now()))
	if err != nil {
		return nil, errors.New("created at: must not be in the future")
	}

	err = validation.Validate(updatedAt, validation.Max(time.Now()))
	if err != nil {
		return nil, errors.New("updated at: must not be in the future")
	}

	return &Inway{
		name:             name,
		organizationName: organizationName,
		address:          address,
		nlxVersion:       nlxVersion,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
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
