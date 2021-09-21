// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Organization struct {
	name         string
	serialNumber string
}

var organizationNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-._\s]{1,100}$`)
var organizationSerialNumberRegex = regexp.MustCompile(`^\d{20}$`)

func NewOrganization(name, serialNumber string) (*Organization, error) {
	err := validation.Validate(name, validation.Required, validation.Match(organizationNameRegex))
	if err != nil {
		return nil, fmt.Errorf("error validating organization name: %s", err)
	}

	err = validation.Validate(serialNumber, validation.Required, validation.Match(organizationSerialNumberRegex))
	if err != nil {
		return nil, fmt.Errorf("error validating organization serial number: %s", err)
	}

	return &Organization{
		name:         name,
		serialNumber: serialNumber,
	}, nil
}

func (i *Organization) Name() string {
	return i.name
}

func (i *Organization) SerialNumber() string {
	return i.serialNumber
}
