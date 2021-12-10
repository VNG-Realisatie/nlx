package pgadapter

import (
	"context"
	"encoding/json"

	"github.com/tabbed/pqtype"
	"go.nlx.io/nlx/txlog-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/txlog-api/domain"
)

func (r *PostgreSQLRepository) CreateRecord(ctx context.Context, record *domain.Record) error {
	err := r.queries.CreateRecord(ctx, &queries.CreateRecordParams{
		Direction:        queries.TransactionlogDirection(record.Direction()),
		SrcOrganization:  record.Source().SerialNumber(),
		DestOrganization: record.Destination().SerialNumber(),
		ServiceName:      record.Service().Name(),
		Created:          record.CreatedAt(),
		Delegator:        record.Order().Delegator(),
		OrderReference:   record.Order().Reference(),
		Data: pqtype.NullRawMessage{
			Valid:      true,
			RawMessage: record.Data(),
		},
		LogrecordID: record.TransactionID(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLRepository) ListRecords(ctx context.Context, limit int32) ([]*domain.Record, error) {
	dbRecords, err := r.queries.ListRecords(ctx, limit)
	if err != nil {
		return nil, err
	}

	records := make([]*domain.Record, len(dbRecords))

	for i, r := range dbRecords {
		src, err := domain.NewOrganization(r.SrcOrganization)
		if err != nil {
			return nil, err
		}

		dest, err := domain.NewOrganization(r.DestOrganization)
		if err != nil {
			return nil, err
		}

		service, err := domain.NewService(r.ServiceName)
		if err != nil {
			return nil, err
		}

		order, err := domain.NewOrder(&domain.NewOrderArgs{
			Delegator: r.Delegator,
			Reference: r.OrderReference,
		})
		if err != nil {
			return nil, err
		}

		var data json.RawMessage
		if r.Data.Valid {
			data = r.Data.RawMessage
		}

		records[i], err = domain.NewRecord(&domain.NewRecordArgs{
			Source:        src,
			Destination:   dest,
			Direction:     domain.OrderDirection(r.Direction),
			Service:       service,
			Order:         order,
			Data:          data,
			TransactionID: r.LogrecordID,
			CreatedAt:     r.Created,
		})
		if err != nil {
			return nil, err
		}
	}

	return records, nil
}
