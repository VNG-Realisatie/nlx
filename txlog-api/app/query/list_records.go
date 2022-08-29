// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package query

import (
	"context"

	"go.nlx.io/nlx/txlog-api/domain/record"
)

type ListRecordsHandler struct {
	recordsRepository record.Repository
}

func NewListRecordsHandler(repository record.Repository) *ListRecordsHandler {
	return &ListRecordsHandler{
		recordsRepository: repository,
	}
}

func (l *ListRecordsHandler) Handle(ctx context.Context, limit uint) ([]*Record, error) {
	records, err := l.recordsRepository.ListRecords(ctx, limit)
	if err != nil {
		return nil, err
	}

	result := make([]*Record, len(records))

	for i, model := range records {
		result[i] = &Record{
			SourceOrganization:      model.SourceOrganization(),
			DestinationOrganization: model.DestinationOrganization(),
			Direction:               string(model.Direction()),
			ServiceName:             model.ServiceName(),
			OrderReference:          model.OrderReference(),
			Delegator:               model.Delegator(),
			Data:                    model.Data(),
			TransactionID:           model.TransactionID(),
			CreatedAt:               model.CreatedAt(),
		}
	}

	return result, nil
}
