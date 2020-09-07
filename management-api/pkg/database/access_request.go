// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type AccessRequest struct {
	ID               string             `json:"id,omitempty"`
	OrganizationName string             `json:"organizationName,omitempty"`
	ServiceName      string             `json:"serviceName,omitempty"`
	State            AccessRequestState `json:"state"`
	CreatedAt        time.Time          `json:"createdAt,omitempty"`
	UpdatedAt        time.Time          `json:"updatedAt,omitempty"`
}

type AccessRequestState int

const (
	AccessRequestFailed AccessRequestState = iota
	AccessRequestCreated
	AccessRequestReceived
)

var ErrActiveAccessRequest = errors.New("already active access request")

func (db ETCDConfigDatabase) ListAllOutgoingAccessRequests(ctx context.Context) ([]*AccessRequest, error) {
	key := path.Join("access-requests", "outgoing")

	r := []*AccessRequest{}

	err := db.list(ctx, key, &r, clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
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
	existing, err := db.GetLatestOutgoingAccessRequest(ctx, accessRequest.OrganizationName, accessRequest.ServiceName)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, ErrActiveAccessRequest
	}

	t := db.clock.Now()
	id := fmt.Sprintf("%x", t.UnixNano())

	key := path.Join("access-requests", "outgoing", accessRequest.OrganizationName, accessRequest.ServiceName, id)

	accessRequest.ID = id
	accessRequest.State = AccessRequestCreated
	accessRequest.CreatedAt = t
	accessRequest.UpdatedAt = t

	err = db.put(ctx, key, accessRequest)
	if err != nil {
		return nil, err
	}

	return accessRequest, nil
}

func (db ETCDConfigDatabase) UpdateAccessRequestState(ctx context.Context, accessRequest *AccessRequest, state AccessRequestState) error {
	key := path.Join("access-requests", "outgoing", accessRequest.OrganizationName, accessRequest.ServiceName, accessRequest.ID)

	request := &*accessRequest
	request.State = state
	request.UpdatedAt = db.clock.Now()

	response, err := db.etcdCli.Get(ctx, db.key(key))
	if err != nil {
		return err
	}

	if response.Count == 0 {
		return fmt.Errorf("no such access request: %s", accessRequest.ID)
	}

	if err := db.put(ctx, key, request); err != nil {
		return err
	}

	accessRequest.State = request.State
	accessRequest.UpdatedAt = request.UpdatedAt

	return nil
}

func (db ETCDConfigDatabase) GetLatestOutgoingAccessRequest(ctx context.Context, organizationName, serviceName string) (*AccessRequest, error) {
	var r *AccessRequest

	key := path.Join("access-requests", "outgoing", organizationName, serviceName)

	err := db.get(ctx, key, &r, clientv3.WithLastKey()...)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db ETCDConfigDatabase) ListAllLatestOutgoingAccessRequests(ctx context.Context) (map[string]*AccessRequest, error) {
	accessRequests, err := db.ListAllOutgoingAccessRequests(ctx)
	if err != nil {
		return nil, err
	}

	latestAccessRequests := make(map[string]*AccessRequest)

	for _, a := range accessRequests {
		key := path.Join(a.OrganizationName, a.ServiceName)
		if _, ok := latestAccessRequests[key]; !ok {
			latestAccessRequests[key] = a
		}
	}

	return latestAccessRequests, nil
}
