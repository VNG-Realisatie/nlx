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

type AccessProof struct {
	ID                      uint
	AccessRequestOutgoingID uint
	OutgoingAccessRequest   *OutgoingAccessRequest `gorm:"foreignKey:access_request_outgoing_id"`
	CreatedAt               time.Time
	RevokedAt               sql.NullTime
	TerminatedAt            sql.NullTime
}

func (AccessProof) TableName() string {
	return "nlx_management.access_proofs"
}

func (db *PostgresConfigDatabase) CreateAccessProof(ctx context.Context, accessRequestOutgoingID uint) (*AccessProof, error) {
	result := &AccessProof{
		AccessRequestOutgoingID: accessRequestOutgoingID,
		CreatedAt:               time.Now(),
	}

	id, err := db.queries.CreateAccessProof(ctx, &queries.CreateAccessProofParams{
		AccessRequestOutgoingID: int32(result.AccessRequestOutgoingID),
		CreatedAt:               result.CreatedAt,
	})
	if err != nil {
		return nil, err
	}

	result.ID = uint(id)

	return result, nil
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

func (db *PostgresConfigDatabase) GetAccessProofs(ctx context.Context, accessProofIDs []uint64) ([]*AccessProof, error) {
	accessProofs := &[]*AccessProof{}

	if err := db.DB.
		WithContext(ctx).
		Preload("OutgoingAccessRequest").
		Where("id IN ?", accessProofIDs).
		Find(accessProofs).Error; err != nil {
		return nil, err
	}

	return *accessProofs, nil
}

// nolint:dupl // is similar to terminate access grant
func (db *PostgresConfigDatabase) TerminateAccessProof(ctx context.Context, accessProofID uint, terminatedAt time.Time) error {
	accessProof, err := db.queries.GetAccessProof(ctx, int32(accessProofID))
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}

		return err
	}

	if accessProof.TerminatedAt.Valid {
		return ErrAccessProofAlreadyTerminated
	}

	err = db.queries.TerminateAccessProof(ctx, &queries.TerminateAccessProofParams{
		ID: int32(accessProofID),
		TerminatedAt: sql.NullTime{
			Time:  terminatedAt,
			Valid: true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
