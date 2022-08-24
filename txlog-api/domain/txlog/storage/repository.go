// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package storage

import (
	"context"

	"go.nlx.io/nlx/txlog-api/domain"
)

type Repository interface {
	CreateRecord(context.Context, *domain.Record) error
	ListRecords(ctx context.Context, limit int32) ([]*domain.Record, error)

	Shutdown() error
}
