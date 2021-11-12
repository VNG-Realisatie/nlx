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

	// Delete services_inway entries for the inway
	err = db.DeleteServicesConnectedToInway(ctx, inway)
	if err != nil {
		return err
	}

	// Load settings for organization inway, and clear it if the to delete inway is set as the organization inway
	settings, err := db.GetSettings(ctx)
	if err != nil {
		return err
	}

	if settings.Inway != nil && settings.Inway.ID == inway.ID {
		_, err = db.PutOrganizationInway(ctx, nil)
		if err != nil {
			return err
		}
	}

	// Remove the inway from the database
	err = db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
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
