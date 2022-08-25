package service

import (
	"context"

	"go.uber.org/zap"

	"go.nlx.io/nlx/txlog-api/app"
	"go.nlx.io/nlx/txlog-api/app/command"
	"go.nlx.io/nlx/txlog-api/app/query"
	"go.nlx.io/nlx/txlog-api/domain/record"
)

type NewApplicationArgs struct {
	Context    context.Context
	Clock      command.Clock
	Logger     *zap.Logger
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
