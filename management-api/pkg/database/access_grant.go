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

func (a *AccessGrant) TableName() string {
	return "nlx_management.access_grants"
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

	if accessGrant.AriServiceID.Valid {
		serviceID = uint(accessGrant.AriServiceID.Int32)
	}

	service := &Service{
		ID:                     serviceID,
		Name:                   accessGrant.SName.String,
		EndpointURL:            accessGrant.SEndpointUrl.String,
		DocumentationURL:       accessGrant.SDocumentationUrl.String,
		APISpecificationURL:    accessGrant.SApiSpecificationUrl.String,
		Internal:               accessGrant.SInternal.Bool,
		TechSupportContact:     accessGrant.STechSupportContact.String,
		PublicSupportContact:   accessGrant.SPublicSupportContact.String,
		Inways:                 nil,
		IncomingAccessRequests: nil,
		OneTimeCosts:           int(accessGrant.SOneTimeCosts.Int32),
		MonthlyCosts:           int(accessGrant.SMonthlyCosts.Int32),
		RequestCosts:           int(accessGrant.SRequestCosts.Int32),
		CreatedAt:              accessGrant.SCreatedAt.Time,
		UpdatedAt:              accessGrant.SUpdatedAt.Time,
	}

	result := &AccessGrant{
		ID:                      uint(accessGrant.ID),
		IncomingAccessRequestID: uint(accessGrant.AccessRequestIncomingID),
		IncomingAccessRequest: &IncomingAccessRequest{
			ID:        uint(accessGrant.AccessRequestIncomingID),
			ServiceID: serviceID,
			Organization: IncomingAccessRequestOrganization{
				Name:         accessGrant.AriOrganizationName.String,
				SerialNumber: accessGrant.AriOrganizationSerialNumber.String,
			},
			State:                IncomingAccessRequestState(accessGrant.AriState.String),
			AccessGrants:         nil,
			PublicKeyFingerprint: accessGrant.AriPublicKeyFingerprint.String,
			PublicKeyPEM:         accessGrant.AriPublicKeyPem.String,
			Service:              service,
			CreatedAt:            accessGrant.AriCreatedAt.Time,
			UpdatedAt:            accessGrant.AriUpdatedAt.Time,
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

	err = qtx.RevokeIncomingAccessRequest(ctx, &queries.RevokeIncomingAccessRequestParams{
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

	var serviceID uint

	if accessGrant.AriServiceID.Valid {
		serviceID = uint(accessGrant.AriServiceID.Int32)
	}

	service := &Service{
		ID:                     serviceID,
		Name:                   accessGrant.SName.String,
		EndpointURL:            accessGrant.SEndpointUrl.String,
		DocumentationURL:       accessGrant.SDocumentationUrl.String,
		APISpecificationURL:    accessGrant.SApiSpecificationUrl.String,
		Internal:               accessGrant.SInternal.Bool,
		TechSupportContact:     accessGrant.STechSupportContact.String,
		PublicSupportContact:   accessGrant.SPublicSupportContact.String,
		Inways:                 nil,
		IncomingAccessRequests: nil,
		OneTimeCosts:           int(accessGrant.SOneTimeCosts.Int32),
		MonthlyCosts:           int(accessGrant.SMonthlyCosts.Int32),
		RequestCosts:           int(accessGrant.SRequestCosts.Int32),
		CreatedAt:              accessGrant.SCreatedAt.Time,
		UpdatedAt:              accessGrant.SUpdatedAt.Time,
	}

	result := &AccessGrant{
		ID:                      uint(accessGrant.ID),
		IncomingAccessRequestID: uint(accessGrant.AccessRequestIncomingID),
		IncomingAccessRequest: &IncomingAccessRequest{
			ID:        uint(accessGrant.AccessRequestIncomingID),
			ServiceID: serviceID,
			Organization: IncomingAccessRequestOrganization{
				Name:         accessGrant.AriOrganizationName.String,
				SerialNumber: accessGrant.AriOrganizationSerialNumber.String,
			},
			State:                IncomingAccessRequestState(accessGrant.AriState.String),
			AccessGrants:         nil,
			PublicKeyFingerprint: accessGrant.AriPublicKeyFingerprint.String,
			PublicKeyPEM:         accessGrant.AriPublicKeyPem.String,
			Service:              service,
			CreatedAt:            accessGrant.AriCreatedAt.Time,
			UpdatedAt:            accessGrant.AriUpdatedAt.Time,
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
		if err == sql.ErrNoRows {
			return result, nil
		}

		return nil, err
	}

	for _, accessGrant := range accessGrants {
		var serviceID uint

		if accessGrant.AriServiceID > 0 {
			serviceID = uint(accessGrant.AriServiceID)
		}

		service := &Service{
			ID:                     serviceID,
			Name:                   accessGrant.SName,
			EndpointURL:            accessGrant.SEndpointUrl,
			DocumentationURL:       accessGrant.SDocumentationUrl,
			APISpecificationURL:    accessGrant.SApiSpecificationUrl,
			Internal:               accessGrant.SInternal,
			TechSupportContact:     accessGrant.STechSupportContact,
			PublicSupportContact:   accessGrant.SPublicSupportContact,
			Inways:                 nil,
			IncomingAccessRequests: nil,
			OneTimeCosts:           int(accessGrant.SOneTimeCosts),
			MonthlyCosts:           int(accessGrant.SMonthlyCosts),
			RequestCosts:           int(accessGrant.SRequestCosts),
			CreatedAt:              accessGrant.SCreatedAt,
			UpdatedAt:              accessGrant.SUpdatedAt,
		}

		newModel := &AccessGrant{
			ID:                      uint(accessGrant.ID),
			IncomingAccessRequestID: uint(accessGrant.AccessRequestIncomingID),
			IncomingAccessRequest: &IncomingAccessRequest{
				ID:        uint(accessGrant.AccessRequestIncomingID),
				ServiceID: serviceID,
				Organization: IncomingAccessRequestOrganization{
					Name:         accessGrant.AriOrganizationName,
					SerialNumber: accessGrant.AriOrganizationSerialNumber,
				},
				State:                IncomingAccessRequestState(accessGrant.AriState),
				AccessGrants:         nil,
				PublicKeyFingerprint: accessGrant.AriPublicKeyFingerprint,
				PublicKeyPEM:         accessGrant.AriPublicKeyPem.String,
				Service:              service,
				CreatedAt:            accessGrant.AriCreatedAt,
				UpdatedAt:            accessGrant.AriUpdatedAt,
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

	if accessGrant.AriServiceID > 0 {
		serviceID = uint(accessGrant.AriServiceID)
	}

	service := &Service{
		ID:                     serviceID,
		Name:                   accessGrant.SName,
		EndpointURL:            accessGrant.SEndpointUrl,
		DocumentationURL:       accessGrant.SDocumentationUrl,
		APISpecificationURL:    accessGrant.SApiSpecificationUrl,
		Internal:               accessGrant.SInternal,
		TechSupportContact:     accessGrant.STechSupportContact,
		PublicSupportContact:   accessGrant.SPublicSupportContact,
		Inways:                 nil,
		IncomingAccessRequests: nil,
		OneTimeCosts:           int(accessGrant.SOneTimeCosts),
		MonthlyCosts:           int(accessGrant.SMonthlyCosts),
		RequestCosts:           int(accessGrant.SRequestCosts),
		CreatedAt:              accessGrant.SCreatedAt,
		UpdatedAt:              accessGrant.SUpdatedAt,
	}

	result := &AccessGrant{
		ID:                      uint(accessGrant.ID),
		IncomingAccessRequestID: uint(accessGrant.AccessRequestIncomingID),
		IncomingAccessRequest: &IncomingAccessRequest{
			ID:        uint(accessGrant.AccessRequestIncomingID),
			ServiceID: serviceID,
			Organization: IncomingAccessRequestOrganization{
				Name:         accessGrant.AriOrganizationName,
				SerialNumber: accessGrant.AriOrganizationSerialNumber,
			},
			State:                IncomingAccessRequestState(accessGrant.AriState),
			AccessGrants:         nil,
			PublicKeyFingerprint: accessGrant.AriPublicKeyFingerprint,
			PublicKeyPEM:         accessGrant.AriPublicKeyPem.String,
			Service:              service,
			CreatedAt:            accessGrant.AriCreatedAt,
			UpdatedAt:            accessGrant.AriUpdatedAt,
		},
		CreatedAt: accessGrant.CreatedAt,
		RevokedAt: accessGrant.RevokedAt,
	}

	return result, nil
}
