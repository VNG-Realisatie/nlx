// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service struct {
	ID                     uint
	Name                   string
	EndpointURL            string
	DocumentationURL       string
	APISpecificationURL    string
	Internal               bool
	TechSupportContact     string
	PublicSupportContact   string
	Inways                 []*Inway `gorm:"many2many:nlx_management.inways_services;"`
	IncomingAccessRequests []*IncomingAccessRequest
	OneTimeCosts           int
	MonthlyCosts           int
	RequestCosts           int
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

var ErrNoIDSpecified = errors.New("unable to update service without a primary key")
var ErrServiceAlreadyExists = errors.New("unable to create service with same name as existing service")

func (s *Service) TableName() string {
	return "nlx_management.services"
}

func (db *PostgresConfigDatabase) ListServices(ctx context.Context) ([]*Service, error) {
	services, err := db.queries.ListServices(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*Service, len(services))

	for i, service := range services {
		inways, err := db.queries.ListInwaysForService(ctx, service.ID)
		if err != nil {
			return nil, err
		}

		inwayModels := make([]*Inway, len(inways))

		for j, inway := range inways {
			inwayModels[j] = &Inway{
				ID:          uint(inway.ID),
				Name:        inway.Name,
				Version:     inway.Version,
				Hostname:    inway.Hostname,
				IPAddress:   inway.IpAddress,
				SelfAddress: inway.SelfAddress,
				CreatedAt:   inway.CreatedAt,
				UpdatedAt:   inway.UpdatedAt,
			}
		}

		result[i] = &Service{
			ID:                     uint(service.ID),
			Name:                   service.Name,
			EndpointURL:            service.EndpointUrl,
			DocumentationURL:       service.DocumentationUrl,
			APISpecificationURL:    service.ApiSpecificationUrl,
			Internal:               service.Internal,
			TechSupportContact:     service.TechSupportContact,
			PublicSupportContact:   service.PublicSupportContact,
			Inways:                 inwayModels,
			IncomingAccessRequests: nil,
			OneTimeCosts:           int(service.OneTimeCosts),
			MonthlyCosts:           int(service.MonthlyCosts),
			RequestCosts:           int(service.RequestCosts),
			CreatedAt:              service.CreatedAt,
			UpdatedAt:              service.UpdatedAt,
		}
	}

	return result, nil
}

func (db *PostgresConfigDatabase) GetService(ctx context.Context, name string) (*Service, error) {
	service := &Service{}
	if err := db.DB.
		WithContext(ctx).
		Preload("Inways").
		First(service, Service{Name: name}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return service, nil
}

func (db *PostgresConfigDatabase) CreateService(ctx context.Context, service *Service) error {
	err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Create(service).Error

	if err == nil {
		return nil
	}

	if err.Error() == `pq: duplicate key value violates unique constraint "services_pkey"` {
		return ErrServiceAlreadyExists
	}

	return err
}

func (db *PostgresConfigDatabase) UpdateService(ctx context.Context, service *Service) error {
	if service.ID == 0 {
		return ErrNoIDSpecified
	}

	return db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Select(
			"endpoint_url",
			"documentation_url",
			"api_specification_url",
			"internal",
			"public_support_contact",
			"tech_support_contact",
			"one_time_costs",
			"monthly_costs",
			"request_costs",
		).
		Save(service).Error
}

func (db *PostgresConfigDatabase) setServiceInways(ctx context.Context, serviceID uint, inwayNames []string) error {
	service := &Service{}
	if err := db.DB.
		WithContext(ctx).
		Where(serviceID).
		First(service).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}

		return err
	}

	inways := []*Inway{}
	if err := db.DB.
		WithContext(ctx).
		Where("name IN ?", inwayNames).
		Find(&inways).Error; err != nil {
		return err
	}

	if len(inways) != len(inwayNames) {
		return ErrNotFound
	}

	return db.DB.Model(service).
		WithContext(ctx).
		Association("Inways").
		Replace(inways)
}

func (db *PostgresConfigDatabase) DeleteService(ctx context.Context, serviceName, organizationSerialNumber string) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}
	service := &Service{}

	err := dbWithTx.
		Where(&Service{Name: serviceName}).
		First(service).
		Error
	if err != nil {
		return err
	}

	err = dbWithTx.Where(&OutgoingAccessRequest{
		ServiceName: serviceName,
		Organization: Organization{
			SerialNumber: organizationSerialNumber,
		},
	}).Delete(&OutgoingAccessRequest{}).Error
	if err != nil {
		dbWithTx.Rollback()
		return err
	}

	err = dbWithTx.DB.
		WithContext(ctx).
		Select(clause.Associations).
		Delete(service).Error
	if err != nil {
		dbWithTx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (db *PostgresConfigDatabase) CreateServiceWithInways(ctx context.Context, service *Service, inwayNames []string) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	if err := dbWithTx.CreateService(ctx, service); err != nil {
		return err
	}

	if err := dbWithTx.setServiceInways(ctx, service.ID, inwayNames); err != nil {
		return err
	}

	return tx.Commit().Error
}

func (db *PostgresConfigDatabase) UpdateServiceWithInways(ctx context.Context, service *Service, inwayNames []string) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	if err := dbWithTx.UpdateService(ctx, service); err != nil {
		return err
	}

	if err := dbWithTx.setServiceInways(ctx, service.ID, inwayNames); err != nil {
		return err
	}

	return tx.Commit().Error
}
