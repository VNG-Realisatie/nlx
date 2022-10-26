// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package migrations

import (
	"embed"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

const (
	driverName     = "embedhttpsessions"
	migrationTable = "http_sessions_schema_migrations"
)

var registerDriverOnce sync.Once

//go:embed sql/*.sql
var migrations embed.FS

func RegisterDriver(driverName string) {
	source.Register(driverName, &driver{})
}

type driver struct {
	httpfs.PartialDriver
}

func (d *driver) Open(string) (source.Driver, error) {
	err := d.PartialDriver.Init(http.FS(migrations), "sql")
	if err != nil {
		return nil, err
	}

	return d, nil
}

func setupMigrator(dsn string) (*migrate.Migrate, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	params := u.Query()
	params.Set("x-migrations-table", migrationTable)
	u.RawQuery = params.Encode()
	dsn = u.String()

	registerDriverOnce.Do(func() {
		RegisterDriver(driverName)
	})

	return migrate.New(fmt.Sprintf("%s://", driverName), dsn)
}

func PerformMigrations(dsn string) error {
	migrator, err := setupMigrator(dsn)
	if err != nil {
		return err
	}

	if err := migrator.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("migrations are up-to-date")
			return nil
		}

		return err
	}

	return nil
}

func MigrationStatus(dsn string) (version uint, dirty bool, err error) {
	migrator, err := setupMigrator(dsn)
	if err != nil {
		return 0, false, err
	}

	return migrator.Version()
}
