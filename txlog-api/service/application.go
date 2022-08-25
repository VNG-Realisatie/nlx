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
	Logger     *zap.Logger
	Repository record.Repository
}

func NewApplication(args *NewApplicationArgs) (*app.Application, error) {
	application := &app.Application{
		Queries: app.Queries{
			ListRecords: query.NewListRecordsHandler(args.Repository),
		},
		Commands: app.Commands{
			CreateRecord: command.NewCreateRecordHandler(args.Repository),
		},
	}

	return application, nil
}
