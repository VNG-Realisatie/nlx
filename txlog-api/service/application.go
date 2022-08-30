// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package service

import (
	"context"

	"go.nlx.io/nlx/txlog-api/app"
	"go.nlx.io/nlx/txlog-api/app/command"
	"go.nlx.io/nlx/txlog-api/app/query"
	"go.nlx.io/nlx/txlog-api/domain/record"
	"go.nlx.io/nlx/txlog-api/ports/logger"
)

type NewApplicationArgs struct {
	Context    context.Context
	Clock      command.Clock
	Logger     logger.Logger
	Repository record.Repository
}

func NewApplication(args *NewApplicationArgs) (*app.Application, error) {
	createRecordHandler, err := command.NewCreateRecordHandler(args.Repository, args.Clock, args.Logger)
	if err != nil {
		return nil, err
	}

	application := &app.Application{
		Queries: app.Queries{
			ListRecords: query.NewListRecordsHandler(args.Repository),
		},
		Commands: app.Commands{
			CreateRecord: createRecordHandler,
		},
	}

	return application, nil
}
