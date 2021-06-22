// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package inway

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidName = errors.New("invalid name")
)

type Inway struct {
	name             string
	organizationName string
	address          string
	nlxVersion       string
}

func NewInway(name, organizationName, address, nlxVersion string) (*Inway, error) {
	if !isNameValid(name) {
		return nil, ErrInvalidName
	}

	return &Inway{
		name:             name,
		organizationName: organizationName,
		address:          address,
		nlxVersion:       nlxVersion,
	}, nil
}

func (i Inway) Name() string {
	return i.name
}

func (i Inway) OrganizationName() string {
	return i.organizationName
}

func (i Inway) Address() string {
	return i.address
}

func (i Inway) NlxVersion() string {
	return i.nlxVersion
}

func isNameValid(name string) bool {
	if len(name) < 1 {
		return true
	}

	var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9-]{1,100}$`)

	return nameRegex.MatchString(name)
}
