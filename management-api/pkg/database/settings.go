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

type Settings struct {
	ID        uint
	InwayID   *uint
	Inway     *Inway `gorm:"foreignkey:InwayID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var ErrInwayNotFound = errors.New("inway not found")

func (s *Settings) TableName() string {
	return "nlx_management.settings"
}

func (db *PostgresConfigDatabase) GetSettings(ctx context.Context) (*Settings, error) {
	organizationSettings := &Settings{}

	if err := db.DB.
		WithContext(ctx).
		Preload("Inway").
		First(organizationSettings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return organizationSettings, nil
}

func (db *PostgresConfigDatabase) PutOrganizationInway(ctx context.Context, inwayID *uint) (*Settings, error) {
	settingsInDB := &Settings{}
	err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Where("id IS NOT NULL").
		Assign(map[string]interface{}{"inway_id": inwayID}).
		FirstOrCreate(settingsInDB).Error

	if err != nil {
		if err.Error() == `pq: insert or update on table "settings" violates foreign key constraint "fk_organization_settings_inway"` {
			return nil, ErrInwayNotFound
		}

		return nil, err
	}

	return settingsInDB, nil
}
