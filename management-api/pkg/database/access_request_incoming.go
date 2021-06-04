// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IncomingAccessRequestState string

const (
	IncomingAccessRequestReceived IncomingAccessRequestState = "received"
	IncomingAccessRequestApproved IncomingAccessRequestState = "approved"
	IncomingAccessRequestRejected IncomingAccessRequestState = "rejected"
	IncomingAccessRequestRevoked  IncomingAccessRequestState = "revoked"
)

type IncomingAccessRequest struct {
	ID                   uint
	ServiceID            uint
	OrganizationName     string
	State                IncomingAccessRequestState
	AccessGrants         []AccessGrant
	PublicKeyPEM         string
	PublicKeyFingerprint string
	Service              *Service
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (i *IncomingAccessRequest) TableName() string {
	return "nlx_management.access_requests_incoming"
}

func (db *PostgresConfigDatabase) ListAllIncomingAccessRequests(ctx context.Context) ([]*IncomingAccessRequest, error) {
	accessRequests := []*IncomingAccessRequest{}

	if err := db.DB.
		WithContext(ctx).
		Preload("Service").
		Find(&accessRequests).Error; err != nil {
		return nil, err
	}

	return accessRequests, nil
}

func (db *PostgresConfigDatabase) ListIncomingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*IncomingAccessRequest, error) {
	accessRequests := []*IncomingAccessRequest{}

	if err := db.DB.
		WithContext(ctx).
		Preload("Service").
		Joins("JOIN nlx_management.services s ON s.service_id = access_requests_incoming.service_id AND organization_name = ? AND s.name = ?", organizationName, serviceName).
		Find(&accessRequests).Error; err != nil {
		return nil, err
	}

	return accessRequests, nil
}

func (db *PostgresConfigDatabase) GetLatestIncomingAccessRequest(ctx context.Context, organizationName, serviceName string) (*IncomingAccessRequest, error) {
	accessRequest := &IncomingAccessRequest{}

	if err := db.DB.Debug().
		WithContext(ctx).
		Preload("Service").
		Joins("JOIN nlx_management.services s ON s.id = access_requests_incoming.service_id AND access_requests_incoming.organization_name = ? AND s.name = ?", organizationName, serviceName).
		Order("created_at DESC").
		First(&accessRequest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return accessRequest, nil
}

func (db *PostgresConfigDatabase) GetIncomingAccessRequestCountByService(ctx context.Context) (map[string]int, error) {
	result := []map[string]interface{}{}

	if err := db.DB.Debug().
		WithContext(ctx).
		Model(&Service{}).
		Select("services.name, COUNT(a.id)").
		Joins("LEFT JOIN nlx_management.access_requests_incoming a ON a.service_id = services.id AND a.state = 'received'").
		Group("services.id").
		Find(&result).Error; err != nil {
		return nil, err
	}

	countPerService := make(map[string]int)

	for _, value := range result {
		countPerService[value["name"].(string)] = int(value["count"].(int64))
	}

	return countPerService, nil
}

func (db *PostgresConfigDatabase) GetIncomingAccessRequest(ctx context.Context, id uint) (*IncomingAccessRequest, error) {
	accessRequest := &IncomingAccessRequest{}
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

func (db *PostgresConfigDatabase) CreateIncomingAccessRequest(ctx context.Context, accessRequest *IncomingAccessRequest) (*IncomingAccessRequest, error) {
	if err := db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Create(accessRequest).Error; err != nil {
		return nil, err
	}

	return accessRequest, nil
}

func (db *PostgresConfigDatabase) UpdateIncomingAccessRequestState(ctx context.Context, accessRequestID uint, state IncomingAccessRequestState) error {
	incomingAccessRequest := &IncomingAccessRequest{}

	if err := db.DB.
		WithContext(ctx).
		First(incomingAccessRequest, accessRequestID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}

		return err
	}

	incomingAccessRequest.State = state

	return db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Select("state", "updated_at").
		Save(incomingAccessRequest).Error
}
