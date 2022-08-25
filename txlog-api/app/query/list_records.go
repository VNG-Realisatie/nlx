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

func (l *ListRecordsHandler) Handle(ctx context.Context, limit uint) ([]*record.Record, error) {
	return l.recordsRepository.ListRecords(ctx, limit)
}
