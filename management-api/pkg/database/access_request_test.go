// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

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
		Status:           database.AccessRequestCreated,
		CreatedAt:        time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
		UpdatedAt:        time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
	}

	assert.Equal(t, expected, actual)

	response, err := client.Get(ctx, "/nlx/access-requests/outgoing/test-organization-a/test-service-1/161c188cfcea1939")
	assert.NoError(t, err)
	assert.Len(t, response.Kvs, 1)
}
