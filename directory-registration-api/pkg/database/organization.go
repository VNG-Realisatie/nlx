// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Organization struct {
	Name string
}

var (
	ErrNoInwayWithAddress   = errors.New("no inway found for address")
	ErrOrganizationNotFound = errors.New("no organization found")
)

func (db PostgreSQLDirectoryDatabase) SetOrganizationInway(ctx context.Context, organizationName, inwayAddress string) error {
	arg := map[string]interface{}{
		"inway_address":     inwayAddress,
		"organization_name": organizationName,
	}

	var ioID struct {
		InwayID        int
		OrganizationID int
	}

	err := db.selectInwayByAddressStatement.GetContext(ctx, &ioID, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoInwayWithAddress
		}

		return err
	}

	_, err = db.setOrganizationInwayStatement.ExecContext(ctx, ioID)
	if err != nil {
		return err
	}

	return nil
}

func (db PostgreSQLDirectoryDatabase) ClearOrganizationInway(ctx context.Context, organizationName string) error {
	r, err := db.clearOrganizationInwayStatement.ExecContext(ctx, organizationName)
	if err != nil {
		return err
	}

	n, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if n != 1 {
		return ErrOrganizationNotFound
	}

	return nil
}

func prepareSelectInwayByAddressStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	s := `
		SELECT i.id AS inway_id, i.organization_id
		FROM directory.inways i
		INNER JOIN directory.organizations o ON o.id = i.organization_id
		WHERE i.address = :inway_address
		AND o.name = :organization_name
	`

	return db.PrepareNamed(s)
}

func prepareSetOrganizationInwayStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	s := `
		UPDATE directory.organizations
		SET inway_id = :inway_id
		WHERE id = :organization_id
	`

	return db.PrepareNamed(s)
}

func prepareClearOrganizationInwayStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	s := `
		UPDATE directory.organizations
		SET inway_id = null
		WHERE name = $1
	`

	return db.Preparex(s)
}
