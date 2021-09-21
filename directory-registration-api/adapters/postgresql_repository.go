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
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-registration-api/domain"
)

var (
	ErrDuplicateAddress     = errors.New("another inway is already registered with this address")
	ErrNoInwayWithAddress   = errors.New("no inway found for address")
	ErrOrganizationNotFound = errors.New("no organization found")
)

type PostgreSQLRepository struct {
	logger                             *zap.Logger
	db                                 *sqlx.DB
	registerInwayStmt                  *sqlx.NamedStmt
	getInwayStmt                       *sqlx.NamedStmt
	registerServiceStmt                *sqlx.NamedStmt
	getServiceStmt                     *sqlx.NamedStmt
	selectInwayByAddressStmt           *sqlx.NamedStmt
	setOrganizationInwayStmt           *sqlx.NamedStmt
	clearOrganizationInwayStmt         *sqlx.NamedStmt
	selectOrganizationInwayAddressStmt *sqlx.NamedStmt
}

//nolint gocyclo: all checks in this function are necessary
func New(logger *zap.Logger, db *sqlx.DB) (*PostgreSQLRepository, error) {
	if logger == nil {
		panic("missing logger")
	}

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

	selectOrganizationInwayAddressStmt, err := prepareSelectOrganizationInwayAddressStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select organization inway address statement: %s", err)
	}

	return &PostgreSQLRepository{
		logger:                             logger.Named("postgres repository"),
		db:                                 db,
		registerInwayStmt:                  registerInwayStmt,
		getInwayStmt:                       getInwayStmt,
		registerServiceStmt:                registerServiceStmt,
		getServiceStmt:                     getServiceStmt,
		selectInwayByAddressStmt:           selectInwayByAddressStmt,
		setOrganizationInwayStmt:           setOrganizationInwayStmt,
		clearOrganizationInwayStmt:         clearOrganizationInwayStmt,
		selectOrganizationInwayAddressStmt: selectOrganizationInwayAddressStmt,
	}, nil
}

