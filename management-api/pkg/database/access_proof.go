// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/pkg/errors"

	"github.com/coreos/etcd/clientv3"
)

var ErrAccessProofAlreadyRevoked = errors.New("access proof is already revoked")

type AccessProof struct {
	ID               string    `json:"id,omitempty"`
	AccessRequestID  string    `json:"accessRequestId,omitempty"`
	OrganizationName string    `json:"organizationName,omitempty"`
	ServiceName      string    `json:"serviceName,omitempty"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	RevokedAt        time.Time `json:"revokedAt,omitempty"`
}

func (a *AccessProof) Revoked() bool {
	return !a.RevokedAt.IsZero()
}

func (db ETCDConfigDatabase) CreateAccessProof(ctx context.Context, accessProof *AccessProof) (*AccessProof, error) {
	now := db.clock.Now()
	id := fmt.Sprintf("%x", now.UnixNano())

	accessProof = &AccessProof{
		ID:               id,
		AccessRequestID:  accessProof.AccessRequestID,
		OrganizationName: accessProof.OrganizationName,
		ServiceName:      accessProof.ServiceName,
		CreatedAt:        now,
		RevokedAt:        accessProof.RevokedAt,
	}

	key := path.Join("access-proofs", accessProof.OrganizationName, accessProof.ServiceName, id)

	if err := db.put(ctx, key, accessProof); err != nil {
		return nil, err
	}

	return accessProof, nil
}

func (db ETCDConfigDatabase) RevokeAccessProof(ctx context.Context, organizationName, serviceName, id string, revokedAt time.Time) (*AccessProof, error) {
	key := path.Join("access-proofs", organizationName, serviceName, id)

	accessProof := &AccessProof{}

	err := db.get(ctx, key, accessProof)
	if err != nil {
		return nil, err
	}

	if accessProof.Revoked() {
		return nil, ErrAccessProofAlreadyRevoked
	}

	accessProof.RevokedAt = revokedAt

	err = db.put(ctx, key, accessProof)
	if err != nil {
		return nil, err
	}

	return accessProof, nil
}

func (db ETCDConfigDatabase) GetLatestAccessProofForService(ctx context.Context, organizationName, serviceName string) (*AccessProof, error) {
	key := path.Join("access-proofs", organizationName, serviceName)

	accessProof := &AccessProof{}

	err := db.get(ctx, key, accessProof, clientv3.WithLastKey()...)
	if err != nil {
		return nil, err
	}

	return accessProof, nil
}
