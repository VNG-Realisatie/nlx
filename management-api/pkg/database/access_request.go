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

	"go.nlx.io/nlx/common/diagnostics"
)

type AccessRequest struct {
	ID                   string             `json:"id,omitempty"`
	OrganizationName     string             `json:"organizationName,omitempty"`
	ServiceName          string             `json:"serviceName,omitempty"`
	State                AccessRequestState `json:"state"`
	PublicKeyFingerprint string             `json:"publicKeyFingerprint"`
	CreatedAt            time.Time          `json:"createdAt,omitempty"`
	UpdatedAt            time.Time          `json:"updatedAt,omitempty"`
}

type IncomingAccessRequest struct {
	AccessRequest
}

type OutgoingAccessRequest struct {
	AccessRequest
	ErrorDetails *diagnostics.ErrorDetails
	ReferenceID  string `json:"referenceId"`
}

// Sendable returns if the access request can be send to the other organization
func (a *OutgoingAccessRequest) Sendable() bool {
	return a.State == AccessRequestCreated ||
		a.State == AccessRequestFailed
}

type AccessRequestState int

const (
	AccessRequestUnspecified AccessRequestState = iota
	AccessRequestFailed
	AccessRequestCreated
	AccessRequestReceived
	AccessRequestApproved
	AccessRequestRejected
)

func (state AccessRequestState) String() string {
	switch state {
	case AccessRequestFailed:
		return "Failed"
	case AccessRequestCreated:
		return "Created"
	case AccessRequestReceived:
		return "Received"
	case AccessRequestApproved:
		return "Approved"
	case AccessRequestRejected:
		return "Rejected"
	}

	return "Unspecified"
}

const locksPrefix = "/nlx/locks/"

var (
	ErrActiveAccessRequest = errors.New("already active access request")
	ErrAccessRequestLocked = errors.New("access request is already locked")

	AccessRequestLockTTL = func() int64 {
		return 60 * 5
	}
)

