// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"

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

// nolint:dupl // conversion of models can be unified once we finished moving from Gorm to Sqlc
func (db *PostgresConfigDatabase) ListLatestOutgoingAccessRequests(ctx context.Context, organizationSerialNumber, serviceName string) ([]*OutgoingAccessRequest, error) {
	accessRequests, err := db.queries.ListLatestOutgoingAccessRequests(ctx, &queries.ListLatestOutgoingAccessRequestsParams{
		OrganizationSerialNumber: organizationSerialNumber,
		ServiceName:              serviceName,
	})
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
	count, err := db.queries.CountReceivedOutgoingAccessRequestsForOutway(ctx, &queries.CountReceivedOutgoingAccessRequestsForOutwayParams{
		OrganizationSerialNumber: accessRequest.Organization.SerialNumber,
		ServiceName:              accessRequest.ServiceName,
		PublicKeyFingerprint:     accessRequest.PublicKeyFingerprint,
	})
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, ErrActiveAccessRequest
	}

	var pem = sql.NullString{}

	if accessRequest.PublicKeyPEM != "" {
		pem.String = accessRequest.PublicKeyPEM
		pem.Valid = true
	}

	id, err := db.queries.CreateOutgoingAccessRequest(ctx, &queries.CreateOutgoingAccessRequestParams{
		State:                    string(accessRequest.State),
		OrganizationName:         accessRequest.Organization.Name,
		OrganizationSerialNumber: accessRequest.Organization.SerialNumber,
		PublicKeyFingerprint:     accessRequest.PublicKeyFingerprint,
		PublicKeyPem:             pem,
		ServiceName:              accessRequest.ServiceName,
		CreatedAt:                accessRequest.CreatedAt,
		UpdatedAt:                accessRequest.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	accessRequest.ID = uint(id)

	return accessRequest, nil
}

func (db *PostgresConfigDatabase) GetOutgoingAccessRequest(ctx context.Context, id uint) (*OutgoingAccessRequest, error) {
	outgoingAccessRequest, err := db.queries.GetOutgoingAccessRequest(ctx, int32(id))
	if err != nil {
		return nil, ErrNotFound
	}

	var pem = ""

	if outgoingAccessRequest.PublicKeyPem.Valid {
		pem = outgoingAccessRequest.PublicKeyPem.String
	}

	var errorCause = ""

	if outgoingAccessRequest.ErrorCause.Valid {
		errorCause = outgoingAccessRequest.ErrorCause.String
	}

	result := &OutgoingAccessRequest{
		ID: uint(outgoingAccessRequest.ID),
		Organization: Organization{
			SerialNumber: outgoingAccessRequest.OrganizationSerialNumber,
			Name:         outgoingAccessRequest.OrganizationName,
		},
		ServiceName:          outgoingAccessRequest.ServiceName,
		ReferenceID:          uint(outgoingAccessRequest.ReferenceID),
		State:                OutgoingAccessRequestState(outgoingAccessRequest.State),
		PublicKeyFingerprint: outgoingAccessRequest.PublicKeyFingerprint,
		PublicKeyPEM:         pem,
		ErrorCode:            int(outgoingAccessRequest.ErrorCode),
		ErrorCause:           errorCause,
		CreatedAt:            outgoingAccessRequest.CreatedAt,
		UpdatedAt:            outgoingAccessRequest.UpdatedAt,
	}

	return result, nil
}

func (db *PostgresConfigDatabase) GetLatestOutgoingAccessRequest(ctx context.Context, organizationSerialNumber, serviceName, publicKeyFingerprint string) (*OutgoingAccessRequest, error) {
	outgoingAccessRequest, err := db.queries.GetLatestOutgoingAccessRequest(ctx, &queries.GetLatestOutgoingAccessRequestParams{
		OrganizationSerialNumber: organizationSerialNumber,
		ServiceName:              serviceName,
		PublicKeyFingerprint:     publicKeyFingerprint,
	})
	if err != nil {
		return nil, ErrNotFound
	}

	var pem = ""

	if outgoingAccessRequest.PublicKeyPem.Valid {
		pem = outgoingAccessRequest.PublicKeyPem.String
	}

	var errorCause = ""

	if outgoingAccessRequest.ErrorCause.Valid {
		errorCause = outgoingAccessRequest.ErrorCause.String
	}

	result := &OutgoingAccessRequest{
		ID: uint(outgoingAccessRequest.ID),
		Organization: Organization{
			SerialNumber: outgoingAccessRequest.OrganizationSerialNumber,
			Name:         outgoingAccessRequest.OrganizationName,
		},
		ServiceName:          outgoingAccessRequest.ServiceName,
		ReferenceID:          uint(outgoingAccessRequest.ReferenceID),
		State:                OutgoingAccessRequestState(outgoingAccessRequest.State),
		PublicKeyFingerprint: outgoingAccessRequest.PublicKeyFingerprint,
		PublicKeyPEM:         pem,
		ErrorCode:            int(outgoingAccessRequest.ErrorCode),
		ErrorCause:           errorCause,
		CreatedAt:            outgoingAccessRequest.CreatedAt,
		UpdatedAt:            outgoingAccessRequest.UpdatedAt,
	}

	return result, nil
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
