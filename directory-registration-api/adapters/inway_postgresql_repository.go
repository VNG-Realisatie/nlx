// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package adapters

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver

	"go.nlx.io/nlx/directory-registration-api/domain/inway"
)

type InwayPostgreSQLRepository struct {
	db              *sqlx.DB
	upsertInwayStmt *sqlx.NamedStmt
	getInwayStmt    *sqlx.NamedStmt
}

func NewInwayPostgreSQLRepository(db *sqlx.DB) (*InwayPostgreSQLRepository, error) {
	if db == nil {
		panic("missing db")
	}

	registerStmt, err := prepareRegisterStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare upsert inway statement: %s", err)
	}

	getStmt, err := prepareGetStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get inway statement: %s", err)
	}

	return &InwayPostgreSQLRepository{
		db:              db,
		upsertInwayStmt: registerStmt,
		getInwayStmt:    getStmt,
	}, nil
}

func (r *InwayPostgreSQLRepository) Register(model *inway.Inway) error {
	type registerParams struct {
		OrganizationName string `db:"organization_name"`
		Name             string `db:"inway_name"`
		Address          string `db:"inway_address"`
		NlxVersion       string `db:"inway_version"`
	}

	_, err := r.upsertInwayStmt.Exec(&registerParams{
		OrganizationName: model.OrganizationName(),
		Name:             model.Name(),
		Address:          model.Address(),
		NlxVersion:       model.NlxVersion(),
	})
	if err != nil {
		return fmt.Errorf("failed to register inway: %s", err)
	}

	return nil
}

func (r *InwayPostgreSQLRepository) GetInway(name, organizationName string) (*inway.Inway, error) {
	type dbInway struct {
		Name             string `db:"name"`
		Address          string `db:"address"`
		NlxVersion       string `db:"nlx_version"`
		OrganizationName string `db:"organization_name"`
	}

	type params struct {
		Name             string `db:"name"`
		OrganizationName string `db:"organization_name"`
	}

	result := dbInway{}

	err := r.getInwayStmt.Get(&result, &params{
		Name:             name,
		OrganizationName: organizationName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get inway (name: %s, organization: %s): %s", name, organizationName, err)
	}

	model, err := inway.NewInway(result.Name, result.OrganizationName, result.Address, result.NlxVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid inway model in database: %v", err)
	}

	return model, nil
}

func prepareRegisterStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
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
		    		(name, organization_id, address, version)
			 select :inway_name, organization.id, :inway_address, nullif(:inway_version, '')
			   from organization
	   	on conflict (name, organization_id) do update set 
	      			              name = 	excluded.name,
	      			              address = excluded.address, 
								  version = excluded.version;
	`

	return db.PrepareNamed(query)
}

func prepareGetStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		select directory.inways.name as name, address, version as nlx_version, directory.organizations.name as organization_name 
		from directory.inways
		join directory.organizations 
		    on directory.inways.organization_id = directory.organizations.id
		where 
		      directory.organizations.name = :organization_name
		  and directory.inways.name = :name;
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

func PostgreSQLPerformMigrations(dsn string) error {
	migrator, err := migrate.New("file://../../directory-db/migrations", dsn)
	if err != nil {
		return fmt.Errorf("setup migrator: %v", err)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("running migrations: %v", err)
	}

	return nil
}