func (r *PostgreSQLRepository) RegisterInway(model *domain.Inway) error {
	type registerParams struct {
		OrganizationName string    `db:"organization_name"`
		SerialNumber     string    `db:"serial_number"`
		Name             string    `db:"inway_name"`
		Address          string    `db:"inway_address"`
		NlxVersion       string    `db:"inway_version"`
		CreatedAt        time.Time `db:"inway_created_at"`
		UpdatedAt        time.Time `db:"inway_updated_at"`
	}

	organization := model.Organization()

	_, err := r.registerInwayStmt.Exec(&registerParams{
		OrganizationName: organization.Name(),
		SerialNumber:     organization.SerialNumber(),
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

// @TODO: Remove name, only use serial number
func (r *PostgreSQLRepository) GetInway(name, serialNumber string) (*domain.Inway, error) {
	type params struct {
		Name         string `db:"name"`
		SerialNumber string `db:"serial_number"`
	}

	type dbInway struct {
		Name                     string    `db:"name"`
		Address                  string    `db:"address"`
		NlxVersion               string    `db:"nlx_version"`
		OrganizationName         string    `db:"organization_name"`
		OrganizationSerialNumber string    `db:"serial_number"`
		CreatedAt                time.Time `db:"created_at"`
		UpdatedAt                time.Time `db:"updated_at"`
	}

	result := dbInway{}

	err := r.getInwayStmt.Get(&result, &params{
		Name:         name,
		SerialNumber: serialNumber,
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

func (r *PostgreSQLRepository) RegisterService(model *domain.Service) error {
	type registerParams struct {
		Name                 string `db:"name"`
		SerialNumber         string `db:"serial_number"`
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
			SerialNumber:         model.SerialNumber(),
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

func (r *PostgreSQLRepository) GetService(id uint) (*domain.Service, error) {
	type params struct {
		ID uint `db:"id"`
	}

	type dbService struct {
		ID                   uint   `db:"id"`
		Name                 string `db:"name"`
		SerialNumber         string `db:"serial_number"`
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

	model, err := domain.NewService(
		&domain.NewServiceArgs{
			Name:                 result.Name,
			SerialNumber:         result.SerialNumber,
			Internal:             result.Internal,
			DocumentationURL:     result.DocumentationURL,
			APISpecificationType: domain.SpecificationType(result.APISpecificationType),
			PublicSupportContact: result.PublicSupportContact,
			TechSupportContact:   result.TechSupportContact,
			OneTimeCosts:         uint(result.OneTimeCosts),
			MonthlyCosts:         uint(result.MonthlyCosts),
			RequestCosts:         uint(result.RequestCosts),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid inway model in database: %v", err)
	}

	model.SetID(result.ID)

	return model, nil
}

func (r *PostgreSQLRepository) SetOrganizationInway(ctx context.Context, organizationSerialNumber, inwayAddress string) error {
	arg := map[string]interface{}{
		"inway_address": inwayAddress,
		"serial_number": organizationSerialNumber,
	}

	var ioID struct {
		InwayID        int `db:"inway_id"`
		OrganizationID int `db:"organization_id"`
	}

	err := r.selectInwayByAddressStmt.GetContext(ctx, &ioID, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoInwayWithAddress
		}

		return err
	}

	setOrgInwayArgs := map[string]interface{}{
		"inway_id":      ioID.InwayID,
		"serial_number": organizationSerialNumber,
	}

	_, err = r.setOrganizationInwayStmt.ExecContext(ctx, setOrgInwayArgs)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLRepository) ClearOrganizationInway(ctx context.Context, serialNumber string) error {
	arg := map[string]interface{}{
		"serial_number": serialNumber,
	}

	res, err := r.clearOrganizationInwayStmt.ExecContext(ctx, arg)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n != 1 {
		return ErrOrganizationNotFound
	}

	return nil
}

func (r *PostgreSQLRepository) GetOrganizationInwayAddress(ctx context.Context, serialNumber string) (string, error) {
	var address sql.NullString

	arg := map[string]interface{}{
		"serial_number": serialNumber,
	}

	err := r.selectOrganizationInwayAddressStmt.GetContext(ctx, &address, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrOrganizationNotFound
		}

		return "", err
	}

	return address.String, nil
}

// ClearIfSetAsOrganizationInway clears the inway for the given organization.
// This method should be called if IsOrganizationInway is false in the request, to ensure the directory has this correctly set as well
func (r *PostgreSQLRepository) ClearIfSetAsOrganizationInway(ctx context.Context, serialNumber, selfAddress string) error {
	organizationSelfAddress, err := r.GetOrganizationInwayAddress(ctx, serialNumber)
	if err != nil {
		if errors.Is(err, ErrOrganizationNotFound) {
			return nil
		}

		return err
	}

	if selfAddress == organizationSelfAddress {
		r.logger.Warn("unexpected state: inway was incorrectly set as organization inway ", zap.String("inway self address", selfAddress), zap.String("organization inway self address", organizationSelfAddress))
		return r.ClearOrganizationInway(ctx, serialNumber)
	}

	return nil
}

func prepareRegisterInwayStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		with organization as (
			insert into directory.organizations
			            (serial_number, name)
			     values (:serial_number, :organization_name)
			on conflict
			    		on constraint organizations_uq_serial_number
			  			do update
			      			set serial_number = excluded.serial_number, -- no-op update to return id
			      			    name 		  = excluded.name
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
								  updated_at = excluded.updated_at;
	`

	return db.PrepareNamed(query)
}

func prepareGetInwayStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		select directory.inways.name as name, address, version as nlx_version, created_at, updated_at, directory.organizations.serial_number, directory.organizations.name as organization_name
		from directory.inways
		join directory.organizations
		    on directory.inways.organization_id = directory.organizations.id
		where
		      directory.organizations.serial_number = :serial_number
		  and directory.inways.name = :name;
	`

	return db.PrepareNamed(query)
}

func prepareRegisterServiceStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		with organization as (
		    select id from directory.organizations where directory.organizations.serial_number = :serial_number
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
		select directory.services.id as id, directory.services.name as name, documentation_url, api_specification_type, internal, tech_support_contact, public_support_contact, directory.organizations.serial_number as serial_number, one_time_costs, monthly_costs, request_costs
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
		AND o.serial_number = :serial_number
	`

	return db.PrepareNamed(query)
}

func prepareSetOrganizationInwayStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		UPDATE directory.organizations
		SET inway_id = :inway_id
		WHERE serial_number = :serial_number
	`

	return db.PrepareNamed(query)
}

func prepareClearOrganizationInwayStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		UPDATE directory.organizations
		SET inway_id = null
		WHERE serial_number = :serial_number
	`

	return db.PrepareNamed(query)
}

func prepareSelectOrganizationInwayAddressStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		SELECT i.address
		FROM directory.organizations o
		LEFT JOIN directory.inways i ON o.inway_id = i.id
		WHERE o.serial_number = :serial_number
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
