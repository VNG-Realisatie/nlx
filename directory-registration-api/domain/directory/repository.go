// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"go.nlx.io/nlx/directory-registration-api/domain/inway"
	"go.nlx.io/nlx/directory-registration-api/domain/service"
)

type Repository interface {
	RegisterInway(*inway.Inway) error
	GetInway(name, organization string) (*inway.Inway, error)

	RegisterService(*service.Service) error
	GetService(id uint) (*service.Service, error)
}
