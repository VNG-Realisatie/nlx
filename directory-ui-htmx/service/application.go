// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package service

import (
	"context"
	"errors"

	"go.nlx.io/nlx/directory-ui-htmx/app"
	"go.nlx.io/nlx/directory-ui-htmx/app/query"
	"go.nlx.io/nlx/directory-ui-htmx/domain"
)

type NewApplicationArgs struct {
	Context             context.Context
	DirectoryRepository domain.Repository
}

func NewApplication(args *NewApplicationArgs) (*app.Application, error) {
	if args.DirectoryRepository == nil {
		return nil, errors.New("directory repository is required")
	}

	application := &app.Application{
		Queries: app.Queries{
			ListServices:     query.NewListServicesHandler(args.DirectoryRepository),
			ListParticipants: query.NewListParticipantsHandler(args.DirectoryRepository),
		},
	}

	return application, nil
}
