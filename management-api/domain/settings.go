// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Settings struct {
	organizationInwayName    string
	organizationEmailAddress string
}

var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9-]{1,100}$`)

func NewSettings(organizationInwayName, organizationEmail string) (*Settings, error) {
	err := validation.Validate(organizationInwayName, validation.Match(nameRegex))
	if err != nil {
		return nil, fmt.Errorf("inway name: %s", err)
	}

	err = validation.Validate(organizationEmail, is.Email)
	if err != nil {
		return nil, fmt.Errorf("organization email address: %s", err)
	}

	return &Settings{
		organizationInwayName:    organizationInwayName,
		organizationEmailAddress: organizationEmail,
	}, nil
}

func (s *Settings) OrganizationEmailAddress() string {
	return s.organizationEmailAddress
}

func (s *Settings) OrganizationInwayName() string {
	return s.organizationInwayName
}
