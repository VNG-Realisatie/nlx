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
)

var ErrAccessGrantAlreadyRevoked = errors.New("accessGrant is already revoked")

type AccessGrant struct {
	ID                      uint `gorm:"primarykey;column:access_grant_id;"`
	IncomingAccessRequestID uint `gorm:"column:access_request_incoming_id;"`
	IncomingAccessRequest   *IncomingAccessRequest
	CreatedAt               time.Time
	RevokedAt               sql.NullTime
}

func (a *AccessGrant) TableName() string {
	return "nlx_management.access_grants"
}

func (db *PostgresConfigDatabase) CreateAccessGrant(ctx context.Context, accessRequest *IncomingAccessRequest) (*AccessGrant, error) {
	accessGrant := &AccessGrant{
		IncomingAccessRequestID: accessRequest.ID,
	}

	if err := db.DB.Debug().
		WithContext(ctx).
		Omit(clause.Associations).
		Create(accessGrant).Error; err != nil {
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

	if err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Select("revoked_at").
		Save(accessGrant).Error; err != nil {
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
		Joins("JOIN nlx_management.access_requests_incoming r ON r.access_request_incoming_id = access_grants.access_request_incoming_id").
		Joins("JOIN nlx_management.services s ON s.service_id = r.service_id AND s.name = ?", serviceName).
		Find(&accessGrants).Error; err != nil {
		return nil, err
	}

	return accessGrants, nil
}

func (db *PostgresConfigDatabase) GetLatestAccessGrantForService(ctx context.Context, organizationName, serviceName string) (*AccessGrant, error) {
	accessGrant := &AccessGrant{}

	if err := db.DB.
		WithContext(ctx).
		Preload("IncomingAccessRequest").
		Preload("IncomingAccessRequest.Service").
		Joins("JOIN nlx_management.access_requests_incoming r ON r.access_request_incoming_id = access_grants.access_request_incoming_id AND r.organization_name = ?", organizationName).
		Joins("JOIN nlx_management.services s ON s.service_id = r.service_id AND s.name = ?", serviceName).
		Order("created_at DESC").
		First(accessGrant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return accessGrant, nil
}
