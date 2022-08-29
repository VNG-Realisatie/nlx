// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package postgresadapter_test

import (
	"context"
	"github.com/huandu/xstrings"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/testing/testingutils"
	"go.nlx.io/nlx/txlog-api/adapters/storage/postgres"
	"go.nlx.io/nlx/txlog-api/domain/record"
)

const dbName = "test_txlog"
const dbDriver = "txdb"

var setupOnce sync.Once

func new(t *testing.T, enableFixtures bool) (record.Repository, func() error) {
	setupOnce.Do(func() {
		setupDatabase(t)
	})

	db, err := sqlx.Open(dbDriver, t.Name())
	require.NoError(t, err)

	db.MapperFunc(xstrings.ToSnakeCase)

	repo, err := postgresadapter.New(db)
	require.NoError(t, err)

	if enableFixtures {
		err := loadFixtures(repo)
		require.NoError(t, err)
	}

	return repo, db.Close
}

func setupDatabase(t *testing.T) {
	dsnBase := os.Getenv("POSTGRES_DSN")
	dsn, err := testingutils.CreateTestDatabase(dsnBase, dbName)
	if err != nil {
		t.Fatal(err)
	}

	dsnForMigrations := testingutils.AddQueryParamToAddress(dsn, "x-migrations-table", dbName)
	err = postgresadapter.PerformMigrations(dsnForMigrations)
	if err != nil {
		t.Fatal(err)
	}

	txdb.Register(dbDriver, "postgres", dsn)

	// This is necessary because the default BindVars for txdb isn't correct
	sqlx.BindDriver(dbDriver, sqlx.DOLLAR)

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
		recordModel, err := record.NewRecord(args)
		if err != nil {
			return err
		}

		err = repo.CreateRecord(context.Background(), recordModel)
		if err != nil {
			return err
		}
	}

	return nil
}
