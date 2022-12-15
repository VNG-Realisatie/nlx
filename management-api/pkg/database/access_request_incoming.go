// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"time"

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
)

type IncomingAccessRequestState string

const (
	IncomingAccessRequestReceived  IncomingAccessRequestState = "received"
	IncomingAccessRequestApproved  IncomingAccessRequestState = "approved"
	IncomingAccessRequestRejected  IncomingAccessRequestState = "rejected"
	IncomingAccessRequestWithdrawn IncomingAccessRequestState = "withdrawn"
)

type IncomingAccessRequestOrganization struct {
	Name         string
	SerialNumber string
}

type IncomingAccessRequest struct {
	ID                   uint
	ServiceID            uint
	Organization         IncomingAccessRequestOrganization `gorm:"embedded;embeddedPrefix:organization_"`
	State                IncomingAccessRequestState
	AccessGrants         []AccessGrant
	PublicKeyFingerprint string
	PublicKeyPEM         string
	Service              *Service
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (i *IncomingAccessRequest) TableName() string {
	return "nlx_management.access_requests_incoming"
}

func (db *PostgresConfigDatabase) ListIncomingAccessRequests(ctx context.Context, serviceName string) ([]*IncomingAccessRequest, error) {
	result := []*IncomingAccessRequest{}

	incomingAccessRequests, err := db.queries.ListIncomingAccessRequests(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	for _, accessRequest := range incomingAccessRequests {
		result = append(result, &IncomingAccessRequest{
			ID: uint(accessRequest.ID),
			Organization: IncomingAccessRequestOrganization{
				Name:         accessRequest.OrganizationName,
				SerialNumber: accessRequest.OrganizationSerialNumber,
			},
			State: IncomingAccessRequestState(accessRequest.State),
			Service: &Service{
				Name: serviceName,
			},
			CreatedAt: accessRequest.CreatedAt,
			UpdatedAt: accessRequest.UpdatedAt,
		})
	}

	return result, nil
}

func (db *PostgresConfigDatabase) GetLatestIncomingAccessRequest(ctx context.Context, organizationSerialNumber, serviceName, publicKeyFingerprint string) (*IncomingAccessRequest, error) {
	incomingAccessRequests, err := db.queries.GetLatestIncomingAccessRequest(ctx, &queries.GetLatestIncomingAccessRequestParams{
		OrganizationSerialNumber: organizationSerialNumber,
		PublicKeyFingerprint:     publicKeyFingerprint,
		Name:                     serviceName,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	result := &IncomingAccessRequest{
		ID:        uint(incomingAccessRequests.ID),
		State:     IncomingAccessRequestState(incomingAccessRequests.State),
		CreatedAt: incomingAccessRequests.CreatedAt,
		UpdatedAt: incomingAccessRequests.UpdatedAt,
	}

	return result, nil
}

func (db *PostgresConfigDatabase) GetIncomingAccessRequestCountByService(ctx context.Context) (map[string]int, error) {
	counts, err := db.queries.GetIncomingAccessRequestsByServiceCount(ctx)
	if err != nil {
		return nil, err
	}

	countPerService := make(map[string]int)

	for _, value := range counts {
		countPerService[value.Name] = int(value.Count)
	}

	return countPerService, nil
}

func (db *PostgresConfigDatabase) GetIncomingAccessRequest(ctx context.Context, id uint) (*IncomingAccessRequest, error) {
	accessRequest, err := db.queries.GetIncomingAccessRequest(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &IncomingAccessRequest{
		ID:        uint(accessRequest.ID),
		ServiceID: uint(accessRequest.ServiceID),
		Organization: IncomingAccessRequestOrganization{
			Name:         accessRequest.OrganizationName,
			SerialNumber: accessRequest.OrganizationSerialNumber,
		},
		State:                IncomingAccessRequestState(accessRequest.State),
		AccessGrants:         nil,
		PublicKeyFingerprint: accessRequest.PublicKeyFingerprint,
		PublicKeyPEM:         accessRequest.PublicKeyPem.String,
		Service:              nil,
		CreatedAt:            accessRequest.CreatedAt,
		UpdatedAt:            accessRequest.UpdatedAt,
	}, nil
}

func (db *PostgresConfigDatabase) CreateIncomingAccessRequest(ctx context.Context, accessRequest *IncomingAccessRequest) (*IncomingAccessRequest, error) {
	pem := sql.NullString{Valid: false}

	if accessRequest.PublicKeyPEM != "" {
		pem.Valid = true
		pem.String = accessRequest.PublicKeyPEM
	}

	id, err := db.queries.CreateIncomingAccessRequest(ctx, &queries.CreateIncomingAccessRequestParams{
		State:                    string(accessRequest.State),
		OrganizationName:         accessRequest.Organization.Name,
		OrganizationSerialNumber: accessRequest.Organization.SerialNumber,
		PublicKeyFingerprint:     accessRequest.PublicKeyFingerprint,
		PublicKeyPem:             pem,
		ServiceID:                int32(accessRequest.ServiceID),
		CreatedAt:                accessRequest.CreatedAt,
		UpdatedAt:                accessRequest.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	return &IncomingAccessRequest{
		ID:    uint(id),
		State: accessRequest.State,
	}, nil
}

func (db *PostgresConfigDatabase) UpdateIncomingAccessRequestState(ctx context.Context, accessRequestID uint, state IncomingAccessRequestState) error {
	rowsAffected, err := db.queries.UpdateIncomingAccessRequestState(ctx, &queries.UpdateIncomingAccessRequestStateParams{
		ID:        int32(accessRequestID),
		State:     string(state),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (db *PostgresConfigDatabase) DeleteIncomingAccessRequest(ctx context.Context, id uint) error {
	return db.queries.DeleteIncomingAccessRequest(ctx, int32(id))
}
