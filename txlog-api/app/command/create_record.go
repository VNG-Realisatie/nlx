package command

import (
	"context"

	"go.nlx.io/nlx/txlog-api/domain"
	"go.nlx.io/nlx/txlog-api/domain/txlog/storage"
)

type CreateRecordHandler struct {
	repository storage.Repository
}

func NewCreateRecordHandler(repository storage.Repository) *CreateRecordHandler {
	return &CreateRecordHandler{repository: repository}
}

func (h *CreateRecordHandler) Handle(ctx context.Context, record *domain.Record) error {
	return h.repository.CreateRecord(ctx, record)
}
