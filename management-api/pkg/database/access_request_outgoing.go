// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
)

var ErrActiveAccessRequest = errors.New("there is already an active AccessRequest")

type OutgoingAccessRequestState string

const (
	OutgoingAccessRequestReceived  OutgoingAccessRequestState = "received"
	OutgoingAccessRequestApproved  OutgoingAccessRequestState = "approved"
	OutgoingAccessRequestRejected  OutgoingAccessRequestState = "rejected"
	OutgoingAccessRequestFailed    OutgoingAccessRequestState = "failed"
	OutgoingAccessRequestWithdrawn OutgoingAccessRequestState = "withdrawn"
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
	return request.State == OutgoingAccessRequestFailed
}

func (db *PostgresConfigDatabase) ListLatestOutgoingAccessRequests(_ context.Context, organizationSerialNumber, serviceName string) ([]*OutgoingAccessRequest, error) {
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
			PublicKeyFingerprint: accessRequest.PublicKeyFingerprint,
			PublicKeyPEM:         pem,
			ErrorCode:            int(accessRequest.ErrorCode),
			ErrorCause:           errorCause,
			CreatedAt:            accessRequest.CreatedAt,
			UpdatedAt:            accessRequest.UpdatedAt,
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

func (db *PostgresConfigDatabase) UpdateOutgoingAccessRequestState(ctx context.Context, accessRequestID uint, state OutgoingAccessRequestState) error {
	outgoingAccessRequest := &OutgoingAccessRequest{}

	if err := db.DB.
		First(outgoingAccessRequest, accessRequestID).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}

		return err
	}

	return db.queries.UpdateOutgoingAccessRequestState(ctx, &queries.UpdateOutgoingAccessRequestStateParams{
		State:     string(state),
		UpdatedAt: time.Now(),
		ID:        int32(accessRequestID),
	})
}

func (db *PostgresConfigDatabase) DeleteOutgoingAccessRequests(ctx context.Context, organizationSerialNumber, serviceName string) error {
	return db.DB.
		WithContext(ctx).
		Where("organization_serial_number = ? AND service_name = ?", organizationSerialNumber, serviceName).
		Delete(&OutgoingAccessRequest{}).
		Error
}

func (db *PostgresConfigDatabase) DeleteOutgoingAccessRequest(ctx context.Context, id uint) error {
	return db.queries.DeleteOutgoingAccessRequest(ctx, int32(id))
}
