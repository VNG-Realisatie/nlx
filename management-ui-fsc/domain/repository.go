// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package domain

type Repository interface {
	Shutdown() error
}
