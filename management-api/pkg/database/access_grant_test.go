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
	defer cluster.Terminate(t)

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
	defer cluster.Terminate(t)

	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()
	client := cluster.GetClient(t)

	a := &database.AccessGrant{
		OrganizationName:     "test-organization-a",
		ServiceName:          "test-service-1",
		PublicKeyFingerprint: "test-fingerprint-a",
	}

	actual, err := cluster.DB.CreateAccessGrant(ctx, a)
	assert.NoError(t, err)

	expected := &database.AccessGrant{
		ID:                   "161c188cfcea1939",
		OrganizationName:     "test-organization-a",
		ServiceName:          "test-service-1",
		PublicKeyFingerprint: "test-fingerprint-a",

		CreatedAt: time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
	}

	assert.Equal(t, expected, actual)

	response, err := client.Get(ctx, "/nlx/access-grants/test-service-1/test-organization-a/161c188cfcea1939")
	assert.NoError(t, err)
	assert.Len(t, response.Kvs, 1)
}
