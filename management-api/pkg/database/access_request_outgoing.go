// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"go.nlx.io/nlx/common/diagnostics"
)

var ErrActiveAccessRequest = errors.New("there is already an active AccessRequest")

const lockTimeOut = 5 * time.Minute

type OutgoingAccessRequestState string

const (
	OutgoingAccessRequestCreated  OutgoingAccessRequestState = "created"
	OutgoingAccessRequestReceived OutgoingAccessRequestState = "received"
	OutgoingAccessRequestApproved OutgoingAccessRequestState = "approved"
	OutgoingAccessRequestRejected OutgoingAccessRequestState = "rejected"
	OutgoingAccessRequestFailed   OutgoingAccessRequestState = "failed"
)

type OutgoingAccessRequest struct {
	ID                   uint
	OrganizationName     string
	ServiceName          string
	ReferenceID          uint
	State                OutgoingAccessRequestState
	LockID               *uuid.UUID
	LockExpiresAt        sql.NullTime
	PublicKeyFingerprint string
	PublicKeyPEM         string
	ErrorCode            int
	ErrorCause           string
	ErrorStackTrace      pq.StringArray `gorm:"type:text[]"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (*OutgoingAccessRequest) TableName() string {
	return "nlx_management.access_requests_outgoing"
}

func (request *OutgoingAccessRequest) IsSendable() bool {
	return request.State == OutgoingAccessRequestCreated ||
		request.State == OutgoingAccessRequestFailed
}

func (db *PostgresConfigDatabase) CreateOutgoingAccessRequest(ctx context.Context, accessRequest *OutgoingAccessRequest) (*OutgoingAccessRequest, error) {
	var count int64

	if err := db.DB.
		WithContext(ctx).
		Model(OutgoingAccessRequest{}).
		Where(
			"organization_name = ? AND service_name = ? AND public_key_fingerprint = ? AND state IN ?",
			accessRequest.OrganizationName,
			accessRequest.ServiceName,
			accessRequest.PublicKeyFingerprint,
			[]string{
				string(OutgoingAccessRequestCreated),
				string(OutgoingAccessRequestReceived),
				string(OutgoingAccessRequestReceived),
			},
		).
		Count(&count).
		Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, ErrActiveAccessRequest
	}

	if err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Create(accessRequest).Error; err != nil {
		return nil, err
	}

	return accessRequest, nil
}

func (db *PostgresConfigDatabase) GetOutgoingAccessRequest(ctx context.Context, id uint) (*OutgoingAccessRequest, error) {
	accessRequest := &OutgoingAccessRequest{}

	if err := db.DB.
		WithContext(ctx).
		First(accessRequest, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return accessRequest, nil
}

func (db *PostgresConfigDatabase) GetLatestOutgoingAccessRequest(ctx context.Context, organizationName, serviceName string) (*OutgoingAccessRequest, error) {
	accessRequest := &OutgoingAccessRequest{}

	if err := db.DB.
		WithContext(ctx).
		Where("organization_name = ? AND service_name = ?", organizationName, serviceName).
		Order("created_at DESC").
		First(accessRequest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return accessRequest, nil
}

func (db *PostgresConfigDatabase) UpdateOutgoingAccessRequestState(ctx context.Context, accessRequestID uint, state OutgoingAccessRequestState, referenceID uint, schedulerErr *diagnostics.ErrorDetails) error {
	outgoingAccessRequest := &OutgoingAccessRequest{}

	if err := db.DB.
		First(outgoingAccessRequest, accessRequestID).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}

		return err
	}

	if referenceID > 0 {
		outgoingAccessRequest.ReferenceID = referenceID
	}

	outgoingAccessRequest.State = state
	outgoingAccessRequest.LockID = nil
	outgoingAccessRequest.LockExpiresAt = sql.NullTime{}

	if schedulerErr != nil {
		outgoingAccessRequest.ErrorCode = int(schedulerErr.Code)
		outgoingAccessRequest.ErrorCause = schedulerErr.Cause
		outgoingAccessRequest.ErrorStackTrace = schedulerErr.StackTrace
	} else {
		outgoingAccessRequest.ErrorCode = 0
		outgoingAccessRequest.ErrorCause = ""
		outgoingAccessRequest.ErrorStackTrace = nil
	}

	return db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Select(
			"state",
			"lock_id",
			"updated_at",
			"reference_id",
			"lock_expires_at",
			"error_code",
			"error_cause",
			"error_stack_trace",
		).
		Save(outgoingAccessRequest).Error
}

func (db *PostgresConfigDatabase) ListAllOutgoingAccessRequests(ctx context.Context) ([]*OutgoingAccessRequest, error) {
	requests := []*OutgoingAccessRequest{}

	if err := db.DB.
		WithContext(ctx).
		Find(&requests).
		Error; err != nil {
		return nil, err
	}

	return requests, nil
}

func (db *PostgresConfigDatabase) ListOutgoingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*OutgoingAccessRequest, error) {
	requests := []*OutgoingAccessRequest{}

	if err := db.DB.
		WithContext(ctx).
		Where("organization_name = ? AND service_name = ?", organizationName, serviceName).
		Find(&requests).
		Error; err != nil {
		return nil, err
	}

	return requests, nil
}

func (db *PostgresConfigDatabase) TakePendingOutgoingAccessRequest(ctx context.Context) (*OutgoingAccessRequest, error) {
	lockID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	result := db.DB.
		WithContext(ctx).
		Table("nlx_management.access_requests_outgoing").
		Where(
			"id = (SELECT id FROM nlx_management.access_requests_outgoing WHERE state IN ? AND (lock_expires_at IS NULL OR NOW() > lock_expires_at) ORDER BY created_at DESC LIMIT 1)",
			[]string{
				string(OutgoingAccessRequestCreated),
				string(OutgoingAccessRequestApproved),
				string(OutgoingAccessRequestReceived),
			},
		).
		Updates(map[string]interface{}{
			"lock_id":         lockID,
			"lock_expires_at": time.Now().Add(lockTimeOut),
		})
	if result.Error != nil {
		return nil, err
	}

	if result.RowsAffected > 0 {
		request := &OutgoingAccessRequest{}

		if err := db.DB.
			WithContext(ctx).
			Table("nlx_management.access_requests_outgoing").
			Where("lock_id = ?", lockID).
			First(request).
			Error; err != nil {
			return nil, err
		}

		return request, nil
	}

	return nil, nil
}

func (db *PostgresConfigDatabase) UnlockOutgoingAccessRequest(ctx context.Context, outgoingAccessRequest *OutgoingAccessRequest) error {
	return db.DB.
		WithContext(ctx).
		Model(&outgoingAccessRequest).
		Updates(map[string]interface{}{"lock_expires_at": nil}).Error
}
