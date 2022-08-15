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

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
)

var ErrAccessGrantAlreadyRevoked = errors.New("accessGrant is already revoked")

type AccessGrant struct {
	ID                      uint
	IncomingAccessRequestID uint `gorm:"column:access_request_incoming_id;"`
	IncomingAccessRequest   *IncomingAccessRequest
	CreatedAt               time.Time
	RevokedAt               sql.NullTime
}

func (a *AccessGrant) TableName() string {
	return "nlx_management.access_grants"
}

func (db *PostgresConfigDatabase) CreateAccessGrant(ctx context.Context, accessRequest *IncomingAccessRequest) (*AccessGrant, error) {
	result := &AccessGrant{
		IncomingAccessRequestID: accessRequest.ID,
		CreatedAt:               time.Now(),
	}

	id, err := db.queries.CreateAccessGrant(ctx, &queries.CreateAccessGrantParams{
		AccessRequestIncomingID: int32(result.IncomingAccessRequestID),
		CreatedAt:               result.CreatedAt,
	})
	if err != nil {
		return nil, err
	}

	result.ID = uint(id)

	return result, nil
}

func (db *PostgresConfigDatabase) GetAccessGrant(ctx context.Context, id uint) (*AccessGrant, error) {
	accessGrant := &AccessGrant{}

	if err := db.DB.
		WithContext(ctx).
		Preload("IncomingAccessRequest").
		Preload("IncomingAccessRequest.Service").
		First(accessGrant, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return accessGrant, nil
}

//nolint:dupl // looks the same as RevokeAccessProof but is different. RevokeAccessGrant is for access grants RevokeAccessProof is for access proofs.
func (db *PostgresConfigDatabase) RevokeAccessGrant(ctx context.Context, accessGrantID uint, revokedAt time.Time) (*AccessGrant, error) {
	accessGrant := &AccessGrant{}

	if err := db.DB.
		WithContext(ctx).
		Preload("IncomingAccessRequest").
		Preload("IncomingAccessRequest.Service").
		First(accessGrant, accessGrantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	if accessGrant.RevokedAt.Valid {
		return nil, ErrAccessGrantAlreadyRevoked
	}

	accessGrant.RevokedAt = sql.NullTime{
		Time:  revokedAt,
		Valid: true,
	}

	if err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			WithContext(ctx).
			Omit(clause.Associations).
			Select("revoked_at").
			Save(accessGrant).Error; err != nil {
			return err
		}

		accessGrant.IncomingAccessRequest.State = IncomingAccessRequestRevoked

		if err := tx.
			WithContext(ctx).
			Omit(clause.Associations).
			Select("state", "updated_at").
			Save(accessGrant.IncomingAccessRequest).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return accessGrant, nil
}

func (db *PostgresConfigDatabase) ListAccessGrantsForService(ctx context.Context, serviceName string) ([]*AccessGrant, error) {
	accessGrants := []*AccessGrant{}

	if err := db.DB.
		WithContext(ctx).
		Preload("IncomingAccessRequest").
		Preload("IncomingAccessRequest.Service").
		Joins("JOIN nlx_management.access_requests_incoming r ON r.id = access_grants.access_request_incoming_id").
		Joins("JOIN nlx_management.services s ON s.id = r.service_id AND s.name = ?", serviceName).
		Find(&accessGrants).Error; err != nil {
		return nil, err
	}

	return accessGrants, nil
}

func (db *PostgresConfigDatabase) GetLatestAccessGrantForService(ctx context.Context, organizationSerialNumber, serviceName, publicKeyFingerprint string) (*AccessGrant, error) {
	accessGrant := &AccessGrant{}

	if err := db.DB.
		WithContext(ctx).
		Preload("IncomingAccessRequest").
		Preload("IncomingAccessRequest.Service").
		Joins("JOIN nlx_management.access_requests_incoming r ON r.id = access_grants.access_request_incoming_id AND r.organization_serial_number = ? AND r.public_key_fingerprint = ?", organizationSerialNumber, publicKeyFingerprint).
		Joins("JOIN nlx_management.services s ON s.id = r.service_id AND s.name = ?", serviceName).
		Order("created_at DESC").
		First(accessGrant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return accessGrant, nil
}
