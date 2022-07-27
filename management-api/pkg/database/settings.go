// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

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
	var inwayID *uint

	var inway = &Inway{}

	if settings.OrganizationInwayName() != "" {
		if err := db.DB.
			WithContext(ctx).
			First(inway, &Inway{Name: settings.OrganizationInwayName()}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrInwayNotFound
			}

			return err
		}

		inwayID = &inway.ID
	}

	settingsInDB := &dbSettings{
		InwayID:                  inwayID,
		Inway:                    inway,
		OrganizationEmailAddress: settings.OrganizationEmailAddress(),
	}

	if err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Where("id IS NOT NULL").
		Assign(map[string]interface{}{"inway_id": inwayID, "organization_email_address": settings.OrganizationEmailAddress()}).
		FirstOrCreate(settingsInDB).Error; err != nil {
		return err
	}

	return nil
}
