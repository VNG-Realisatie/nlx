// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Organization struct {
	Name                string
	InsightIrmaEndpoint string
	InsightLogEndpoint  string
}

var (
	// ErrNoInwayWithAddress is returned when no inway is found for the provided address
	ErrNoInwayWithAddress = errors.New("no inway found for address")
	// ErrNoOrganization is returned when no organization is found
	ErrNoOrganization = errors.New("no organization found")
)

// SetInsightConfiguration updates the insight configuration options for an organization
func (db PostgreSQLDirectoryDatabase) SetInsightConfiguration(ctx context.Context, organizationName, insightAPIURL, irmaServerURL string) error {
	_, err := db.setInsightConfigurationStatement.Exec(organizationName, insightAPIURL, irmaServerURL)
	if err != nil {
		return fmt.Errorf("failed to execute setInsightConfigurationStatement: %v", err)
	}

	return nil
}

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
		return ErrNoOrganization
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
