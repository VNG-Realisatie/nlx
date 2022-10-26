// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package pgsessionstore_test

import (
	"database/sql"
	"os"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"go.nlx.io/nlx/management-api/pkg/oidc/pgsessionstore"
	"go.nlx.io/nlx/management-api/pkg/oidc/pgsessionstore/migrations"
	"go.nlx.io/nlx/testing/testingutils"
)

var setupOnce sync.Once

const dbDriver = "txdb"

func setup(t *testing.T) {
	setupOnce.Do(func() {
		setupPostgreSQL(t)
	})
}

func setupPostgreSQL(t *testing.T) {
	dbName := "test_httpsessions"

	dsnBase := os.Getenv("POSTGRES_DSN")

	dsn, err := testingutils.CreateTestDatabase(dsnBase, dbName)
	if err != nil {
		t.Fatal(err)
	}

	dsnForMigrations := testingutils.AddQueryParamToAddress(dsn, "x-migrations-table", dbName)

	err = migrations.PerformMigrations(dsnForMigrations)
	if err != nil {
		t.Fatal(err)
	}

	txdb.Register(dbDriver, "postgres", dsn)
}

func newDB(t *testing.T, id string) *sql.DB {
	db, err := sql.Open(dbDriver, id)
	require.NoError(t, err)

	return db
}

func New(t *testing.T, secret string) *pgsessionstore.PGStore {
	setup(t)

	s, err := pgsessionstore.New(zaptest.NewLogger(t), newDB(t, t.Name()), []byte(secret))
	require.NoError(t, err)

	return s
}
