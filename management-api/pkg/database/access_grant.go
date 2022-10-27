// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
)

var ErrAccessGrantAlreadyRevoked = errors.New("accessGrant is already revoked")
var ErrAccessGrantAlreadyTerminated = errors.New("accessGrant is already terminated")
var ErrAccessProofAlreadyTerminated = errors.New("accessProof is already terminated")

type AccessGrant struct {
	ID                      uint
	IncomingAccessRequestID uint
	IncomingAccessRequest   *IncomingAccessRequest
	CreatedAt               time.Time
	RevokedAt               sql.NullTime
	TerminatedAt            sql.NullTime
}

func (db *PostgresConfigDatabase) CreateAccessGrant(ctx context.Context, accessRequest *IncomingAccessRequest) (*AccessGrant, error) {
	result := &AccessGrant{
		IncomingAccessRequestID: accessRequest.ID,
		CreatedAt:               time.Now(),
	}

	id, err := db.queries.CreateAccessGrant(ctx, &queries.CreateAccessGrantParams{
		AccessRequestIncomingID: int32(result.IncomingAccessRequestID),
		CreatedAt:               result.CreatedAt,
	})
	if err != nil {
		return nil, err
	}

	result.ID = uint(id)

	return result, nil
}

func (db *PostgresConfigDatabase) GetAccessGrantIDForIncomingAccessRequest(ctx context.Context, accessRequestID uint) (uint, error) {
	accessGrantID, err := db.queries.GetAccessGrantIDOfIncomingAccessRequest(ctx, int32(accessRequestID))
	if err != nil {
		return 0, err
	}

	return uint(accessGrantID), nil
}

