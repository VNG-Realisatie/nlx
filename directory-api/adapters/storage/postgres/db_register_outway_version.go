// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/common/nlxversion"
)

func (r *PostgreSQLRepository) RegisterOutwayVersion(_ context.Context, version nlxversion.Version) error {
	_, err := r.registerOutwayStmt.Exec(version)
	if err != nil {
		return fmt.Errorf("failed to log the outway version: %v", err)
	}

	return nil
}

func prepareRegisterOutwayStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	registerOutwayStatement, err := db.PrepareNamed(`
		INSERT INTO directory.outways (version)
		VALUES (:version)
	`)
	if err != nil {
		return nil, err
	}

	return registerOutwayStatement, nil
}
