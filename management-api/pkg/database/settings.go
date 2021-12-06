// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
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
	organizationSettings := &dbSettings{}

	if err := db.DB.
		WithContext(ctx).
		Preload("Inway").
		First(organizationSettings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	inwayName := ""
	if organizationSettings.Inway != nil {
		inwayName = organizationSettings.Inway.Name
	}

	settings, err := domain.NewSettings(inwayName, organizationSettings.OrganizationEmailAddress)
	if err != nil {
		return nil, err
	}

	return settings, nil
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
