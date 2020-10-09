// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

// Snippet for generating an id for an AccessRequest:
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
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestListAccessRequests(t *testing.T) {
	cluster := newTestCluster(t)
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

	expected := []*database.OutgoingAccessRequest{
		{
			AccessRequest: database.AccessRequest{
				OrganizationName: "test-organization-a",
				ServiceName:      "test-service-1",
			},
		},
		{
			AccessRequest: database.AccessRequest{
				OrganizationName: "test-organization-a",
				ServiceName:      "test-service-1",
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestCreateAccessRequest(t *testing.T) {
	cluster := newTestCluster(t)
	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()
	client := cluster.GetClient(t)

	a := &database.OutgoingAccessRequest{
		AccessRequest: database.AccessRequest{
			OrganizationName:     "test-organization-a",
			ServiceName:          "test-service-1",
			PublicKeyFingerprint: "public_key",
		},
	}

	actual, err := cluster.DB.CreateOutgoingAccessRequest(ctx, a)
	assert.NoError(t, err)

	expected := &database.OutgoingAccessRequest{
		AccessRequest: database.AccessRequest{
			ID:                   "161c188cfcea1939",
			OrganizationName:     "test-organization-a",
			ServiceName:          "test-service-1",
			PublicKeyFingerprint: "public_key",
			State:                database.AccessRequestCreated,
			CreatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
			UpdatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
		},
	}

	assert.Equal(t, expected, actual)

	response, err := client.Get(ctx, "/nlx/access-requests/outgoing/test-organization-a/test-service-1/161c188cfcea1939")
	assert.NoError(t, err)
	assert.Len(t, response.Kvs, 1)

	a = &database.OutgoingAccessRequest{
		AccessRequest: database.AccessRequest{
			OrganizationName:     "test-organization-a",
			ServiceName:          "test-service-1",
			PublicKeyFingerprint: "public_key",
		},
	}

	_, err = cluster.DB.CreateOutgoingAccessRequest(ctx, a)

	assert.Error(t, err)
	assert.Equal(t, err, database.ErrActiveAccessRequest)
}

func TestUpdateAccessRequestState(t *testing.T) {
	cluster := newTestCluster(t)
	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()
	client := cluster.GetClient(t)

	a := &database.OutgoingAccessRequest{
		AccessRequest: database.AccessRequest{
			OrganizationName: "test-organization-a",
			ServiceName:      "test-service-1",
			State:            database.AccessRequestCreated,
		},
	}

	actual, err := cluster.DB.CreateOutgoingAccessRequest(ctx, a)
	assert.NoError(t, err)

	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 43, 1337, time.UTC))

	err = cluster.DB.UpdateOutgoingAccessRequestState(ctx, a, database.AccessRequestFailed)
	assert.NoError(t, err)

	expected := &database.OutgoingAccessRequest{
		AccessRequest: database.AccessRequest{
			ID:               "161c188cfcea1939",
			OrganizationName: "test-organization-a",
			ServiceName:      "test-service-1",
			State:            database.AccessRequestFailed,
			CreatedAt:        time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
			UpdatedAt:        time.Date(2020, time.June, 26, 12, 42, 43, 1337, time.UTC),
		},
	}

	assert.Equal(t, expected, actual)

	response, err := client.Get(ctx, "/nlx/access-requests/outgoing/test-organization-a/test-service-1/161c188cfcea1939")
	assert.NoError(t, err)
	assert.Len(t, response.Kvs, 1)
}

func TestLockOutgoingAccessRequest(t *testing.T) {
	cluster := newTestCluster(t)
	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()

	a := &database.OutgoingAccessRequest{
		AccessRequest: database.AccessRequest{
			OrganizationName: "test-organization-a",
			ServiceName:      "test-service-1",
		},
	}

	_, err := cluster.DB.CreateOutgoingAccessRequest(ctx, a)
	assert.NoError(t, err)

	err = cluster.DB.LockOutgoingAccessRequest(ctx, a)
	assert.NoError(t, err)

	err = cluster.DB.LockOutgoingAccessRequest(ctx, a)
	assert.Error(t, err)
	assert.Equal(t, database.ErrAccessRequestLocked, err)
}

func TestUnlockOutgoingAccessRequest(t *testing.T) {
	cluster := newTestCluster(t)
	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()

	a := &database.OutgoingAccessRequest{
		AccessRequest: database.AccessRequest{
			OrganizationName: "test-organization-a",
			ServiceName:      "test-service-1",
		},
	}

	_, err := cluster.DB.CreateOutgoingAccessRequest(ctx, a)
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
	cluster.Clock.SetTime(time.Date(2020, time.July, 9, 14, 44, 55, 0, time.UTC))

	ctx := context.Background()

	create := func(o, s string) {
		cluster.Clock.Step(5 * time.Second)

		_, err := cluster.DB.CreateOutgoingAccessRequest(ctx, &database.OutgoingAccessRequest{
			AccessRequest: database.AccessRequest{
				OrganizationName: o,
				ServiceName:      s,
			},
		})

		assert.NoError(t, err)
	}

	create("test-organization-a", "test-service-1") // 14:45:00

	expected := &database.OutgoingAccessRequest{
		AccessRequest: database.AccessRequest{
			ID:               "16201cc4e0cf3800",
			OrganizationName: "test-organization-a",
			ServiceName:      "test-service-1",
			State:            database.AccessRequestCreated,
			CreatedAt:        time.Date(2020, time.July, 9, 14, 45, 0, 0, time.UTC),
			UpdatedAt:        time.Date(2020, time.July, 9, 14, 45, 0, 0, time.UTC),
		},
	}

	actual, err := cluster.DB.GetLatestOutgoingAccessRequest(ctx, "test-organization-a", "test-service-1")

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestListAllLatestOutgoingAccessRequests(t *testing.T) {
	cluster := newTestCluster(t)
	cluster.Clock.SetTime(time.Date(2020, time.July, 10, 10, 11, 40, 0, time.UTC))

	ctx := context.Background()

	create := func(o, s string) {
		cluster.Clock.Step(5 * time.Second)

		_, err := cluster.DB.CreateOutgoingAccessRequest(ctx, &database.OutgoingAccessRequest{
			AccessRequest: database.AccessRequest{
				OrganizationName: o,
				ServiceName:      s,
			},
		})

		assert.NoError(t, err)
	}

	create("test-organization-b", "test-service-1") // 10:11:45
	create("test-organization-c", "test-service-1") // 10:11:50
	create("test-organization-a", "test-service-2") // 10:12:55

	expected := map[string]*database.OutgoingAccessRequest{
		"test-organization-a/test-service-2": {
			AccessRequest: database.AccessRequest{
				ID:               "16205c7284036e00",
				OrganizationName: "test-organization-a",
				ServiceName:      "test-service-2",
				State:            database.AccessRequestCreated,
				CreatedAt:        time.Date(2020, time.July, 10, 10, 11, 55, 0, time.UTC),
				UpdatedAt:        time.Date(2020, time.July, 10, 10, 11, 55, 0, time.UTC),
			},
		},
		"test-organization-b/test-service-1": {
			AccessRequest: database.AccessRequest{
				ID:               "16205c702ff78a00",
				OrganizationName: "test-organization-b",
				ServiceName:      "test-service-1",
				State:            database.AccessRequestCreated,
				CreatedAt:        time.Date(2020, time.July, 10, 10, 11, 45, 0, time.UTC),
				UpdatedAt:        time.Date(2020, time.July, 10, 10, 11, 45, 0, time.UTC),
			},
		},
		"test-organization-c/test-service-1": {
			AccessRequest: database.AccessRequest{
				ID:               "16205c7159fd7c00",
				OrganizationName: "test-organization-c",
				ServiceName:      "test-service-1",
				State:            database.AccessRequestCreated,
				CreatedAt:        time.Date(2020, time.July, 10, 10, 11, 50, 0, time.UTC),
				UpdatedAt:        time.Date(2020, time.July, 10, 10, 11, 50, 0, time.UTC),
			},
		},
	}

	actual, err := cluster.DB.ListAllLatestOutgoingAccessRequests(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetIncomingAccessRequest(t *testing.T) {
	cluster := newTestCluster(t)
	ctx := context.Background()
	client := cluster.GetClient(t)

	// Test with no incoming access requests
	actual, err := cluster.DB.GetIncomingAccessRequest(ctx, "1")
	assert.Nil(t, actual)
	assert.Nil(t, err)

	createAccessRequest := func(id, organization, service string) {
		bytes, _ := json.Marshal(
			&database.IncomingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               id,
					OrganizationName: organization,
					ServiceName:      service,
				},
			},
		)

		_, err := client.Put(ctx, path.Join("/nlx/access-requests/incoming", organization, service, id), string(bytes))
		assert.NoError(t, err)
	}

	createAccessRequest("1", "test-organization-a", "test-service-1")
	createAccessRequest("2", "test-organization-a", "test-service-1")
	createAccessRequest("3", "test-organization-a", "test-service-2")
	createAccessRequest("4", "test-organization-b", "test-service-1")

	tests := []struct {
		name     string
		id       string
		expected *database.IncomingAccessRequest
		err      error
	}{
		{
			"existing_access_request",
			"1",
			&database.IncomingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "1",
					OrganizationName: "test-organization-a",
					ServiceName:      "test-service-1",
				},
			},
			nil,
		},
		{
			"non_existing_access_request",
			"5",
			nil,
			nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			actual, err := cluster.DB.GetIncomingAccessRequest(ctx, test.id)

			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.err, err)
		})
	}
}
