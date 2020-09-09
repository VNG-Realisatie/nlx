// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/coreos/etcd/clientv3"
	"go.uber.org/zap"
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

var (
	ErrActiveAccessRequest = errors.New("already active access request")
	ErrAccessRequestLocked = errors.New("access request is already locked")
)

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

	if err := db.put(ctx, key, accessRequest); err != nil {
		return nil, err
	}

	if _, err := db.etcdCli.Put(ctx, path.Join(key, "locked"), "false"); err != nil {
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

func (db ETCDConfigDatabase) LockOutgoingAccessRequest(ctx context.Context, accessRequest *AccessRequest) error {
	key := path.Join(
		"access-requests",
		"outgoing",
		accessRequest.OrganizationName,
		accessRequest.ServiceName,
		accessRequest.ID,
		"locked",
	)

	response, err := db.etcdCli.Txn(ctx).
		If(clientv3.Compare(clientv3.Value(key), "=", "false")).
		Then(clientv3.OpPut(key, "true")).
		Commit()
	if err != nil {
		return err
	}

	if !response.Succeeded {
		return ErrAccessRequestLocked
	}

	return nil
}

func (db ETCDConfigDatabase) UnlockOutgoingAccessRequest(ctx context.Context, accessRequest *AccessRequest) error {
	key := path.Join(
		"access-requests",
		"outgoing",
		accessRequest.OrganizationName,
		accessRequest.ServiceName,
		accessRequest.ID,
		"locked",
	)

	_, err := db.etcdCli.Put(ctx, key, "false")
	return err
}

func (db ETCDConfigDatabase) WatchOutgoingAccessRequests(ctx context.Context, output chan *AccessRequest) error {
	key := path.Join("access-requests", "outgoing")
	watchChannel := db.etcdCli.Watch(ctx, key)

	for response := range watchChannel {
		for _, event := range response.Events {
			if event.IsCreate() {
				request := &AccessRequest{}

				if err := json.Unmarshal(event.Kv.Value, request); err != nil {
					db.logger.Error("failed to unmarshal created AccessRequest", zap.Error(err))
					continue
				}

				output <- request
			}
		}
	}

	return nil
}
