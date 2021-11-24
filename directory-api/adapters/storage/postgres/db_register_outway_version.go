// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/common/nlxversion"
)

func (r *PostgreSQLRepository) RegisterOutwayVersion(_ context.Context, version nlxversion.Version, announcedAt time.Time) error {
	type params struct {
		AnnouncedAt time.Time `db:"announced_at"`
		Version     string    `db:"version"`
	}

	_, err := r.registerOutwayStmt.Exec(&params{
		AnnouncedAt: announcedAt,
		Version:     version.Version,
	})
	if err != nil {
		return fmt.Errorf("failed to log the outway version: %v", err)
	}

	return nil
}

func prepareRegisterOutwayStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	registerOutwayStatement, err := db.PrepareNamed(`
		INSERT INTO directory.outways (version, announced)
		VALUES (:version, :announced_at)
	`)
	if err != nil {
		return nil, err
	}

	return registerOutwayStatement, nil
}
