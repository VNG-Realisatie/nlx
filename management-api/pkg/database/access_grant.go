// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"time"

	"github.com/pkg/errors"

	"github.com/coreos/etcd/clientv3"
)

var ErrAccessRequestModified = errors.New("access request modified")

type AccessGrant struct {
	ID                   string    `json:"id,omitempty"`
	AccessRequestID      string    `json:"accessRequestId,omitempty"`
	OrganizationName     string    `json:"organizationName,omitempty"`
	ServiceName          string    `json:"serviceName,omitempty"`
	PublicKeyFingerprint string    `json:"publicKeyFingerprint,omitempty"`
	CreatedAt            time.Time `json:"createdAt,omitempty"`
}

func (db ETCDConfigDatabase) CreateAccessGrant(ctx context.Context, accessRequest *IncomingAccessRequest) (*AccessGrant, error) {
	data, err := json.Marshal(accessRequest)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling access request")
	}

	key := db.key("access-requests", "incoming", accessRequest.OrganizationName, accessRequest.ServiceName, accessRequest.ID)

	accessRequestCompare := clientv3.Compare(
		clientv3.Value(key), "=", string(data),
	)

	accessRequest.State = AccessRequestApproved
	accessRequest.UpdatedAt = db.clock.Now()

	data, err = json.Marshal(accessRequest)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling approved access request")
	}

	accessRequestOp := clientv3.OpPut(key, string(data))

	now := db.clock.Now()
	id := fmt.Sprintf("%x", now.UnixNano())

	accessGrant := &AccessGrant{
		ID:                   id,
		AccessRequestID:      accessRequest.ID,
		OrganizationName:     accessRequest.OrganizationName,
		ServiceName:          accessRequest.ServiceName,
		PublicKeyFingerprint: accessRequest.PublicKeyFingerprint,
		CreatedAt:            now,
	}

	key = db.key("access-grants", accessGrant.ServiceName, accessGrant.OrganizationName, id)

	data, err = json.Marshal(accessGrant)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling access grant")
	}

	accessGrantOp := clientv3.OpPut(key, string(data))

	transaction := db.etcdCli.KV.Txn(ctx).
		If(accessRequestCompare).
		Then(accessRequestOp, accessGrantOp)

	response, err := transaction.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "committing transaction")
	}

	if !response.Succeeded {
		return nil, ErrAccessRequestModified
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
