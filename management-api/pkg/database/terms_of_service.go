// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/management-api/domain"
)

var ErrInvalidDate = errors.New("date cannot be in the future")

func (db *PostgresConfigDatabase) AcceptTermsOfService(ctx context.Context, username string, createdAt time.Time) (bool, error) {
	if createdAt.After(time.Now()) {
		return false, ErrInvalidDate
	}

	tx, err := db.db.Begin()
	if err != nil {
		return false, err
	}

	defer func() {
		err = tx.Rollback()
		if err != nil {
			if errors.Is(err, sql.ErrTxDone) {
				return
			}

			fmt.Printf("cannot rollback database transaction for accepting the terms of service: %e", err)
		}
	}()

	qtx := db.queries.WithTx(tx)

	termsOfService, err := qtx.ListTermsOfService(ctx)
	if err != nil {
		return false, err
	}

	if len(termsOfService) > 0 {
		return true, nil
	}

	err = qtx.CreateTermsOfService(ctx, &queries.CreateTermsOfServiceParams{
		Username:  username,
		CreatedAt: createdAt,
	})
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return false, nil
}

func (db *PostgresConfigDatabase) GetTermsOfServiceStatus(ctx context.Context) (*domain.TermsOfServiceStatus, error) {
	allTos, err := db.queries.ListTermsOfService(ctx)
	if err != nil {
		return nil, err
	}

	if len(allTos) == 0 {
		return nil, ErrNotFound
	}

	tos := allTos[0]

	model, err := domain.NewTermsOfServiceStatus(&domain.NewTermsOfServiceStatusArgs{
		Username:  tos.Username,
		CreatedAt: tos.CreatedAt,
	})
	if err != nil {
		return nil, err
	}

	return model, nil
}
