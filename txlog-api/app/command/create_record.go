// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package command

import (
	"context"

	"go.nlx.io/nlx/txlog-api/domain/record"
)

type CreateRecordHandler struct {
	repository record.Repository
}

func NewCreateRecordHandler(repository record.Repository) *CreateRecordHandler {
	return &CreateRecordHandler{repository: repository}
}

func (h *CreateRecordHandler) Handle(ctx context.Context, model *record.Record) error {
	return h.repository.CreateRecord(ctx, model)
}
