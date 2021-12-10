// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Service struct {
	name string
}

func NewService(name string) (*Service, error) {
	err := validation.Validate(name, validation.Required)
	if err != nil {
		return nil, err
	}

	return &Service{
		name: name,
	}, nil
}

func (i *Service) Name() string {
	return i.name
}
