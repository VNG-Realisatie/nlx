// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package postgresadapter

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tabbed/pqtype"

	"go.nlx.io/nlx/txlog-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/txlog-api/domain/record"
)

func (r *PostgreSQLRepository) CreateRecord(ctx context.Context, model *record.Record) error {
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
		Direction:        queries.TransactionlogDirection(model.Direction()),
		SrcOrganization:  model.SourceOrganization(),
		DestOrganization: model.DestinationOrganization(),
		ServiceName:      model.ServiceName(),
		Created:          model.CreatedAt(),
		Data: pqtype.NullRawMessage{
			Valid:      true,
			RawMessage: model.Data(),
		},
		TransactionID:  model.TransactionID(),
		Delegator:      model.Delegator(),
		OrderReference: model.OrderReference(),
	}

	recordID, err := qtx.CreateRecord(ctx, dbRecord)
	if err != nil {
		return err
	}

	for key, value := range model.DataSubjects() {
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

func (r *PostgreSQLRepository) ListRecords(ctx context.Context, limit uint) ([]*record.Record, error) {
	dbRecords, err := r.queries.ListRecords(ctx, int32(limit))
	if err != nil {
		return nil, err
	}

	records := make([]*record.Record, len(dbRecords))

	for i, r := range dbRecords {
		var data json.RawMessage
		if r.Data.Valid {
			data = r.Data.RawMessage
		}

		records[i], err = record.NewRecord(&record.NewRecordArgs{
			SourceOrganization:      r.SrcOrganization,
			DestinationOrganization: r.DestOrganization,
			Direction:               record.OrderDirection(r.Direction),
			ServiceName:             r.ServiceName,
			OrderReference:          r.OrderReference,
			Delegator:               r.Delegator,
			Data:                    data,
			TransactionID:           r.TransactionID,
			CreatedAt:               r.Created,
		})
		if err != nil {
			return nil, err
		}
	}

	return records, nil
}
