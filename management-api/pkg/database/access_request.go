// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"fmt"
	"path"
	"time"
)

type AccessRequest struct {
	ID               string              `json:"id,omitempty"`
	OrganizationName string              `json:"organizationName,omitempty"`
	ServiceName      string              `json:"serviceName,omitempty"`
	Status           AccessRequestStatus `json:"status,omitempty"`
	CreatedAt        time.Time           `json:"createdAt,omitempty"`
	UpdatedAt        time.Time           `json:"updatedAt,omitempty"`
}

type AccessRequestStatus int

const (
	AccessRequestFailed AccessRequestStatus = iota
	AccessRequestCreated
	AccessRequestSent
)

func (db ETCDConfigDatabase) ListAllOutgoingAccessRequests(ctx context.Context) ([]*AccessRequest, error) {
	key := path.Join("access-requests", "outgoing")

	r := []*AccessRequest{}

	err := db.list(ctx, key, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db ETCDConfigDatabase) ListOutgoingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*AccessRequest, error) {
	key := path.Join("access-requests", "outgoing", organizationName, serviceName)

	r := []*AccessRequest{}

	err := db.list(ctx, key, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db ETCDConfigDatabase) CreateAccessRequest(ctx context.Context, accessRequest *AccessRequest) (*AccessRequest, error) {
	t := db.clock.Now()
	id := fmt.Sprintf("%x", t.UnixNano())

	key := path.Join("access-requests", "outgoing", accessRequest.OrganizationName, accessRequest.ServiceName, id)

	accessRequest.ID = id
	accessRequest.Status = AccessRequestCreated
	accessRequest.CreatedAt = t
	accessRequest.UpdatedAt = t

	err := db.put(ctx, key, accessRequest)
	if err != nil {
		return nil, err
	}

	return accessRequest, nil
}