// nolint:dupl // is similar to terminate access proof
func (db *PostgresConfigDatabase) TerminateAccessGrant(ctx context.Context, accessGrantID uint, terminatedAt time.Time) error {
	accessGrant, err := db.queries.GetAccessGrant(ctx, int32(accessGrantID))
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}

		return err
	}

	if accessGrant.TerminatedAt.Valid {
		return ErrAccessGrantAlreadyTerminated
	}

	err = db.queries.TerminateAccessGrant(ctx, &queries.TerminateAccessGrantParams{
		ID: int32(accessGrantID),
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

//nolint:dupl // looks the same as other methods but we want to keep these separate to avoid abstracting too soon
func (db *PostgresConfigDatabase) GetAccessGrant(ctx context.Context, id uint) (*AccessGrant, error) {
	accessGrant, err := db.queries.GetAccessGrant(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	serviceID := uint(accessGrant.AccessRequestIncomingServiceID)

	service := &Service{
		ID:                     serviceID,
		Name:                   accessGrant.ServiceName,
		EndpointURL:            accessGrant.ServiceEndpointUrl,
		DocumentationURL:       accessGrant.ServiceDocumentationUrl,
		APISpecificationURL:    accessGrant.ServiceApiSpecificationUrl,
		Internal:               accessGrant.ServiceInternal,
		TechSupportContact:     accessGrant.ServiceTechSupportContact,
		PublicSupportContact:   accessGrant.ServicePublicSupportContact,
		Inways:                 nil,
		IncomingAccessRequests: nil,
		OneTimeCosts:           int(accessGrant.ServiceOneTimeCosts),
		MonthlyCosts:           int(accessGrant.ServiceMonthlyCosts),
		RequestCosts:           int(accessGrant.ServiceRequestCosts),
		CreatedAt:              accessGrant.ServiceCreatedAt,
		UpdatedAt:              accessGrant.ServiceUpdatedAt,
	}

	result := &AccessGrant{
		ID:                      uint(accessGrant.ID),
		IncomingAccessRequestID: uint(accessGrant.AccessRequestIncomingID),
		IncomingAccessRequest: &IncomingAccessRequest{
			ID:        uint(accessGrant.AccessRequestIncomingID),
			ServiceID: serviceID,
			Organization: IncomingAccessRequestOrganization{
				Name:         accessGrant.AccessRequestIncomingOrganizationName,
				SerialNumber: accessGrant.AccessRequestIncomingOrganizationSerialNumber,
			},
			State:                IncomingAccessRequestState(accessGrant.AccessRequestIncomingState),
			AccessGrants:         nil,
			PublicKeyFingerprint: accessGrant.AccessRequestIncomingPublicKeyFingerprint,
			PublicKeyPEM:         accessGrant.AccessRequestIncomingPublicKeyPem.String,
			Service:              service,
			CreatedAt:            accessGrant.AccessRequestIncomingCreatedAt,
			UpdatedAt:            accessGrant.AccessRequestIncomingUpdatedAt,
		},
		CreatedAt:    accessGrant.CreatedAt,
		TerminatedAt: accessGrant.TerminatedAt,
		RevokedAt:    accessGrant.RevokedAt,
	}

	return result, nil
}

//nolint:dupl // looks the same as other methods but we want to keep these separate to avoid abstracting too soon
func (db *PostgresConfigDatabase) RevokeAccessGrant(ctx context.Context, accessGrantID uint, revokedAt time.Time) (*AccessGrant, error) {
	accessGrant, err := db.queries.GetAccessGrant(ctx, int32(accessGrantID))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	if accessGrant.RevokedAt.Valid {
		return nil, ErrAccessGrantAlreadyRevoked
	}

	err = db.queries.RevokeAccessGrant(ctx, &queries.RevokeAccessGrantParams{
		ID: int32(accessGrantID),
		RevokedAt: sql.NullTime{
			Time:  revokedAt,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	accessGrant.RevokedAt.Valid = true
	accessGrant.RevokedAt.Time = revokedAt

	serviceID := uint(accessGrant.AccessRequestIncomingServiceID)

	service := &Service{
		ID:                     serviceID,
		Name:                   accessGrant.ServiceName,
		EndpointURL:            accessGrant.ServiceEndpointUrl,
		DocumentationURL:       accessGrant.ServiceDocumentationUrl,
		APISpecificationURL:    accessGrant.ServiceApiSpecificationUrl,
		Internal:               accessGrant.ServiceInternal,
		TechSupportContact:     accessGrant.ServiceTechSupportContact,
		PublicSupportContact:   accessGrant.ServicePublicSupportContact,
		Inways:                 nil,
		IncomingAccessRequests: nil,
		OneTimeCosts:           int(accessGrant.ServiceOneTimeCosts),
		MonthlyCosts:           int(accessGrant.ServiceMonthlyCosts),
		RequestCosts:           int(accessGrant.ServiceRequestCosts),
		CreatedAt:              accessGrant.ServiceCreatedAt,
		UpdatedAt:              accessGrant.ServiceUpdatedAt,
	}

	result := &AccessGrant{
		ID:                      uint(accessGrant.ID),
		IncomingAccessRequestID: uint(accessGrant.AccessRequestIncomingID),
		IncomingAccessRequest: &IncomingAccessRequest{
			ID:        uint(accessGrant.AccessRequestIncomingID),
			ServiceID: serviceID,
			Organization: IncomingAccessRequestOrganization{
				Name:         accessGrant.AccessRequestIncomingOrganizationName,
				SerialNumber: accessGrant.AccessRequestIncomingOrganizationSerialNumber,
			},
			State:                IncomingAccessRequestState(accessGrant.AccessRequestIncomingState),
			AccessGrants:         nil,
			PublicKeyFingerprint: accessGrant.AccessRequestIncomingPublicKeyFingerprint,
			PublicKeyPEM:         accessGrant.AccessRequestIncomingPublicKeyPem.String,
			Service:              service,
			CreatedAt:            accessGrant.AccessRequestIncomingCreatedAt,
			UpdatedAt:            accessGrant.AccessRequestIncomingUpdatedAt,
		},
		CreatedAt:    accessGrant.CreatedAt,
		TerminatedAt: accessGrant.TerminatedAt,
		RevokedAt:    accessGrant.RevokedAt,
	}

	return result, nil
}

func (db *PostgresConfigDatabase) ListAccessGrantsForService(ctx context.Context, serviceName string) ([]*AccessGrant, error) {
	accessGrants, err := db.queries.ListAccessGrantsForService(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	result := make([]*AccessGrant, len(accessGrants))

	for i, accessGrant := range accessGrants {
		serviceID := uint(accessGrant.AccessRequestIncomingServiceID)

		service := &Service{
			ID:                     serviceID,
			Name:                   accessGrant.ServiceName,
			EndpointURL:            accessGrant.ServiceEndpointUrl,
			DocumentationURL:       accessGrant.ServiceDocumentationUrl,
			APISpecificationURL:    accessGrant.ServiceApiSpecificationUrl,
			Internal:               accessGrant.ServiceInternal,
			TechSupportContact:     accessGrant.ServiceTechSupportContact,
			PublicSupportContact:   accessGrant.ServicePublicSupportContact,
			Inways:                 nil,
			IncomingAccessRequests: nil,
			OneTimeCosts:           int(accessGrant.ServiceOneTimeCosts),
			MonthlyCosts:           int(accessGrant.ServiceMonthlyCosts),
			RequestCosts:           int(accessGrant.ServiceRequestCosts),
			CreatedAt:              accessGrant.ServiceCreatedAt,
			UpdatedAt:              accessGrant.ServiceUpdatedAt,
		}

		newModel := &AccessGrant{
			ID:                      uint(accessGrant.ID),
			IncomingAccessRequestID: uint(accessGrant.AccessRequestIncomingID),
			IncomingAccessRequest: &IncomingAccessRequest{
				ID:        uint(accessGrant.AccessRequestIncomingID),
				ServiceID: serviceID,
				Organization: IncomingAccessRequestOrganization{
					Name:         accessGrant.AccessRequestIncomingOrganizationName,
					SerialNumber: accessGrant.AccessRequestIncomingOrganizationSerialNumber,
				},
				State:                IncomingAccessRequestState(accessGrant.AccessRequestIncomingState),
				AccessGrants:         nil,
				PublicKeyFingerprint: accessGrant.AccessRequestIncomingPublicKeyFingerprint,
				PublicKeyPEM:         accessGrant.AccessRequestIncomingPublicKeyPem.String,
				Service:              service,
				CreatedAt:            accessGrant.AccessRequestIncomingCreatedAt,
				UpdatedAt:            accessGrant.AccessRequestIncomingUpdatedAt,
			},
			CreatedAt:    accessGrant.CreatedAt,
			RevokedAt:    accessGrant.RevokedAt,
			TerminatedAt: accessGrant.TerminatedAt,
		}

		result[i] = newModel
	}

	return result, nil
}

func (db *PostgresConfigDatabase) GetLatestAccessGrantForService(ctx context.Context, organizationSerialNumber, serviceName, publicKeyFingerprint string) (*AccessGrant, error) {
	accessGrant, err := db.queries.GetLatestAccessGrantForService(ctx, &queries.GetLatestAccessGrantForServiceParams{
		ServiceName:              serviceName,
		OrganizationSerialNumber: organizationSerialNumber,
		PublicKeyFingerprint:     publicKeyFingerprint,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	serviceID := uint(accessGrant.AccessRequestIncomingServiceID)

	service := &Service{
		ID:                     serviceID,
		Name:                   accessGrant.ServiceName,
		EndpointURL:            accessGrant.ServiceEndpointUrl,
		DocumentationURL:       accessGrant.ServiceDocumentationUrl,
		APISpecificationURL:    accessGrant.ServiceApiSpecificationUrl,
		Internal:               accessGrant.ServiceInternal,
		TechSupportContact:     accessGrant.ServiceTechSupportContact,
		PublicSupportContact:   accessGrant.ServicePublicSupportContact,
		Inways:                 nil,
		IncomingAccessRequests: nil,
		OneTimeCosts:           int(accessGrant.ServiceOneTimeCosts),
		MonthlyCosts:           int(accessGrant.ServiceMonthlyCosts),
		RequestCosts:           int(accessGrant.ServiceRequestCosts),
		CreatedAt:              accessGrant.ServiceCreatedAt,
		UpdatedAt:              accessGrant.ServiceUpdatedAt,
	}

	result := &AccessGrant{
		ID:                      uint(accessGrant.ID),
		IncomingAccessRequestID: uint(accessGrant.AccessRequestIncomingID),
		IncomingAccessRequest: &IncomingAccessRequest{
			ID:        uint(accessGrant.AccessRequestIncomingID),
			ServiceID: serviceID,
			Organization: IncomingAccessRequestOrganization{
				Name:         accessGrant.AccessRequestIncomingOrganizationName,
				SerialNumber: accessGrant.AccessRequestIncomingOrganizationSerialNumber,
			},
			State:                IncomingAccessRequestState(accessGrant.AccessRequestIncomingState),
			AccessGrants:         nil,
			PublicKeyFingerprint: accessGrant.AccessRequestIncomingPublicKeyFingerprint,
			PublicKeyPEM:         accessGrant.AccessRequestIncomingPublicKeyPem.String,
			Service:              service,
			CreatedAt:            accessGrant.AccessRequestIncomingCreatedAt,
			UpdatedAt:            accessGrant.AccessRequestIncomingUpdatedAt,
		},
		CreatedAt:    accessGrant.CreatedAt,
		RevokedAt:    accessGrant.RevokedAt,
		TerminatedAt: accessGrant.TerminatedAt,
	}

	return result, nil
}
