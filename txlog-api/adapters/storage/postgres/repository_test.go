// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package postgresadapter_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/txlog-api/domain/record"
)

func new(t *testing.T, enableFixtures bool) (record.Repository, func() error) {
	repo, close := NewTestRepository(t)
	if enableFixtures {
		err := loadFixtures(repo)
		require.NoError(t, err)
	}

	return repo, close
}

func loadFixtures(repo record.Repository) error {
	newRecordsArgs := []*record.NewRecordArgs{
		{
			SourceOrganization:      "0001",
			DestinationOrganization: "0002",
			Direction:               record.IN,
			ServiceName:             "test-service",
			OrderReference:          "test-reference",
			Delegator:               "0003",
			Data:                    []byte(`{"test": "data"}`),
			CreatedAt:               time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			TransactionID:           "abcde",
		},
	}

	for _, args := range newRecordsArgs {
		record, err := record.NewRecord(args)
		if err != nil {
			return err
		}

		err = repo.CreateRecord(context.Background(), record)
		if err != nil {
			return err
		}
	}

	return nil
}

func assertRecordInRepository(t *testing.T, repo record.Repository, r *record.Record) {
	require.NotNil(t, r)

	records, err := repo.ListRecords(context.Background(), 100)
	require.NoError(t, err)

	var found bool
	for _, record := range records {
		if record.TransactionID() == r.TransactionID() {
			found = true
			break
		}
	}

	require.Equal(t, true, found)
}
