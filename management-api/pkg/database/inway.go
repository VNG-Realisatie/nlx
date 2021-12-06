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
	return db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoUpdates: clause.AssignmentColumns([]string{"version"}),
		}).
		Create(inway).Error
}

func (db *PostgresConfigDatabase) UpdateInway(ctx context.Context, inway *Inway) error {
	if err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Save(inway).
		Error; err != nil {
		return err
	}

	return nil
}

func (db *PostgresConfigDatabase) DeleteInway(ctx context.Context, name string) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	inway, err := db.GetInway(ctx, name)
	if err != nil {
		return err
	}

	// Load settings for organization inway, and clear it if the to delete inway is set as the organization inway
	settings, err := db.GetSettings(ctx)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return err
	}

	if settings != nil && settings.OrganizationInwayName() != "" && settings.OrganizationInwayName() == name {
		err = tx.
			WithContext(ctx).
			Omit(clause.Associations).
			Model(settings).
			Update("inway_id", nil).Error
		if err != nil {
			return err
		}
	}

	// Remove inway service records
	err = tx.
		WithContext(ctx).
		Exec(
			"DELETE FROM nlx_management.inways_services WHERE inway_id = ?", inway.ID).
		Error
	if err != nil {
		return err
	}

	// Remove the inway from the database
	err = tx.
		WithContext(ctx).
		Where(&Inway{Name: name}).
		Delete(&Inway{}).Error
	if err != nil {
		return err
	}

	return tx.Commit().Error
}

func (db *PostgresConfigDatabase) ListInways(ctx context.Context) ([]*Inway, error) {
	inways := []*Inway{}

	if err := db.DB.
		WithContext(ctx).
		Preload("Services").
		Find(&inways).Error; err != nil {
		return nil, err
	}

	return inways, nil
}

func (db *PostgresConfigDatabase) GetInway(ctx context.Context, name string) (*Inway, error) {
	inway := &Inway{}

	if err := db.DB.
		WithContext(ctx).
		Preload("Services").
		First(inway, Inway{Name: name}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return inway, nil
}
