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

	locksPrefix = "/nlx/locks/"
)

var (
	ErrActiveAccessRequest = errors.New("already active access request")
	ErrAccessRequestLocked = errors.New("access request is already locked")

	AccessRequestLockTTL = func() int64 {
		return 60 * 5
	}
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

	return accessRequest, nil
}

func (db ETCDConfigDatabase) UpdateAccessRequestState(ctx context.Context, accessRequest *AccessRequest, state AccessRequestState) error {
	key := path.Join("access-requests", "outgoing", accessRequest.OrganizationName, accessRequest.ServiceName, accessRequest.ID)

	response, err := db.etcdCli.Get(ctx, db.key(key))
	if err != nil {
		return err
	}

	if response.Count == 0 {
		return fmt.Errorf("no such access request: %s", accessRequest.ID)
	}

	accessRequest.State = state
	accessRequest.UpdatedAt = db.clock.Now()

	if err := db.put(ctx, key, accessRequest); err != nil {
		return err
	}

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
	key := db.key(
		"locks",
		"access-requests",
		"outgoing",
		accessRequest.OrganizationName,
		accessRequest.ServiceName,
		accessRequest.ID,
	)

	leaseResponse, err := db.etcdCli.Grant(ctx, AccessRequestLockTTL())
	if err != nil {
		return err
	}

	response, err := db.etcdCli.Txn(ctx).
		If(clientv3.Compare(clientv3.CreateRevision(key), "=", 0)).
		Then(clientv3.OpPut(key, "", clientv3.WithLease(leaseResponse.ID))).
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
	key := db.key(
		"locks",
		"access-requests",
		"outgoing",
		accessRequest.OrganizationName,
		accessRequest.ServiceName,
		accessRequest.ID,
	)

	_, err := db.etcdCli.Delete(ctx, key)

	return err
}

func (db ETCDConfigDatabase) parseDeleteAccessRequestLeaseEvent(ctx context.Context, event *clientv3.Event) (*AccessRequest, error) {
	if event.Kv == nil {
		return nil, nil
	}

	leaseID := clientv3.LeaseID(event.Kv.Lease)

	if leaseID != clientv3.NoLease {
		ttlResponse, err := db.etcdCli.TimeToLive(ctx, leaseID)
		if err != nil {
			return nil, err
		}

		if ttlResponse.TTL <= 0 {
			accessRequestKey := string(event.Kv.Key)[len(locksPrefix):]

			request := &AccessRequest{}
			if err := db.get(ctx, accessRequestKey, request); err != nil {
				return nil, err
			}

			if request.ID == "" {
				return nil, nil
			}

			return request, nil
		}
	}

	return nil, nil
}

func (db ETCDConfigDatabase) parseCreateAccessRequestEvent(event *clientv3.Event) (*AccessRequest, error) {
	if !event.IsCreate() {
		return nil, nil
	}

	request := &AccessRequest{}

	if err := json.Unmarshal(event.Kv.Value, request); err != nil {
		return nil, err
	}

	return request, nil
}

func (db ETCDConfigDatabase) WatchOutgoingAccessRequests(ctx context.Context, output chan *AccessRequest) {
	leaderCtx := clientv3.WithRequireLeader(ctx)

	accessRequestsChannel := db.etcdCli.Watch(
		leaderCtx,
		db.key("access-requests", "outgoing"),
		clientv3.WithPrefix(),
	)
	locksChannel := db.etcdCli.Watch(
		leaderCtx,
		db.key("locks", "access-requests", "outgoing"),
		clientv3.WithPrefix(),
	)

	for {
		select {
		// re-schedule dead-locked AccessRequests due to a system failure
		case response := <-locksChannel:
			for _, event := range response.Events {
				request, err := db.parseDeleteAccessRequestLeaseEvent(ctx, event)
				if err != nil {
					db.logger.Error("failed to parse delete lease event", zap.Error(err))
					continue
				}

				if request != nil {
					output <- request
				}
			}
			// filter AccessRequests on create and send to the output channel
		case response := <-accessRequestsChannel:
			for _, event := range response.Events {
				request, err := db.parseCreateAccessRequestEvent(event)
				if err != nil {
					db.logger.Error("failed to parse create access request event", zap.Error(err))
					continue
				}

				if request != nil {
					output <- request
				}
			}
		}
	}
}
