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

			fmt.Printf("cannot rollback database transaction while creating record: %e", err)
		}
	}()

	qtx := r.queries.WithTx(tx)

	dbRecord := &queries.CreateRecordParams{
		Direction:        queries.TransactionlogDirection(record.Direction()),
		SrcOrganization:  record.SourceOrganization(),
		DestOrganization: record.DestinationOrganization(),
		ServiceName:      record.ServiceName(),
		Created:          record.CreatedAt(),
		Data: pqtype.NullRawMessage{
			Valid:      true,
			RawMessage: record.Data(),
		},
		LogrecordID:    record.TransactionID(),
		Delegator:      record.Delegator(),
		OrderReference: record.OrderReference(),
	}

	recordID, err := qtx.CreateRecord(ctx, dbRecord)
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

func (r *PostgreSQLRepository) ListRecords(ctx context.Context, limit uint) ([]*domain.Record, error) {
	dbRecords, err := r.queries.ListRecords(ctx, int32(limit))
	if err != nil {
		return nil, err
	}

	records := make([]*domain.Record, len(dbRecords))

	for i, r := range dbRecords {
		var data json.RawMessage
		if r.Data.Valid {
			data = r.Data.RawMessage
		}

		records[i], err = domain.NewRecord(&domain.NewRecordArgs{
			SourceOrganization:      r.SrcOrganization,
			DestinationOrganization: r.DestOrganization,
			Direction:               domain.OrderDirection(r.Direction),
			ServiceName:             r.ServiceName,
			OrderReference:          r.OrderReference,
			Delegator:               r.Delegator,
			Data:                    data,
			TransactionID:           r.LogrecordID,
			CreatedAt:               r.Created,
		})
		if err != nil {
			return nil, err
		}
	}

	return records, nil
}
