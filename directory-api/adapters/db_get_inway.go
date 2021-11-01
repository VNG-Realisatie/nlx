package adapters

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) GetInway(name, serialNumber string) (*domain.Inway, error) {
	type params struct {
		Name                     string `db:"name"`
		OrganizationSerialNumber string `db:"organization_serial_number"`
	}

	type dbInway struct {
		Name                     string    `db:"name"`
		Address                  string    `db:"address"`
		NlxVersion               string    `db:"nlx_version"`
		OrganizationName         string    `db:"organization_name"`
		OrganizationSerialNumber string    `db:"organization_serial_number"`
		CreatedAt                time.Time `db:"created_at"`
		UpdatedAt                time.Time `db:"updated_at"`
	}

	result := dbInway{}

	err := r.getInwayStmt.Get(&result, &params{
		Name:                     name,
		OrganizationSerialNumber: serialNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get inway (name: %s, serialNumber: %s): %s", name, serialNumber, err)
	}

	organizationModel, err := domain.NewOrganization(result.OrganizationName, result.OrganizationSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid organization model in database: %v", err)
	}

	model, err := domain.NewInway(&domain.NewInwayArgs{
		Name:         result.Name,
		Organization: organizationModel,
		Address:      result.Address,
		NlxVersion:   result.NlxVersion,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid inway model in database: %v", err)
	}

	return model, nil
}

func prepareGetInwayStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		select directory.inways.name as name, address, version as nlx_version, created_at, updated_at, directory.organizations.serial_number as organization_serial_number, directory.organizations.name as organization_name
		from directory.inways
		join directory.organizations
		    on directory.inways.organization_id = directory.organizations.id
		where
		      directory.organizations.serial_number = :organization_serial_number
		  and directory.inways.name = :name;
	`

	return db.PrepareNamed(query)
}
