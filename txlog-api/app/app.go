package app

import (
	"go.nlx.io/nlx/txlog-api/app/command"
	"go.nlx.io/nlx/txlog-api/app/query"
)

type Application struct {
	Queries  Queries
	Commands Commands
}

type Queries struct {
	ListRecords *query.ListRecordsHandler
}

type Commands struct {
	CreateRecord *command.CreateRecordHandler
}
