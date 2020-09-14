// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

// Snippet for generating an id for an AccessRuquest:
//
//  func TestId(t *testing.T) {
//  	fmt.Printf("%x\n", time.Date(2020, time.July, 9, 14, 45, 0, 0, time.UTC).UnixNano())
//  	t.Fail()
//  }
//
package database_test

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestAccessRequest(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()
	client := cluster.GetClient(t)

	createAccessRequest := func(id, organization, service string) {
		a, _ := json.Marshal(
			&database.AccessRequest{
				OrganizationName: organization,
				ServiceName:      service,
			},
		)

		_, err := client.Put(ctx, path.Join("/nlx/access-requests/outgoing", organization, service, id), string(a))
		assert.NoError(t, err)
	}

	createAccessRequest("1", "test-organization-a", "test-service-1")
	createAccessRequest("2", "test-organization-a", "test-service-1")
	createAccessRequest("3", "test-organization-a", "test-service-2")
	createAccessRequest("4", "test-organization-b", "test-service-1")

	actual, err := cluster.DB.ListOutgoingAccessRequests(ctx, "test-organization-a", "test-service-1")

	assert.NoError(t, err)
	assert.Len(t, actual, 2)

	expected := []*database.AccessRequest{
		{
			OrganizationName: "test-organization-a",
			ServiceName:      "test-service-1",
		},
		{
			OrganizationName: "test-organization-a",
			ServiceName:      "test-service-1",
		},
	}

	assert.Equal(t, expected, actual)
}

func TestCreateAccessRequest(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()
	client := cluster.GetClient(t)

	a := &database.AccessRequest{
		OrganizationName: "test-organization-a",
		ServiceName:      "test-service-1",
	}

	actual, err := cluster.DB.CreateAccessRequest(ctx, a)
	assert.NoError(t, err)

	expected := &database.AccessRequest{
		ID:               "161c188cfcea1939",
		OrganizationName: "test-organization-a",
		ServiceName:      "test-service-1",
		State:            database.AccessRequestCreated,
		CreatedAt:        time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
		UpdatedAt:        time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
	}

	assert.Equal(t, expected, actual)

	response, err := client.Get(ctx, "/nlx/access-requests/outgoing/test-organization-a/test-service-1/161c188cfcea1939")
	assert.NoError(t, err)
	assert.Len(t, response.Kvs, 1)
}

func TestUpdateAccessRequestState(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()
	client := cluster.GetClient(t)

	a := &database.AccessRequest{
		OrganizationName: "test-organization-a",
		ServiceName:      "test-service-1",
		State:            database.AccessRequestCreated,
	}

	actual, err := cluster.DB.CreateAccessRequest(ctx, a)
	assert.NoError(t, err)

	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 43, 1337, time.UTC))

	err = cluster.DB.UpdateAccessRequestState(ctx, a, database.AccessRequestFailed)
	assert.NoError(t, err)

	expected := &database.AccessRequest{
		ID:               "161c188cfcea1939",
		OrganizationName: "test-organization-a",
		ServiceName:      "test-service-1",
		State:            database.AccessRequestFailed,
		CreatedAt:        time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
		UpdatedAt:        time.Date(2020, time.June, 26, 12, 42, 43, 1337, time.UTC),
	}

	assert.Equal(t, expected, actual)

	response, err := client.Get(ctx, "/nlx/access-requests/outgoing/test-organization-a/test-service-1/161c188cfcea1939")
	assert.NoError(t, err)
	assert.Len(t, response.Kvs, 1)
}

func TestLockOutgoingAccessRequest(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()

	a := &database.AccessRequest{
		OrganizationName: "test-organization-a",
		ServiceName:      "test-service-1",
	}

	_, err := cluster.DB.CreateAccessRequest(ctx, a)
	assert.NoError(t, err)

	err = cluster.DB.LockOutgoingAccessRequest(ctx, a)
	assert.NoError(t, err)

	err = cluster.DB.LockOutgoingAccessRequest(ctx, a)
	assert.Error(t, err)
	assert.Equal(t, database.ErrAccessRequestLocked, err)
}

func TestUnlockOutgoingAccessRequest(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()

	a := &database.AccessRequest{
		OrganizationName: "test-organization-a",
		ServiceName:      "test-service-1",
	}

	_, err := cluster.DB.CreateAccessRequest(ctx, a)
	assert.NoError(t, err)

	err = cluster.DB.LockOutgoingAccessRequest(ctx, a)
	assert.NoError(t, err)

	err = cluster.DB.LockOutgoingAccessRequest(ctx, a)
	assert.Error(t, err)
	assert.Equal(t, database.ErrAccessRequestLocked, err)

	err = cluster.DB.UnlockOutgoingAccessRequest(ctx, a)
	assert.NoError(t, err)

	err = cluster.DB.LockOutgoingAccessRequest(ctx, a)
	assert.NoError(t, err)
}

