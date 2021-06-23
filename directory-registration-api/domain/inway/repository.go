// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package inway

type Repository interface {
	Register(i *Inway) error
}
