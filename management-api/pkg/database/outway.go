// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Outway struct {
	ID                   uint
	Name                 string
	IPAddress            pgtype.Inet `gorm:"type:inet"`
	SelfAddressAPI       string
	PublicKeyPEM         string
	PublicKeyFingerprint string
	Version              string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (i *Outway) TableName() string {
	return "nlx_management.outways"
}

func (db *PostgresConfigDatabase) RegisterOutway(ctx context.Context, outway *Outway) error {
	return db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoUpdates: clause.AssignmentColumns([]string{"version", "ip_address", "self_address_api", "public_key_pem", "public_key_fingerprint"}),
		}).
		Create(outway).Error
}

func (db *PostgresConfigDatabase) ListOutways(ctx context.Context) ([]*Outway, error) {
	outways := []*Outway{}

	if err := db.DB.
		WithContext(ctx).
		Find(&outways).Error; err != nil {
		return nil, err
	}

	return outways, nil
}

func (db *PostgresConfigDatabase) GetOutway(ctx context.Context, name string) (*Outway, error) {
	outway := &Outway{}

	if err := db.DB.
		WithContext(ctx).
		First(outway, Outway{Name: name}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return outway, nil
}

func (db *PostgresConfigDatabase) GetOutwaysByPublicKeyFingerprint(ctx context.Context, publicKeyFingerprint string) ([]*Outway, error) {
	outways := []*Outway{}

	if err := db.DB.
		WithContext(ctx).
		Where("public_key_fingerprint = ?", publicKeyFingerprint).
		Find(&outways).Error; err != nil {
		return nil, err
	}

	if len(outways) == 0 {
		return nil, ErrNotFound
	}

	return outways, nil
}
