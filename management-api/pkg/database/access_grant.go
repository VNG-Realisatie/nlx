// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
)

var ErrAccessGrantAlreadyRevoked = errors.New("accessGrant is already revoked")

type AccessGrant struct {
	ID                      uint
	IncomingAccessRequestID uint
	IncomingAccessRequest   *IncomingAccessRequest
	CreatedAt               time.Time
	RevokedAt               sql.NullTime
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

//nolint:dupl // looks the same as other methods but we want to keep these separate to avoid abstracting too soon
func (db *PostgresConfigDatabase) GetAccessGrant(ctx context.Context, id uint) (*AccessGrant, error) {
	accessGrant, err := db.queries.GetAccessGrant(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	var serviceID uint

	if accessGrant.AccessRequestIncomingServiceID.Valid {
		serviceID = uint(accessGrant.AccessRequestIncomingServiceID.Int32)
	}

	service := &Service{
		ID:                     serviceID,
		Name:                   accessGrant.ServiceName.String,
		EndpointURL:            accessGrant.ServiceEndpointUrl.String,
		DocumentationURL:       accessGrant.ServiceDocumentationUrl.String,
		APISpecificationURL:    accessGrant.ServiceApiSpecificationUrl.String,
		Internal:               accessGrant.ServiceInternal.Bool,
		TechSupportContact:     accessGrant.ServiceTechSupportContact.String,
		PublicSupportContact:   accessGrant.ServicePublicSupportContact.String,
		Inways:                 nil,
		IncomingAccessRequests: nil,
		OneTimeCosts:           int(accessGrant.ServiceOneTimeCosts.Int32),
		MonthlyCosts:           int(accessGrant.ServiceMonthlyCosts.Int32),
		RequestCosts:           int(accessGrant.ServiceRequestCosts.Int32),
		CreatedAt:              accessGrant.ServiceCreatedAt.Time,
		UpdatedAt:              accessGrant.ServiceUpdatedAt.Time,
	}

	result := &AccessGrant{
		ID:                      uint(accessGrant.ID),
		IncomingAccessRequestID: uint(accessGrant.AccessRequestIncomingID),
		IncomingAccessRequest: &IncomingAccessRequest{
			ID:        uint(accessGrant.AccessRequestIncomingID),
			ServiceID: serviceID,
			Organization: IncomingAccessRequestOrganization{
				Name:         accessGrant.AccessRequestIncomingOrganizationName.String,
				SerialNumber: accessGrant.AccessRequestIncomingOrganizationSerialNumber.String,
			},
			State:                IncomingAccessRequestState(accessGrant.AccessRequestIncomingState.String),
			AccessGrants:         nil,
			PublicKeyFingerprint: accessGrant.AccessRequestIncomingPublicKeyFingerprint.String,
			PublicKeyPEM:         accessGrant.AccessRequestIncomingPublicKeyPem.String,
			Service:              service,
			CreatedAt:            accessGrant.AccessRequestIncomingCreatedAt.Time,
			UpdatedAt:            accessGrant.AccessRequestIncomingUpdatedAt.Time,
		},
		CreatedAt: accessGrant.CreatedAt,
		RevokedAt: accessGrant.RevokedAt,
	}

	return result, nil
}

//nolint:dupl // looks the same as other methods but we want to keep these separate to avoid abstracting too soon
func (db *PostgresConfigDatabase) RevokeAccessGrant(ctx context.Context, accessGrantID uint, revokedAt time.Time) (*AccessGrant, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		err = tx.Rollback()
		if err != nil {
			if errors.Is(err, sql.ErrTxDone) {
				return
			}

			fmt.Printf("cannot rollback database transaction for revoke access grant: %e", err)
		}
	}()

	qtx := db.queries.WithTx(tx)

	accessGrant, err := qtx.GetAccessGrant(ctx, int32(accessGrantID))
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

	err = qtx.UpdateIncomingAccessRequest(ctx, &queries.UpdateIncomingAccessRequestParams{
		ID:        accessGrant.AccessRequestIncomingID,
		State:     string(IncomingAccessRequestRevoked),
		UpdatedAt: revokedAt,
	})
	if err != nil {
		return nil, err
	}

	accessGrant, err = qtx.GetAccessGrant(ctx, int32(accessGrantID))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	serviceID := uint(accessGrant.AccessRequestIncomingServiceID.Int32)

	service := &Service{
		ID:                     serviceID,
		Name:                   accessGrant.ServiceName.String,
		EndpointURL:            accessGrant.ServiceEndpointUrl.String,
		DocumentationURL:       accessGrant.ServiceDocumentationUrl.String,
		APISpecificationURL:    accessGrant.ServiceApiSpecificationUrl.String,
		Internal:               accessGrant.ServiceInternal.Bool,
		TechSupportContact:     accessGrant.ServiceTechSupportContact.String,
		PublicSupportContact:   accessGrant.ServicePublicSupportContact.String,
		Inways:                 nil,
		IncomingAccessRequests: nil,
		OneTimeCosts:           int(accessGrant.ServiceOneTimeCosts.Int32),
		MonthlyCosts:           int(accessGrant.ServiceMonthlyCosts.Int32),
		RequestCosts:           int(accessGrant.ServiceRequestCosts.Int32),
		CreatedAt:              accessGrant.ServiceCreatedAt.Time,
		UpdatedAt:              accessGrant.ServiceUpdatedAt.Time,
	}

	result := &AccessGrant{
		ID:                      uint(accessGrant.ID),
		IncomingAccessRequestID: uint(accessGrant.AccessRequestIncomingID),
		IncomingAccessRequest: &IncomingAccessRequest{
			ID:        uint(accessGrant.AccessRequestIncomingID),
			ServiceID: serviceID,
			Organization: IncomingAccessRequestOrganization{
				Name:         accessGrant.AccessRequestIncomingOrganizationName.String,
				SerialNumber: accessGrant.AccessRequestIncomingOrganizationSerialNumber.String,
			},
			State:                IncomingAccessRequestState(accessGrant.AccessRequestIncomingState.String),
			AccessGrants:         nil,
			PublicKeyFingerprint: accessGrant.AccessRequestIncomingPublicKeyFingerprint.String,
			PublicKeyPEM:         accessGrant.AccessRequestIncomingPublicKeyPem.String,
			Service:              service,
			CreatedAt:            accessGrant.AccessRequestIncomingCreatedAt.Time,
			UpdatedAt:            accessGrant.AccessRequestIncomingUpdatedAt.Time,
		},
		CreatedAt: accessGrant.CreatedAt,
		RevokedAt: accessGrant.RevokedAt,
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db *PostgresConfigDatabase) ListAccessGrantsForService(ctx context.Context, serviceName string) ([]*AccessGrant, error) {
	result := []*AccessGrant{}

	accessGrants, err := db.queries.ListAccessGrantsForService(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	for _, accessGrant := range accessGrants {
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
			CreatedAt: accessGrant.CreatedAt,
			RevokedAt: accessGrant.RevokedAt,
		}

		result = append(result, newModel)
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

	var serviceID uint

	if accessGrant.AccessRequestIncomingServiceID > 0 {
		serviceID = uint(accessGrant.AccessRequestIncomingServiceID)
	}

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
		CreatedAt: accessGrant.CreatedAt,
		RevokedAt: accessGrant.RevokedAt,
	}

	return result, nil
}
