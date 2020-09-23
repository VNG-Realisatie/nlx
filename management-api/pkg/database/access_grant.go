// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"fmt"
	"path"
	"time"
)

type AccessGrant struct {
	ID                   string    `json:"id,omitempty"`
	AccessRequestID      string    `json:"accessRequestId,omitempty"`
	OrganizationName     string    `json:"organizationName,omitempty"`
	ServiceName          string    `json:"serviceName,omitempty"`
	PublicKeyFingerprint string    `json:"publicKeyFingerprint,omitempty"`
	CreatedAt            time.Time `json:"createdAt,omitempty"`
}

func (db ETCDConfigDatabase) CreateAccessGrant(ctx context.Context, accessGrant *AccessGrant) (*AccessGrant, error) {
	t := db.clock.Now()
	id := fmt.Sprintf("%x", t.UnixNano())

	key := path.Join("access-grants", accessGrant.ServiceName, accessGrant.OrganizationName, id)

	accessGrant.ID = id
	accessGrant.CreatedAt = t

	if err := db.put(ctx, key, accessGrant); err != nil {
		return nil, err
	}

	return accessGrant, nil
}

func (db ETCDConfigDatabase) ListAccessGrantsForService(ctx context.Context, serviceName string) ([]*AccessGrant, error) {
	key := path.Join("access-grants", serviceName)

	r := []*AccessGrant{}

	err := db.list(ctx, key, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
