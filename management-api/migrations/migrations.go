// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package migrations

import (
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

//go:embed sql/*.sql
var migrations embed.FS

func RegisterDriver(driverName string) {
	source.Register(driverName, &driver{})
}

type driver struct {
	httpfs.PartialDriver
}

func (d *driver) Open(rawURL string) (source.Driver, error) {
	err := d.PartialDriver.Init(http.FS(migrations), "sql")
	if err != nil {
		return nil, err
	}

	return d, nil
}
