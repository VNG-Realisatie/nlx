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

type AccessProof struct {
	ID                      uint `gorm:"primarykey;column:access_proof_id"`
	AccessRequestOutgoingID uint
	OutgoingAccessRequest   *OutgoingAccessRequest `gorm:"foreignKey:access_request_outgoing_id"`
	CreatedAt               time.Time
	RevokedAt               sql.NullTime
}

func (AccessProof) TableName() string {
	return "nlx_management.access_proofs"
}

func (db *PostgresConfigDatabase) CreateAccessProof(ctx context.Context, accessRequest *OutgoingAccessRequest) (*AccessProof, error) {
	accessProof := &AccessProof{
		AccessRequestOutgoingID: accessRequest.ID,
		OutgoingAccessRequest:   accessRequest,
	}

	result := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Create(accessProof)
	if result.Error != nil {
		return nil, result.Error
	}

	return accessProof, nil
}

func (db *PostgresConfigDatabase) GetLatestAccessProofForService(ctx context.Context, organizationName, serviceName string) (*AccessProof, error) {
	accessProof := &AccessProof{}

	if err := db.DB.
		WithContext(ctx).
		Preload("OutgoingAccessRequest").
		Joins("JOIN nlx_management.access_requests_outgoing r ON r.organization_name = ? AND r.service_name = ?", organizationName, serviceName).
		Order("created_at DESC").
		First(accessProof).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	return accessProof, nil
}

func (db *PostgresConfigDatabase) GetAccessProofForOutgoingAccessRequest(ctx context.Context, accessRequestID uint) (*AccessProof, error) {
	accessProof := &AccessProof{}

	if err := db.DB.
		WithContext(ctx).
		Preload("OutgoingAccessRequest").
		Where("access_request_outgoing_id = ?", accessRequestID).
		First(accessProof).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return accessProof, nil
}

//nolint:dupl // looks the same as RevokeAccessGrant but is different. RevokeAccessGrant is for access grants RevokeAccessProof is for access proofs.
func (db *PostgresConfigDatabase) RevokeAccessProof(ctx context.Context, accessProofID uint, revokedAt time.Time) (*AccessProof, error) {
	accessProof := &AccessProof{}

	if err := db.DB.
		WithContext(ctx).
		First(accessProof, accessProofID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	accessProof.RevokedAt = sql.NullTime{
		Time:  revokedAt,
		Valid: true,
	}

	if err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Select("revoked_at").
		Save(accessProof).Error; err != nil {
		return nil, err
	}

	return accessProof, nil
}
