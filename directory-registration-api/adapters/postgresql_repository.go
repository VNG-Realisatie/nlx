// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package adapters

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver

	"go.nlx.io/nlx/directory-registration-api/domain/inway"
	"go.nlx.io/nlx/directory-registration-api/domain/service"
)

const timeLayout = time.RFC3339

var (
	ErrDuplicateAddress     = errors.New("another inway is already registered with this address")
	ErrNoInwayWithAddress   = errors.New("no inway found for address")
	ErrOrganizationNotFound = errors.New("no organization found")
)

type PostgreSQLRepository struct {
	db                         *sqlx.DB
	registerInwayStmt          *sqlx.NamedStmt
	getInwayStmt               *sqlx.NamedStmt
	registerServiceStmt        *sqlx.NamedStmt
	getServiceStmt             *sqlx.NamedStmt
	selectInwayByAddressStmt   *sqlx.NamedStmt
	setOrganizationInwayStmt   *sqlx.NamedStmt
	clearOrganizationInwayStmt *sqlx.NamedStmt
}

func NewPostgreSQLRepository(db *sqlx.DB) (*PostgreSQLRepository, error) {
	if db == nil {
		panic("missing db")
	}

	registerInwayStmt, err := prepareRegisterInwayStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare register inway statement: %s", err)
	}

	getInwayStmt, err := prepareGetInwayStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get inway statement: %s", err)
	}

	registerServiceStmt, err := prepareRegisterServiceStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare register service statement: %s", err)
	}

	getServiceStmt, err := prepareGetServiceStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get service statement: %s", err)
	}

	selectInwayByAddressStmt, err := prepareSelectInwayByAddressStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select inway by address statement: %s", err)
	}

	setOrganizationInwayStmt, err := prepareSetOrganizationInwayStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare set organization inway statement: %s", err)
	}

	clearOrganizationInwayStmt, err := prepareClearOrganizationInwayStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare clear organization inway statement: %s", err)
	}

	return &PostgreSQLRepository{
		db:                         db,
		registerInwayStmt:          registerInwayStmt,
		getInwayStmt:               getInwayStmt,
		registerServiceStmt:        registerServiceStmt,
		getServiceStmt:             getServiceStmt,
		selectInwayByAddressStmt:   selectInwayByAddressStmt,
		setOrganizationInwayStmt:   setOrganizationInwayStmt,
		clearOrganizationInwayStmt: clearOrganizationInwayStmt,
	}, nil
}

func (r *PostgreSQLRepository) RegisterInway(model *inway.Inway) error {
	type registerParams struct {
		OrganizationName string    `db:"organization_name"`
		Name             string    `db:"inway_name"`
		Address          string    `db:"inway_address"`
		NlxVersion       string    `db:"inway_version"`
		CreatedAt        time.Time `db:"inway_created_at"`
		UpdatedAt        time.Time `db:"inway_updated_at"`
	}

	_, err := r.registerInwayStmt.Exec(&registerParams{
		OrganizationName: model.OrganizationName(),
		Name:             model.Name(),
		Address:          model.Address(),
		NlxVersion:       model.NlxVersion(),
		CreatedAt:        model.CreatedAt(),
		UpdatedAt:        model.UpdatedAt(),
	})

	if err != nil && err.Error() == "pq: duplicate key value violates unique constraint \"inways_uq_address\"" {
		return ErrDuplicateAddress
	}

	return err
}

