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
	ID        uint `gorm:"primarykey;column:settings_id"`
	InwayID   *uint
	Inway     *Inway `gorm:"foreignkey:InwayID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

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
	if err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Where("settings_id IS NOT NULL").
		Assign(map[string]interface{}{"inway_id": inwayID}).
		FirstOrCreate(settingsInDB).Error; err != nil {
		return nil, err
	}

	return settingsInDB, nil
}
