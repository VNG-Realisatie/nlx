// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/management-api/domain"
)

type dbSettings struct {
	ID                       uint
	InwayID                  *uint
	OrganizationEmailAddress string
	Inway                    *Inway `gorm:"foreignkey:InwayID;references:ID"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

var ErrInwayNotFound = errors.New("inway not found")

func (s *dbSettings) TableName() string {
	return "nlx_management.settings"
}

func (db *PostgresConfigDatabase) GetSettings(ctx context.Context) (*domain.Settings, error) {
	settings, err := db.queries.GetSettings(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewSettings("", "")
		}

		return nil, err
	}

	return domain.NewSettings(
		settings.Name.String,
		settings.OrganizationEmailAddress.String,
	)
}

func (db *PostgresConfigDatabase) UpdateSettings(ctx context.Context, settings *domain.Settings) error {
	if settings.OrganizationInwayName() != "" {
		amount, err := db.queries.CountInwaysByName(ctx, settings.OrganizationInwayName())
		if err != nil {
			return err
		}

		if amount < 1 {
			return ErrInwayNotFound
		}
	}

	return db.queries.UpdateSettings(ctx, &queries.UpdateSettingsParams{
		OrganizationEmailAddress: settings.OrganizationEmailAddress(),
		InwayName:                settings.OrganizationInwayName(),
	})
}