func (db ETCDConfigDatabase) parseDeleteAccessRequestLeaseEvent(ctx context.Context, event *clientv3.Event) (*OutgoingAccessRequest, error) {
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

			request := &OutgoingAccessRequest{}
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

func (db ETCDConfigDatabase) parseCreateAccessRequestEvent(event *clientv3.Event) (*OutgoingAccessRequest, error) {
	request := &OutgoingAccessRequest{}

	if err := json.Unmarshal(event.Kv.Value, request); err != nil {
		return nil, err
	}

	return request, nil
}

func (db ETCDConfigDatabase) listAccessRequests(ctx context.Context, prefix string, requests interface{}) error {
	key := path.Join("access-requests", prefix)

	err := db.list(ctx, key, requests, clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return err
	}

	return nil
}

func (db ETCDConfigDatabase) getLatestAccessRequest(ctx context.Context, prefix, organizationName, serviceName string, request interface{}) error {
	key := path.Join("access-requests", prefix, organizationName, serviceName)

	err := db.get(ctx, key, request, clientv3.WithLastKey()...)
	if err != nil {
		return err
	}

	return nil
}

func (db ETCDConfigDatabase) createAccessRequest(ctx context.Context, prefix string, accessRequest *AccessRequest) error {
	t := db.clock.Now()
	id := fmt.Sprintf("%x", t.UnixNano())

	key := path.Join("access-requests", prefix, accessRequest.OrganizationName, accessRequest.ServiceName, id)

	accessRequest.ID = id
	accessRequest.CreatedAt = t
	accessRequest.UpdatedAt = t

	if err := db.put(ctx, key, accessRequest); err != nil {
		return err
	}

	return nil
}

func (db ETCDConfigDatabase) ListAllOutgoingAccessRequests(ctx context.Context) (requests []*OutgoingAccessRequest, err error) {
	err = db.listAccessRequests(ctx, "outgoing", &requests)
	return
}

func (db ETCDConfigDatabase) ListOutgoingAccessRequests(ctx context.Context, organizationName, serviceName string) (requests []*OutgoingAccessRequest, err error) {
	err = db.listAccessRequests(ctx, path.Join("outgoing", organizationName, serviceName), &requests)
	return
}

func (db ETCDConfigDatabase) GetOutgoingAccessRequest(ctx context.Context, id string) (*OutgoingAccessRequest, error) {
	key := path.Join("access-requests", "outgoing")

	var result []*OutgoingAccessRequest

	err := db.list(ctx, key, &result)
	if err != nil {
		return nil, err
	}

	var accessRequest *OutgoingAccessRequest

	for _, item := range result {
		if item.ID == id {
			accessRequest = item
			break
		}
	}

	if accessRequest == nil {
		return nil, ErrNotFound
	}

	return accessRequest, nil
}

func (db ETCDConfigDatabase) GetLatestOutgoingAccessRequest(ctx context.Context, organizationName, serviceName string) (request *OutgoingAccessRequest, err error) {
	err = db.getLatestAccessRequest(ctx, "outgoing", organizationName, serviceName, &request)
	return
}

func (db ETCDConfigDatabase) ListAllLatestOutgoingAccessRequests(ctx context.Context) (map[string]*OutgoingAccessRequest, error) {
	accessRequests, err := db.ListAllOutgoingAccessRequests(ctx)
	if err != nil {
		return nil, err
	}

	latestAccessRequests := make(map[string]*OutgoingAccessRequest)

	for _, a := range accessRequests {
		key := path.Join(a.OrganizationName, a.ServiceName)
		if _, ok := latestAccessRequests[key]; !ok {
			latestAccessRequests[key] = a
		}
	}

	return latestAccessRequests, nil
}

func (db ETCDConfigDatabase) LockOutgoingAccessRequest(ctx context.Context, accessRequest *OutgoingAccessRequest) error {
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

func (db ETCDConfigDatabase) UnlockOutgoingAccessRequest(ctx context.Context, accessRequest *OutgoingAccessRequest) error {
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

// nolint:dupl,gocyclo // verification is very similar for outgoing and incoming access requests
func (db ETCDConfigDatabase) verifyOutgoingAccessRequestUniqueConstraint(
	ctx context.Context,
	organizationName,
	serviceName,
	publicKeyFingerprint string,
) error {
	proof, err := db.GetLatestAccessProofForService(ctx, organizationName, serviceName)

	if errors.Is(err, ErrNotFound) {
		requests := []*AccessRequest{}

		err = db.listAccessRequests(ctx, path.Join("outgoing", organizationName, serviceName), &requests)
		if err != nil {
			return err
		}

		for _, request := range requests {
			if request.PublicKeyFingerprint == publicKeyFingerprint &&
				request.State != AccessRequestRejected {
				return ErrActiveAccessRequest
			}
		}

		return nil
	}

	if err != nil {
		return err
	}

	if !proof.RevokedAt.IsZero() {
		request, err := db.GetLatestOutgoingAccessRequest(ctx, organizationName, serviceName)
		if err != nil {
			return err
		}

		if request.ID == proof.AccessRequestID ||
			request.State == AccessRequestRejected {
			return nil
		}
	}

	return ErrActiveAccessRequest
}

//nolint:dupl,gocyclo // incoming and outgoing logic is very similar - but not the same
func (db ETCDConfigDatabase) verifyIncomingAccessRequestUniqueConstraint(
	ctx context.Context,
	organizationName,
	serviceName,
	publicKeyFingerprint string,
) error {
	grant, err := db.GetLatestAccessGrantForService(ctx, organizationName, serviceName)

	if errors.Is(err, ErrNotFound) {
		requests := []*AccessRequest{}

		err = db.listAccessRequests(ctx, path.Join("incoming", organizationName, serviceName), &requests)
		if err != nil {
			return err
		}

		for _, request := range requests {
			if request.PublicKeyFingerprint == publicKeyFingerprint &&
				request.State != AccessRequestRejected {
				return ErrActiveAccessRequest
			}
		}

		return nil
	}

	if err != nil {
		return err
	}

	if !grant.RevokedAt.IsZero() {
		request, err := db.GetLatestIncomingAccessRequest(ctx, organizationName, serviceName)
		if err != nil {
			return err
		}

		if request.ID == grant.AccessRequestID ||
			request.State == AccessRequestRejected {
			return nil
		}
	}

	return ErrActiveAccessRequest
}

func (db ETCDConfigDatabase) CreateOutgoingAccessRequest(ctx context.Context, accessRequest *OutgoingAccessRequest) (*OutgoingAccessRequest, error) {
	if err := db.verifyOutgoingAccessRequestUniqueConstraint(
		ctx,
		accessRequest.OrganizationName,
		accessRequest.ServiceName,
		accessRequest.PublicKeyFingerprint,
	); err != nil {
		return nil, err
	}

	accessRequest.State = AccessRequestCreated

	if err := db.createAccessRequest(ctx, "outgoing", &accessRequest.AccessRequest); err != nil {
		return nil, err
	}

	return accessRequest, nil
}

// nolint:gocyclo // it's indeed somewhat complex but we do want to have these sanity checks in place
func (db ETCDConfigDatabase) UpdateOutgoingAccessRequestState(ctx context.Context, accessRequest *OutgoingAccessRequest, state AccessRequestState, referenceID string, errDetails *diagnostics.ErrorDetails) error {
	key := path.Join("access-requests", "outgoing", accessRequest.OrganizationName, accessRequest.ServiceName, accessRequest.ID)

	response, err := db.etcdCli.Get(ctx, db.key(key))
	if err != nil {
		return err
	}

	if response.Count == 0 {
		return fmt.Errorf("no such access request: %s", accessRequest.ID)
	}

	accessRequest.ErrorDetails = errDetails
	accessRequest.UpdatedAt = db.clock.Now()

	switch state {
	case AccessRequestFailed:
		if errDetails == nil {
			return errors.New("unable to update AccessRequest state to failed without error details")
		}
	case AccessRequestReceived:
		if referenceID == "" {
			return fmt.Errorf("unable to update AccessRequest to state received without referenceID")
		}

		accessRequest.ReferenceID = referenceID
	default:
		if referenceID != "" {
			return fmt.Errorf("unable to set referenceID when state is %s", state.String())
		}

		if errDetails != nil {
			return fmt.Errorf("unable to update AccessRequest state to :%s while error details are provided", state.String())
		}
	}

	if err := db.put(ctx, key, accessRequest); err != nil {
		return err
	}

	return nil
}

func (db ETCDConfigDatabase) WatchOutgoingAccessRequests(ctx context.Context, output chan *OutgoingAccessRequest) {
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

				if request.State == AccessRequestCreated {
					output <- request
				}
			}
		}
	}
}

