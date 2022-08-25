package query

import (
	"context"

	"go.nlx.io/nlx/txlog-api/domain"
	"go.nlx.io/nlx/txlog-api/domain/txlog/storage"
)

type ListRecordsHandler struct {
	recordsRepository storage.Repository
}

func NewListRecordsHandler(repository storage.Repository) *ListRecordsHandler {
	return &ListRecordsHandler{
		recordsRepository: repository,
	}
}

func (l *ListRecordsHandler) Handle(ctx context.Context, limit uint) ([]*domain.Record, error) {
	return l.recordsRepository.ListRecords(ctx, limit)
}
