// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package adapters

import (
	"fmt"
	"time"

	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-registration-api/domain/inway"
)

type InwayPostgreSQLRepository struct {
	db              *sqlx.DB
	upsertInwayStmt *sqlx.NamedStmt
}

func NewInwayPostgreSQLRepository(db *sqlx.DB) (*InwayPostgreSQLRepository, error) {
	if db == nil {
		panic("missing db")
	}

	upsertInwayStmt, err := prepareUpsertInwayStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare upsert inway statement: %s", err)
	}

	return &InwayPostgreSQLRepository{
		db:              db,
		upsertInwayStmt: upsertInwayStmt,
	}, nil
}

func (r *InwayPostgreSQLRepository) Register(model *inway.Inway) error {
	type upsertInwayParams struct {
		OrganizationName string `db:"organization_name"`
		InwayAddress     string `db:"inway_address"`
		InwayVersion     string `db:"inway_version"`
	}

	_, err := r.upsertInwayStmt.Exec(&upsertInwayParams{
		OrganizationName: model.OrganizationName(),
		InwayAddress:     model.Address(),
		InwayVersion:     model.NlxVersion(),
	})
	if err != nil {
		return fmt.Errorf("failed upsert inway and its organization: %s", err)
	}

	return nil
}

func prepareUpsertInwayStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		with organization as (
			insert into directory.organizations 
			            (name)
			     values (:organization_name)
			on conflict 
			    		on constraint organizations_uq_name
			  			do update 
			      			set name = excluded.name -- no-op update to return id
						returning id
		) 
		insert into directory.inways 
		    		(organization_id, address, version)
			 select organization.id, :inway_address, nullif(:inway_version, '')
			   from organization
	   	on conflict 
	   	    		on constraint inways_uq_address
	      			do update set 
	      			              address = excluded.address, 
								  version = excluded.version;
	`

	return db.PrepareNamed(query)
}

func NewPostgreSQLConnection(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open connection to postgres: %s", err)
	}

	const (
		FiveMinutes        = 5 * time.Minute
		MaxIdleConnections = 2
	)

	db.SetConnMaxLifetime(FiveMinutes)
	db.SetMaxIdleConns(MaxIdleConnections)
	db.MapperFunc(xstrings.ToSnakeCase)

	return db, nil
}
