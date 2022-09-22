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
	return request.State == OutgoingAccessRequestFailed
}

func (db *PostgresConfigDatabase) ListLatestOutgoingAccessRequests(ctx context.Context, organizationSerialNumber, serviceName string) ([]*OutgoingAccessRequest, error) {
	outgoingAccessRequests := &[]*OutgoingAccessRequest{}

	if err := db.DB.
		Raw(`
			SELECT
				distinct on (public_key_fingerprint, service_name, organization_serial_number) nlx_management.access_requests_outgoing.*
			FROM
				nlx_management.access_requests_outgoing
			WHERE
				organization_serial_number = ? AND service_name = ?
			ORDER BY
				organization_serial_number, public_key_fingerprint, service_name, created_at DESC;`, organizationSerialNumber, serviceName).Scan(outgoingAccessRequests).Error; err != nil {
		return nil, err
	}

	return *outgoingAccessRequests, nil
}

func (db *PostgresConfigDatabase) ListAllLatestOutgoingAccessRequests(ctx context.Context) ([]*OutgoingAccessRequest, error) {
	accessRequests, err := db.queries.ListAllLatestOutgoingAccessRequests(ctx)
	if err != nil {
		return nil, err
	}

	var outgoingAccessRequests = make([]*OutgoingAccessRequest, len(accessRequests))

	for i, accessRequest := range accessRequests {
		var lockID *uuid.UUID = nil

		if accessRequest.LockID.Valid {
			lockID = &accessRequest.LockID.UUID
		}

		var pem = ""

		if accessRequest.PublicKeyPem.Valid {
			pem = accessRequest.PublicKeyPem.String
		}

		var errorCause = ""

		if accessRequest.ErrorCause.Valid {
			errorCause = accessRequest.ErrorCause.String
		}

		outgoingAccessRequests[i] = &OutgoingAccessRequest{
			ID: uint(accessRequest.ID),
			Organization: Organization{
				SerialNumber: accessRequest.OrganizationSerialNumber,
				Name:         accessRequest.OrganizationName,
			},
			ServiceName:          accessRequest.ServiceName,
			ReferenceID:          uint(accessRequest.ReferenceID),
			State:                OutgoingAccessRequestState(accessRequest.State),
			LockID:               lockID,
			LockExpiresAt:        accessRequest.LockExpiresAt,
			PublicKeyFingerprint: accessRequest.PublicKeyFingerprint,
			PublicKeyPEM:         pem,
			ErrorCode:            int(accessRequest.ErrorCode),
			ErrorCause:           errorCause,
			CreatedAt:            accessRequest.CreatedAt,
			UpdatedAt:            accessRequest.UpdatedAt,
			SynchronizeAt:        accessRequest.SynchronizeAt,
		}
	}

	return outgoingAccessRequests, nil
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

func (db *PostgresConfigDatabase) GetLatestOutgoingAccessRequest(ctx context.Context, organizationSerialNumber, serviceName, publicKeyFingerprint string) (*OutgoingAccessRequest, error) {
	accessRequest := &OutgoingAccessRequest{}

	if err := db.DB.
		WithContext(ctx).
		Where("organization_serial_number = ? AND service_name = ? AND public_key_fingerprint = ?", organizationSerialNumber, serviceName, publicKeyFingerprint).
		Order("created_at DESC").
		First(accessRequest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return accessRequest, nil
}

func (db *PostgresConfigDatabase) GetLatestOutgoingAccessRequestsPerCertificate(ctx context.Context, organizationSerialNumber, serviceName, publicKeyFingerprint string) ([]*OutgoingAccessRequest, error) {
	accessRequests := []*OutgoingAccessRequest{}

	if err := db.DB.
		WithContext(ctx).
		Find(accessRequests).
		Where("organization_serial_number = ? AND service_name = ? AND public_key_fingerprint", organizationSerialNumber, serviceName, publicKeyFingerprint).
		Order("created_at DESC").Error; err != nil {
		return nil, err
	}

	return accessRequests, nil
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

func (db *PostgresConfigDatabase) UnlockOutgoingAccessRequest(ctx context.Context, outgoingAccessRequest *OutgoingAccessRequest) error {
	return db.DB.
		WithContext(ctx).
		Model(&outgoingAccessRequest).
		Updates(map[string]interface{}{"lock_expires_at": nil, "lock_id": nil}).Error
}
