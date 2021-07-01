// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import "go.nlx.io/nlx/directory-registration-api/domain/inway"

type Repository interface {
	Register(i *inway.Inway) error
	GetInway(name, organization string) (*inway.Inway, error)
}
