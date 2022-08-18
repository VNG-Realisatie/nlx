package pgadapter

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tabbed/pqtype"

	"go.nlx.io/nlx/txlog-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/txlog-api/domain"
)

func (r *PostgreSQLRepository) CreateRecord(ctx context.Context, record *domain.Record) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		err = tx.Rollback()
		if err != nil {
			if errors.Is(err, sql.ErrTxDone) {
				return
			}

			fmt.Printf("cannot rollback database transaction for create record: %e", err)
		}
	}()

	qtx := r.queries.WithTx(tx)

	recordID, err := qtx.CreateRecord(ctx, &queries.CreateRecordParams{
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

	for key, value := range record.DataSubjects() {
		err = qtx.CreateDataSubject(ctx, &queries.CreateDataSubjectParams{
			RecordID: recordID,
			Key:      key,
			Value:    value,
		})
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
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