func (db ETCDConfigDatabase) ListAllIncomingAccessRequests(ctx context.Context) (requests []*IncomingAccessRequest, err error) {
	err = db.listAccessRequests(ctx, "incoming", &requests)
	return
}

func (db ETCDConfigDatabase) ListIncomingAccessRequests(ctx context.Context, organizationName, serviceName string) (requests []*IncomingAccessRequest, err error) {
	err = db.listAccessRequests(ctx, path.Join("incoming", organizationName, serviceName), &requests)
	return
}

func (db ETCDConfigDatabase) GetLatestIncomingAccessRequest(ctx context.Context, organizationName, serviceName string) (request *IncomingAccessRequest, err error) {
	err = db.getLatestAccessRequest(ctx, "incoming", organizationName, serviceName, &request)
	return
}

func (db ETCDConfigDatabase) ListAllLatestIncomingAccessRequests(ctx context.Context) (map[string]*IncomingAccessRequest, error) {
	accessRequests, err := db.ListAllIncomingAccessRequests(ctx)
	if err != nil {
		return nil, err
	}

	latestAccessRequests := make(map[string]*IncomingAccessRequest)

	for _, a := range accessRequests {
		key := path.Join(a.OrganizationName, a.ServiceName)
		if _, ok := latestAccessRequests[key]; !ok {
			latestAccessRequests[key] = a
		}
	}

	return latestAccessRequests, nil
}

func (db ETCDConfigDatabase) GetIncomingAccessRequest(ctx context.Context, id string) (*IncomingAccessRequest, error) {
	key := path.Join("access-requests", "incoming")

	var result []*IncomingAccessRequest

	err := db.list(ctx, key, &result)
	if err != nil {
		return nil, err
	}

	var accessRequest *IncomingAccessRequest

	for _, item := range result {
		if item.ID == id {
			accessRequest = item
			break
		}
	}

	if accessRequest == nil {
		return nil, ErrNotFound
	}

	return accessRequest, nil
}

func (db ETCDConfigDatabase) CreateIncomingAccessRequest(ctx context.Context, accessRequest *IncomingAccessRequest) (*IncomingAccessRequest, error) {
	if err := db.verifyIncomingAccessRequestUniqueConstraint(
		ctx,
		accessRequest.OrganizationName,
		accessRequest.ServiceName,
		accessRequest.PublicKeyFingerprint,
	); err != nil {
		return nil, err
	}

	accessRequest.State = AccessRequestReceived

	if err := db.createAccessRequest(ctx, "incoming", &accessRequest.AccessRequest); err != nil {
		return nil, err
	}

	return accessRequest, nil
}

func (db ETCDConfigDatabase) UpdateIncomingAccessRequestState(ctx context.Context, accessRequest *IncomingAccessRequest, state AccessRequestState) error {
	key := path.Join("access-requests", "incoming", accessRequest.OrganizationName, accessRequest.ServiceName, accessRequest.ID)

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
