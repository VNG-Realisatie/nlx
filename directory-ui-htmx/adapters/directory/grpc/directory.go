// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package grpcdirectory

import (
	"go.nlx.io/nlx/directory-ui-htmx/domain"
)

type Directory struct {
	client Client
}

func New(client Client) domain.Repository {
	return &Directory{
		client: client,
	}
}

func (l *Directory) Shutdown() error {
	return l.client.Close()
}
