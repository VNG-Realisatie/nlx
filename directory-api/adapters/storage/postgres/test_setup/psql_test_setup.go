// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package pgadapter_test_setup

import (
	"os"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	pgadapter "go.nlx.io/nlx/directory-api/adapters/storage/postgres"
	"go.nlx.io/nlx/testing/testingutils"
)

const dbName = "test_directory"
const dbDriver = "txdb"

var setupOnce sync.Once

func setupDatabase(t *testing.T) {
	dsnBase := os.Getenv("POSTGRES_DSN")
	dsn, err := testingutils.CreateTestDatabase(dsnBase, dbName)
	if err != nil {
		t.Fatal(err)
	}

	dsnForMigrations := testingutils.AddQueryParamToAddress(dsn, "x-migrations-table", dbName)
	err = pgadapter.PostgreSQLPerformMigrations(dsnForMigrations)
	if err != nil {
		t.Fatal(err)
	}

	txdb.Register(dbDriver, "postgres", dsn)

	// This is necessary because the default BindVars for txdb isn't correct
	sqlx.BindDriver(dbDriver, sqlx.DOLLAR)

}

func New(t *testing.T) (*pgadapter.PostgreSQLRepository, func() error) {
	setupOnce.Do(func() {
		setupDatabase(t)
	})

	db, err := sqlx.Open(dbDriver, t.Name())
	require.NoError(t, err)

	db.MapperFunc(xstrings.ToSnakeCase)

	repo, err := pgadapter.New(zap.NewNop(), db)
	require.NoError(t, err)

	return repo, db.Close
}
