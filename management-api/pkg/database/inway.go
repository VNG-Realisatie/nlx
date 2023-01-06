// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
)

type Inway struct {
	ID          uint
	Name        string
	Version     string
	Hostname    string
	IPAddress   string
	SelfAddress string
	Services    []*Service `gorm:"many2many:nlx_management.inways_services"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (i *Inway) TableName() string {
	return "nlx_management.inways"
}

func (db *PostgresConfigDatabase) RegisterInway(ctx context.Context, inway *Inway) error {
	return db.queries.UpsertInway(ctx, &queries.UpsertInwayParams{
		Name:        inway.Name,
		SelfAddress: inway.SelfAddress,
		Version:     inway.Version,
		Hostname:    inway.Hostname,
		IpAddress:   inway.IPAddress,
		CreatedAt:   inway.CreatedAt,
		UpdatedAt:   inway.UpdatedAt,
	})
}

func (db *PostgresConfigDatabase) UpdateInway(ctx context.Context, inway *Inway) error {
	err := db.queries.UpdateInway(ctx, &queries.UpdateInwayParams{
		ID:          int32(inway.ID),
		Name:        inway.Name,
		Version:     inway.Version,
		Hostname:    inway.Hostname,
		SelfAddress: inway.SelfAddress,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

var ErrInwayUsedByService = errors.New("inway still in use by at least one service")
var ErrInwayIsOrganizationInway = errors.New("inway used as organization inway")

func (db *PostgresConfigDatabase) DeleteInway(ctx context.Context, name string) error {
	err := db.queries.RemoveInwayByName(ctx, name)
	if err != nil {
		if strings.Contains(err.Error(), "constraint \"fk_inway_services_inway\"") {
			return ErrInwayUsedByService
		}

		if strings.Contains(err.Error(), "constraint \"fk_settings_inway\"") {
			return ErrInwayIsOrganizationInway
		}

		return err
	}

	return nil
}

func (db *PostgresConfigDatabase) ListInways(ctx context.Context) ([]*Inway, error) {
	inwayRows, err := db.queries.ListInways(ctx)
	if err != nil {
		return nil, err
	}

	inways := make([]*Inway, len(inwayRows))

	for i, inwayRow := range inwayRows {
		serviceRows, servicesErr := db.queries.ListServicesForInway(ctx, inwayRow.ID)
		if servicesErr != nil {
			return nil, servicesErr
		}

		services := make([]*Service, len(serviceRows))

		for j, serviceRow := range serviceRows {
			services[j] = &Service{
				ID:                     uint(serviceRow.ID),
				Name:                   serviceRow.Name,
				EndpointURL:            serviceRow.EndpointUrl,
				DocumentationURL:       serviceRow.DocumentationUrl,
				APISpecificationURL:    serviceRow.ApiSpecificationUrl,
				Internal:               serviceRow.Internal,
				TechSupportContact:     serviceRow.TechSupportContact,
				PublicSupportContact:   serviceRow.PublicSupportContact,
				Inways:                 nil,
				IncomingAccessRequests: nil,
				OneTimeCosts:           int(serviceRow.OneTimeCosts),
				MonthlyCosts:           int(serviceRow.MonthlyCosts),
				RequestCosts:           int(serviceRow.RequestCosts),
				CreatedAt:              serviceRow.CreatedAt,
				UpdatedAt:              serviceRow.UpdatedAt,
			}
		}

		inways[i] = &Inway{
			ID:          uint(inwayRow.ID),
			Name:        inwayRow.Name,
			Version:     inwayRow.Version,
			Hostname:    inwayRow.Hostname,
			IPAddress:   inwayRow.IpAddress,
			SelfAddress: inwayRow.SelfAddress,
			Services:    services,
			CreatedAt:   inwayRow.CreatedAt,
			UpdatedAt:   inwayRow.UpdatedAt,
		}
	}

	return inways, nil
}

func (db *PostgresConfigDatabase) GetInway(ctx context.Context, name string) (*Inway, error) {
	inwayRow, err := db.queries.GetInwayByName(ctx, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	serviceRows, servicesErr := db.queries.ListServicesForInway(ctx, inwayRow.ID)
	if servicesErr != nil {
		return nil, servicesErr
	}

	services := make([]*Service, len(serviceRows))

	for j, serviceRow := range serviceRows {
		services[j] = &Service{
			ID:                     uint(serviceRow.ID),
			Name:                   serviceRow.Name,
			EndpointURL:            serviceRow.EndpointUrl,
			DocumentationURL:       serviceRow.DocumentationUrl,
			APISpecificationURL:    serviceRow.ApiSpecificationUrl,
			Internal:               serviceRow.Internal,
			TechSupportContact:     serviceRow.TechSupportContact,
			PublicSupportContact:   serviceRow.PublicSupportContact,
			Inways:                 nil,
			IncomingAccessRequests: nil,
			OneTimeCosts:           int(serviceRow.OneTimeCosts),
			MonthlyCosts:           int(serviceRow.MonthlyCosts),
			RequestCosts:           int(serviceRow.RequestCosts),
			CreatedAt:              serviceRow.CreatedAt,
			UpdatedAt:              serviceRow.UpdatedAt,
		}
	}

	return &Inway{
		ID:          uint(inwayRow.ID),
		Name:        inwayRow.Name,
		Version:     inwayRow.Version,
		Hostname:    inwayRow.Hostname,
		IPAddress:   inwayRow.IpAddress,
		SelfAddress: inwayRow.SelfAddress,
		Services:    services,
		CreatedAt:   inwayRow.CreatedAt,
		UpdatedAt:   inwayRow.UpdatedAt,
	}, nil
}
