// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package organization

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Organization struct {
	name string
}

var organizationNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-._\s]{1,100}$`)

func NewOrganization(name string) (*Organization, error) {
	err := validation.Validate(name, validation.Required, validation.Match(organizationNameRegex))
	if err != nil {
		return nil, fmt.Errorf("organization name: %s", err)
	}

	return &Organization{
		name: name,
	}, nil
}

func (i *Organization) Name() string {
	return i.name
}
