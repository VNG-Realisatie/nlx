// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	app_errors "go.nlx.io/nlx/txlog-api/app/errors"
	"go.nlx.io/nlx/txlog-api/domain/record"
	"go.nlx.io/nlx/txlog-api/ports/logger"
)

type Clock interface {
	Now() time.Time
}

type CreateRecordHandler struct {
	clock      Clock
	repository record.Repository
	logger     logger.Logger
}

func NewCreateRecordHandler(repository record.Repository, clock Clock, lgr logger.Logger) (*CreateRecordHandler, error) {
	if repository == nil {
		return nil, errors.New("repository is required")
	}

	if clock == nil {
		return nil, errors.New("repository is required")
	}

	if lgr == nil {
		return nil, errors.New("logger is required")
	}

	return &CreateRecordHandler{
		repository: repository,
		clock:      clock,
		logger:     lgr,
	}, nil
}

type NewRecordArgs struct {
	SourceOrganization      string
	DestinationOrganization string
	Direction               string
	ServiceName             string
	OrderReference          string
	Delegator               string
	Data                    json.RawMessage
	TransactionID           string
	DataSubjects            map[string]string
}

func (h *CreateRecordHandler) Handle(ctx context.Context, args *NewRecordArgs) error {
	direction := record.OUT

	if args.Direction == "in" {
		direction = record.IN
	}

	model, err := record.NewRecord(&record.NewRecordArgs{
		SourceOrganization:      args.SourceOrganization,
		DestinationOrganization: args.DestinationOrganization,
		Direction:               direction,
		ServiceName:             args.ServiceName,
		OrderReference:          args.OrderReference,
		Delegator:               args.Delegator,
		Data:                    args.Data,
		TransactionID:           args.TransactionID,
		CreatedAt:               h.clock.Now(),
		DataSubjects:            args.DataSubjects,
	})
	if err != nil {
		return app_errors.NewIncorrectInputError(fmt.Sprintf("invalid input: %s", err))
	}

	err = h.repository.CreateRecord(ctx, model)
	if err != nil {
		h.logger.Error("create record", err)
		return err
	}

	return nil
}
