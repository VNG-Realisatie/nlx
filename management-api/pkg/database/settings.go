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
	ID            uint `gorm:"primarykey;column:settings_id"`
	IrmaServerURL string
	InsightAPIURL string
	InwayID       *uint
	Inway         *Inway `gorm:"foreignkey:InwayID;references:ID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
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

func (db *PostgresConfigDatabase) UpdateSettings(ctx context.Context, irmaServerURL, insightAPIURL string, inwayID *uint) (*Settings, error) {
	settingsInDB := &Settings{}
	if err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Where("settings_id IS NOT NULL").
		Assign(&Settings{
			InwayID:       inwayID,
			IrmaServerURL: irmaServerURL,
			InsightAPIURL: insightAPIURL,
		}).
		FirstOrCreate(settingsInDB).Error; err != nil {
		return nil, err
	}

	return settingsInDB, nil
}

func (db *PostgresConfigDatabase) PutInsightConfiguration(ctx context.Context, irmaServerURL, insightAPIURL string) (*Settings, error) {
	settingsInDB := &Settings{}
	if err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Where("settings_id IS NOT NULL").
		Assign(map[string]interface{}{"insight_api_url": insightAPIURL, "irma_server_url": irmaServerURL}).
		FirstOrCreate(settingsInDB).Error; err != nil {
		return nil, err
	}

	return settingsInDB, nil
}
