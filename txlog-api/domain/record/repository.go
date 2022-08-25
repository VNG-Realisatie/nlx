// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package record

import (
	"context"
)

type Repository interface {
	CreateRecord(context.Context, *Record) error
	ListRecords(ctx context.Context, limit uint) ([]*Record, error)

	Shutdown() error
}
