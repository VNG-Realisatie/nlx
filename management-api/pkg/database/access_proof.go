// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"time"

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

// nolint:dupl // migration from Gorm to Sqlc
func (db *PostgresConfigDatabase) GetAccessProofForOutgoingAccessRequest(ctx context.Context, accessRequestID uint) (*AccessProof, error) {
	outgoingAccessRequest, err := db.queries.GetOutgoingAccessRequest(ctx, int32(accessRequestID))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	accessProof, err := db.queries.GetAccessProofByOutgoingAccessRequest(ctx, outgoingAccessRequest.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	var errorCause = ""

	if outgoingAccessRequest.ErrorCause.Valid {
		errorCause = outgoingAccessRequest.ErrorCause.String
	}

	result := &AccessProof{
		ID:                      uint(accessProof.ID),
		AccessRequestOutgoingID: uint(outgoingAccessRequest.ID),
		OutgoingAccessRequest: &OutgoingAccessRequest{
			ID: uint(outgoingAccessRequest.ID),
			Organization: Organization{
				SerialNumber: outgoingAccessRequest.OrganizationSerialNumber,
				Name:         outgoingAccessRequest.OrganizationName,
			},
			ServiceName:          outgoingAccessRequest.ServiceName,
			ReferenceID:          uint(outgoingAccessRequest.ReferenceID),
			State:                OutgoingAccessRequestState(outgoingAccessRequest.State),
			PublicKeyFingerprint: outgoingAccessRequest.PublicKeyFingerprint,
			PublicKeyPEM:         outgoingAccessRequest.PublicKeyPem,
			ErrorCode:            int(outgoingAccessRequest.ErrorCode),
			ErrorCause:           errorCause,
			CreatedAt:            outgoingAccessRequest.CreatedAt,
			UpdatedAt:            outgoingAccessRequest.UpdatedAt,
		},
		CreatedAt:    accessProof.CreatedAt,
		RevokedAt:    accessProof.RevokedAt,
		TerminatedAt: accessProof.TerminatedAt,
	}

	return result, nil
}

func (db *PostgresConfigDatabase) RevokeAccessProof(ctx context.Context, accessProofID uint, revokedAt time.Time) (*AccessProof, error) {
	accessProof, err := db.queries.GetAccessProof(ctx, int32(accessProofID))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	err = db.queries.RevokeAccessProof(ctx, &queries.RevokeAccessProofParams{
		ID: int32(accessProofID),
		RevokedAt: sql.NullTime{
			Valid: true,
			Time:  revokedAt,
		},
	})
	if err != nil {
		return nil, err
	}

	result := &AccessProof{
		ID:                      uint(accessProof.ID),
		AccessRequestOutgoingID: uint(accessProof.AccessRequestOutgoingID),
		CreatedAt:               accessProof.CreatedAt,
		RevokedAt: sql.NullTime{
			Valid: true,
			Time:  revokedAt,
		},
	}

	return result, nil
}

// nolint:dupl // migration from Gorm to Sqlc
func (db *PostgresConfigDatabase) GetAccessProofs(ctx context.Context, accessProofIDs []uint64) ([]*AccessProof, error) {
	result := []*AccessProof{}

	for _, id := range accessProofIDs {
		accessProof, err := db.queries.GetAccessProof(ctx, int32(id))
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, ErrNotFound
			}

			return nil, err
		}

		outgoingAccessRequest, err := db.queries.GetOutgoingAccessRequest(ctx, accessProof.AccessRequestOutgoingID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, ErrNotFound
			}

			return nil, err
		}

		var errorCause = ""

		if outgoingAccessRequest.ErrorCause.Valid {
			errorCause = outgoingAccessRequest.ErrorCause.String
		}

		model := &AccessProof{
			ID:                      uint(accessProof.ID),
			AccessRequestOutgoingID: uint(outgoingAccessRequest.ID),
			OutgoingAccessRequest: &OutgoingAccessRequest{
				ID: uint(outgoingAccessRequest.ID),
				Organization: Organization{
					SerialNumber: outgoingAccessRequest.OrganizationSerialNumber,
					Name:         outgoingAccessRequest.OrganizationName,
				},
				ServiceName:          outgoingAccessRequest.ServiceName,
				ReferenceID:          uint(outgoingAccessRequest.ReferenceID),
				State:                OutgoingAccessRequestState(outgoingAccessRequest.State),
				PublicKeyFingerprint: outgoingAccessRequest.PublicKeyFingerprint,
				PublicKeyPEM:         outgoingAccessRequest.PublicKeyPem,
				ErrorCode:            int(outgoingAccessRequest.ErrorCode),
				ErrorCause:           errorCause,
				CreatedAt:            outgoingAccessRequest.CreatedAt,
				UpdatedAt:            outgoingAccessRequest.UpdatedAt,
			},
			CreatedAt:    accessProof.CreatedAt,
			RevokedAt:    accessProof.RevokedAt,
			TerminatedAt: accessProof.TerminatedAt,
		}

		result = append(result, model)
	}

	return result, nil
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
