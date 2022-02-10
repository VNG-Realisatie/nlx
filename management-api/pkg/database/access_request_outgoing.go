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

type Organization struct {
	SerialNumber string
	Name         string
}

type OutgoingAccessRequest struct {
	ID                   uint
	Organization         Organization `gorm:"embedded;embeddedPrefix:organization_"`
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
	SynchronizeAt        time.Time
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
			"organization_serial_number = ? AND service_name = ? AND public_key_fingerprint = ? AND state IN ?",
			accessRequest.Organization.SerialNumber,
			accessRequest.ServiceName,
			accessRequest.PublicKeyFingerprint,
			[]string{
				string(OutgoingAccessRequestCreated),
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

func (db *PostgresConfigDatabase) GetLatestOutgoingAccessRequest(ctx context.Context, organizationSerialNumber, serviceName string) (*OutgoingAccessRequest, error) {
	accessRequest := &OutgoingAccessRequest{}

	if err := db.DB.
		WithContext(ctx).
		Where("organization_serial_number = ? AND service_name = ?", organizationSerialNumber, serviceName).
		Order("created_at DESC").
		First(accessRequest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return accessRequest, nil
}

func (db *PostgresConfigDatabase) UpdateOutgoingAccessRequestState(ctx context.Context, accessRequestID uint, state OutgoingAccessRequestState, referenceID uint, schedulerErr *diagnostics.ErrorDetails, synchronizeAt time.Time) error {
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
	outgoingAccessRequest.SynchronizeAt = synchronizeAt

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
			"reference_id",
			"updated_at",
			"error_code",
			"error_cause",
			"error_stack_trace",
			"synchronize_at",
		).
		Save(outgoingAccessRequest).Error
}

func (db *PostgresConfigDatabase) DeleteOutgoingAccessRequests(ctx context.Context, organizationSerialNumber, serviceName string) error {
	return db.DB.
		WithContext(ctx).
		Where("organization_serial_number = ? AND service_name = ?", organizationSerialNumber, serviceName).
		Delete(&OutgoingAccessRequest{}).
		Error
}

func (db *PostgresConfigDatabase) TakePendingOutgoingAccessRequests(ctx context.Context) ([]*OutgoingAccessRequest, error) {
	lockID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	result := db.DB.
		WithContext(ctx).
		Table("nlx_management.access_requests_outgoing").
		Where(
			`id IN (
				SELECT
					access_requests_outgoing.id FROM nlx_management.access_requests_outgoing
				LEFT JOIN
					nlx_management.access_proofs ON (nlx_management.access_proofs.access_request_outgoing_id = nlx_management.access_requests_outgoing.id)
				WHERE
					(state IN ? OR (nlx_management.access_proofs.id IS NOT NULL AND nlx_management.access_proofs.revoked_at IS NULL))
				AND
					(lock_expires_at IS NULL OR NOW() > lock_expires_at)
				AND
					(access_requests_outgoing.synchronize_at < NOW())
				ORDER BY
					updated_at ASC
				)`,
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

	requests := []*OutgoingAccessRequest{}
	if result.RowsAffected > 0 {
		if err := db.DB.
			WithContext(ctx).
			Table("nlx_management.access_requests_outgoing").
			Where("lock_id = ?", lockID).
			Find(&requests).
			Error; err != nil {
			return nil, err
		}
	}

	return requests, nil
}

func (db *PostgresConfigDatabase) UnlockOutgoingAccessRequest(ctx context.Context, outgoingAccessRequest *OutgoingAccessRequest) error {
	return db.DB.
		WithContext(ctx).
		Model(&outgoingAccessRequest).
		Updates(map[string]interface{}{"lock_expires_at": nil, "lock_id": nil}).Error
}
