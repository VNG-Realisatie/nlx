// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) SetOrganizationEmailAddress(ctx context.Context, organization *domain.Organization, emailAddress string) error {
	arg := map[string]interface{}{
		"email_address":              emailAddress,
		"organization_name":          organization.Name(),
		"organization_serial_number": organization.SerialNumber(),
	}

	_, err := r.setOrganizationEmailAddressStmt.ExecContext(ctx, arg)

	return err
}

func prepareSetOrganizationEmailStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		INSERT into directory.organizations
			            (serial_number, name, email_address)
			     values (:organization_serial_number, :organization_name, :email_address)
			on conflict
			    		on constraint organizations_uq_serial_number
			  			do update
			      			set serial_number = excluded.serial_number,
			      			    name 		  = excluded.name,
								email_address = excluded.email_address
						returning id
	`

	return db.PrepareNamed(query)
}