func TestGetLatestOutgoingAccessRequest(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	cluster.Clock.SetTime(time.Date(2020, time.July, 9, 14, 44, 55, 0, time.UTC))

	ctx := context.Background()

	create := func(o, s string) {
		cluster.Clock.Step(5 * time.Second)

		_, err := cluster.DB.CreateAccessRequest(ctx, &database.AccessRequest{
			OrganizationName: o,
			ServiceName:      s,
		})

		assert.NoError(t, err)
	}

	create("test-organization-a", "test-service-1") // 14:45:00

	expected := &database.AccessRequest{
		ID:               "16201cc4e0cf3800",
		OrganizationName: "test-organization-a",
		ServiceName:      "test-service-1",
		State:            database.AccessRequestCreated,
		CreatedAt:        time.Date(2020, time.July, 9, 14, 45, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, time.July, 9, 14, 45, 0, 0, time.UTC),
	}

	actual, err := cluster.DB.GetLatestOutgoingAccessRequest(ctx, "test-organization-a", "test-service-1")

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestListAllLatestOutgoingAccessRequests(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	cluster.Clock.SetTime(time.Date(2020, time.July, 10, 10, 11, 40, 0, time.UTC))

	ctx := context.Background()

	create := func(o, s string) {
		cluster.Clock.Step(5 * time.Second)

		r, err := cluster.DB.CreateAccessRequest(ctx, &database.AccessRequest{
			OrganizationName: o,
			ServiceName:      s,
		})

		fmt.Println(r)

		assert.NoError(t, err)
	}

	create("test-organization-b", "test-service-1") // 10:11:45
	create("test-organization-c", "test-service-1") // 10:11:50
	create("test-organization-a", "test-service-2") // 10:12:55

	expected := map[string]*database.AccessRequest{
		"test-organization-a/test-service-2": {
			ID:               "16205c7284036e00",
			OrganizationName: "test-organization-a",
			ServiceName:      "test-service-2",
			State:            database.AccessRequestCreated,
			CreatedAt:        time.Date(2020, time.July, 10, 10, 11, 55, 0, time.UTC),
			UpdatedAt:        time.Date(2020, time.July, 10, 10, 11, 55, 0, time.UTC),
		},
		"test-organization-b/test-service-1": {
			ID:               "16205c702ff78a00",
			OrganizationName: "test-organization-b",
			ServiceName:      "test-service-1",
			State:            database.AccessRequestCreated,
			CreatedAt:        time.Date(2020, time.July, 10, 10, 11, 45, 0, time.UTC),
			UpdatedAt:        time.Date(2020, time.July, 10, 10, 11, 45, 0, time.UTC),
		},
		"test-organization-c/test-service-1": {
			ID:               "16205c7159fd7c00",
			OrganizationName: "test-organization-c",
			ServiceName:      "test-service-1",
			State:            database.AccessRequestCreated,
			CreatedAt:        time.Date(2020, time.July, 10, 10, 11, 50, 0, time.UTC),
			UpdatedAt:        time.Date(2020, time.July, 10, 10, 11, 50, 0, time.UTC),
		},
	}

	actual, err := cluster.DB.ListAllLatestOutgoingAccessRequests(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestWatchOutgoingAccessRequests(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	cluster.Clock.SetTime(time.Date(2020, time.July, 10, 10, 11, 40, 0, time.UTC))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	request1 := &database.AccessRequest{
		ID:               "id-1",
		OrganizationName: "organization-a",
		ServiceName:      "service-x",
		State:            database.AccessRequestCreated,
		CreatedAt:        time.Date(2020, time.July, 10, 10, 11, 40, 0, time.UTC),
		UpdatedAt:        time.Date(2020, time.July, 10, 10, 11, 40, 0, time.UTC),
	}
	request2 := &database.AccessRequest{
		ID:               "id-2",
		OrganizationName: "organization-b",
		ServiceName:      "service-y",
		State:            database.AccessRequestCreated,
		CreatedAt:        time.Date(2020, time.July, 10, 10, 11, 45, 0, time.UTC),
		UpdatedAt:        time.Date(2020, time.July, 10, 10, 11, 45, 0, time.UTC),
	}
	expected := []*database.AccessRequest{request1, request2, nil}

	database.AccessRequestLockTTL = func() int64 {
		return 0
	}

	_, err := cluster.DB.CreateAccessRequest(ctx, request1)
	assert.NoError(t, err)

	requests := make(chan *database.AccessRequest, 3)

	go func() {
		<-ctx.Done()
		close(requests)
	}()

	go func() {
		cluster.DB.WatchOutgoingAccessRequests(ctx, requests)
	}()

	err = cluster.DB.LockOutgoingAccessRequest(ctx, request1)
	assert.NoError(t, err)

	cluster.Clock.Step(5 * time.Second)

	database.AccessRequestLockTTL = func() int64 {
		return 300
	}

	_, err = cluster.DB.CreateAccessRequest(ctx, request2)
	assert.NoError(t, err)

	err = cluster.DB.LockOutgoingAccessRequest(ctx, request2)
	assert.NoError(t, err)

	err = cluster.DB.UnlockOutgoingAccessRequest(ctx, request2)
	assert.NoError(t, err)

	actual := []*database.AccessRequest{}

	for i := 0; i < len(expected); i++ {
		actual = append(actual, <-requests)
	}

	assert.Equal(t, expected, actual)
}