func (r *PostgreSQLRepository) GetInway(name, organizationName string) (*inway.Inway, error) {
	type params struct {
		Name             string `db:"name"`
		OrganizationName string `db:"organization_name"`
	}

	type dbInway struct {
		Name             string `db:"name"`
		Address          string `db:"address"`
		NlxVersion       string `db:"nlx_version"`
		OrganizationName string `db:"organization_name"`
		CreatedAt        string `db:"created_at"`
		UpdatedAt        string `db:"updated_at"`
	}

	result := dbInway{}

	err := r.getInwayStmt.Get(&result, &params{
		Name:             name,
		OrganizationName: organizationName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get inway (name: %s, organization: %s): %s", name, organizationName, err)
	}

	createdAt, err := time.Parse(timeLayout, result.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse created at timestamp (created at: %s): %s", result.CreatedAt, err)
	}

	updatedAt, err := time.Parse(timeLayout, result.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse updated at timestamp (updated at: %s): %s", result.UpdatedAt, err)
	}

	model, err := inway.NewInway(result.Name, result.OrganizationName, result.Address, result.NlxVersion, createdAt, updatedAt)
	if err != nil {
		return nil, fmt.Errorf("invalid inway model in database: %v", err)
	}

	return model, nil
}

func (r *PostgreSQLRepository) RegisterService(model *service.Service) error {
	type registerParams struct {
		Name                 string `db:"name"`
		OrganizationName     string `db:"organization_name"`
		DocumentationURL     string `db:"documentation_url"`
		APISpecificationType string `db:"api_specification_type"`
		PublicSupportContact string `db:"public_support_contact"`
		TechSupportContact   string `db:"tech_support_contact"`
		OneTimeCosts         int32  `db:"one_time_costs"`
		MonthlyCosts         int32  `db:"monthly_costs"`
		RequestCosts         int32  `db:"request_costs"`
		Internal             bool   `db:"internal"`
	}

	type dbResult struct {
		ID uint `db:"id"`
	}

	result := dbResult{}

	err := r.registerServiceStmt.Get(
		&result,
		&registerParams{
			Name:                 model.Name(),
			OrganizationName:     model.OrganizationName(),
			Internal:             model.Internal(),
			DocumentationURL:     model.DocumentationURL(),
			APISpecificationType: string(model.APISpecificationType()),
			PublicSupportContact: model.PublicSupportContact(),
			TechSupportContact:   model.TechSupportContact(),
			OneTimeCosts:         int32(model.OneTimeCosts()),
			MonthlyCosts:         int32(model.MonthlyCosts()),
			RequestCosts:         int32(model.RequestCosts()),
		})

	if err != nil {
		return fmt.Errorf("failed to register service: %v", err)
	}

	model.SetID(result.ID)

	return err
}

func (r *PostgreSQLRepository) GetService(id uint) (*service.Service, error) {
	type params struct {
		ID uint `db:"id"`
	}

	type dbService struct {
		ID                   uint   `db:"id"`
		Name                 string `db:"name"`
		OrganizationName     string `db:"organization_name"`
		DocumentationURL     string `db:"documentation_url"`
		APISpecificationType string `db:"api_specification_type"`
		PublicSupportContact string `db:"public_support_contact"`
		TechSupportContact   string `db:"tech_support_contact"`
		OneTimeCosts         int32  `db:"one_time_costs"`
		MonthlyCosts         int32  `db:"monthly_costs"`
		RequestCosts         int32  `db:"request_costs"`
		Internal             bool   `db:"internal"`
	}

	result := dbService{}

	err := r.getServiceStmt.Get(&result, &params{ID: id})
	if err != nil {
		return nil, fmt.Errorf("failed to get service with id %v: %s", id, err)
	}

	model, err := service.NewService(
		result.Name,
		result.OrganizationName,
		result.DocumentationURL,
		service.SpecificationType(result.APISpecificationType),
		result.PublicSupportContact,
		result.TechSupportContact,
		uint(result.OneTimeCosts),
		uint(result.MonthlyCosts),
		uint(result.RequestCosts),
		result.Internal,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid inway model in database: %v", err)
	}

	model.SetID(result.ID)

	return model, nil
}

func (r *PostgreSQLRepository) SetOrganizationInway(ctx context.Context, organizationName, inwayAddress string) error {
	arg := map[string]interface{}{
		"inway_address":     inwayAddress,
		"organization_name": organizationName,
	}

	var ioID struct {
		InwayID        int
		OrganizationID int
	}

	err := r.selectInwayByAddressStmt.GetContext(ctx, &ioID, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoInwayWithAddress
		}

		return err
	}

	_, err = r.setOrganizationInwayStmt.ExecContext(ctx, ioID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLRepository) ClearOrganizationInway(ctx context.Context, organizationName string) error {
	arg := map[string]interface{}{
		"name": organizationName,
	}

	row, err := r.clearOrganizationInwayStmt.ExecContext(ctx, arg)
	if err != nil {
		return err
	}

	n, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if n != 1 {
		return ErrOrganizationNotFound
	}

	return nil
}

func prepareRegisterInwayStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
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
		    		(name, organization_id, address, version, created_at, updated_at)
			 select :inway_name, organization.id, :inway_address, nullif(:inway_version, ''), :inway_created_at, :inway_updated_at
			   from organization
	   	on conflict (name, organization_id) do update set 
	      			              name = 	excluded.name,
	      			              address = excluded.address, 
								  version = excluded.version,
								  created_at = excluded.created_at,
								  updated_at = excluded.updated_at;
	`

	return db.PrepareNamed(query)
}

func prepareGetInwayStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
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

func prepareRegisterServiceStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		with organization as (
		    select id from directory.organizations where directory.organizations.name = :organization_name
		),
	    inway as (
		    select directory.inways.id from directory.inways, organization where organization_id = organization.id
		),
		service as (
			insert into directory.services 
			    		(organization_id, name, internal, documentation_url, api_specification_type, public_support_contact, tech_support_contact, request_costs, monthly_costs, one_time_costs)
				 select organization.id, :name, :internal, nullif(:documentation_url, ''), nullif(:api_specification_type, ''), nullif(:public_support_contact, ''), nullif(:tech_support_contact, ''), :request_costs, :monthly_costs, :one_time_costs
				   from organization
			on conflict on constraint services_uq_name 
		  do update set internal = excluded.internal,
		  				documentation_url = excluded.documentation_url,
	 					api_specification_type = excluded.api_specification_type,
	 					public_support_contact = excluded.public_support_contact,
	   					tech_support_contact = excluded.tech_support_contact,
						request_costs = excluded.request_costs,
	          			monthly_costs = excluded.monthly_costs,
	         			one_time_costs = excluded.one_time_costs
		      returning id
		),
		availabilities as (
					insert into directory.availabilities 
								(inway_id, service_id, last_announced)
						 select inway.id, service.id, now()
						   from inway, service
					on conflict on constraint availabilities_uq_inway_service
				  do update set last_announced = now(),
								active = true		    
		) select id from service;
		
	`

	return db.PrepareNamed(query)
}

func prepareGetServiceStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		select directory.services.id as id, directory.services.name as name, documentation_url, api_specification_type, internal, tech_support_contact, public_support_contact, directory.organizations.name as organization_name, one_time_costs, monthly_costs, request_costs
		from directory.services
		join directory.organizations 
		    on directory.services.organization_id = directory.organizations.id
		where 
		      directory.services.id = :id;
	`

	return db.PrepareNamed(query)
}

func prepareSelectInwayByAddressStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		SELECT i.id AS inway_id, i.organization_id
		FROM directory.inways i
		INNER JOIN directory.organizations o ON o.id = i.organization_id
		WHERE i.address = :inway_address
		AND o.name = :organization_name
	`

	return db.PrepareNamed(query)
}

func prepareSetOrganizationInwayStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		UPDATE directory.organizations
		SET inway_id = :inway_id
		WHERE id = :organization_id
	`

	return db.PrepareNamed(query)
}

func prepareClearOrganizationInwayStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		UPDATE directory.organizations
		SET inway_id = null
		WHERE name = :name
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
