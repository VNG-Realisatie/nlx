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

func TestAccessGrantRequest(t *testing.T) {
	cluster := newTestCluster(t)
	ctx := context.Background()
	client := cluster.GetClient(t)

	createAccessGrant := func(id, organization, service, fingerprint string) {
		a, _ := json.Marshal(
			&database.AccessGrant{
				OrganizationName:     organization,
				ServiceName:          service,
				PublicKeyFingerprint: fingerprint,
			},
		)

		_, err := client.Put(ctx, path.Join("/nlx/access-grants", service, organization, id), string(a))
		assert.NoError(t, err)
	}

	createAccessGrant("1", "test-organization-a", "test-service-1", "test-fingerprint-a")
	createAccessGrant("2", "test-organization-b", "test-service-1", "test-fingerprint-b")
	createAccessGrant("3", "test-organization-a", "test-service-2", "test-fingerprint-a")
	createAccessGrant("4", "test-organization-c", "test-service-1", "test-fingerprint-c")

	actual, err := cluster.DB.ListAccessGrantsForService(ctx, "test-service-1")

	assert.NoError(t, err)
	assert.Len(t, actual, 3)

	expected := []*database.AccessGrant{
		{
			OrganizationName:     "test-organization-a",
			ServiceName:          "test-service-1",
			PublicKeyFingerprint: "test-fingerprint-a",
		},
		{
			OrganizationName:     "test-organization-b",
			ServiceName:          "test-service-1",
			PublicKeyFingerprint: "test-fingerprint-b",
		},
		{
			OrganizationName:     "test-organization-c",
			ServiceName:          "test-service-1",
			PublicKeyFingerprint: "test-fingerprint-c",
		},
	}

	assert.Equal(t, expected, actual)
}

func TestCreateAccessGrant(t *testing.T) {
	cluster := newTestCluster(t)
	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()
	client := cluster.GetClient(t)

	accessRequest := &database.IncomingAccessRequest{
		AccessRequest: database.AccessRequest{
			ID:                   "access-request-id-1",
			OrganizationName:     "test-organization-a",
			ServiceName:          "test-service-1",
			PublicKeyFingerprint: "test-public-key-fingerprint",
		},
	}

	data, _ := json.Marshal(accessRequest)

	_, err := client.Put(ctx, path.Join("/nlx/access-requests/incoming", accessRequest.OrganizationName, accessRequest.ServiceName, "access-request-id-1"), string(data))
	assert.NoError(t, err)

	actual, err := cluster.DB.CreateAccessGrant(ctx, accessRequest)
	assert.NoError(t, err)

	expected := &database.AccessGrant{
		ID:                   "161c188cfcea1939",
		AccessRequestID:      "access-request-id-1",
		OrganizationName:     "test-organization-a",
		ServiceName:          "test-service-1",
		PublicKeyFingerprint: "test-public-key-fingerprint",
		CreatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
	}

	assert.Equal(t, expected, actual)
}

func TestCreateAccessGrantModified(t *testing.T) {
	cluster := newTestCluster(t)
	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()
	client := cluster.GetClient(t)

	accessRequest := &database.IncomingAccessRequest{
		AccessRequest: database.AccessRequest{
			ID:               "access-request-id-1",
			OrganizationName: "test-organization-a",
			ServiceName:      "test-service-1",
			State:            database.AccessRequestReceived,
		},
	}

	data, _ := json.Marshal(accessRequest)

	_, err := client.Put(ctx, path.Join("/nlx/access-requests/incoming", accessRequest.OrganizationName, accessRequest.ServiceName, "access-request-id-1"), string(data))
	assert.NoError(t, err)

	// Simulate that the access request is changed during the creation of the access grant
	// Possibility is that the access request is already approved.
	accessRequest.State = database.AccessRequestApproved
	accessRequest.UpdatedAt = time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)

	actual, err := cluster.DB.CreateAccessGrant(ctx, accessRequest)

	assert.Equal(t, err, database.ErrAccessRequestModified)
	assert.Nil(t, actual)
}
