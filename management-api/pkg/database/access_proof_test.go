// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestCreateAccessProof(t *testing.T) {
	cluster := newTestCluster(t)
	cluster.Clock.SetTime(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC))

	ctx := context.Background()

	AccessProof := &database.AccessProof{
		ID:               "access-request-id-1",
		OrganizationName: "test-organization-a",
		ServiceName:      "test-service-1",
	}

	actual, err := cluster.DB.CreateAccessProof(ctx, AccessProof)
	assert.NoError(t, err)

	expected := &database.AccessProof{
		ID:               "161c188cfcea1939",
		OrganizationName: "test-organization-a",
		ServiceName:      "test-service-1",
		CreatedAt:        time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
	}

	assert.Equal(t, expected, actual)
}

func TestRevokeAccessProof(t *testing.T) {
	clusterTime := time.Date(2020, time.October, 12, 13, 24, 16, 1337, time.UTC)
	cluster := newTestCluster(t)
	cluster.Clock.SetTime(clusterTime)

	ctx := context.Background()

	createdAccessProof, err := cluster.DB.CreateAccessProof(ctx, &database.AccessProof{
		ID:               "163d417ef83a0539",
		OrganizationName: "test-organization",
		ServiceName:      "test-service",
	})

	assert.NoError(t, err)

	tests := []struct {
		name             string
		serviceName      string
		organizationName string
		accessProofID    string
		revokeDate       time.Time
		accessProof      *database.AccessProof
		err              error
	}{
		{
			"unknown_access_proof",
			"test-service",
			"test-organization",
			"unknown-id",
			time.Time{},
			nil,
			database.ErrNotFound,
		},
		{
			"service_argument_mismatch",
			"other-service",
			"test-organization",
			"163d417ef83a0539",
			time.Time{},
			nil,
			database.ErrNotFound,
		},
		{
			"organization_argument_mismatch",
			"test-service",
			"other-organization",
			"163d417ef83a0539",
			time.Time{},
			nil,
			database.ErrNotFound,
		},
		{
			"happy_flow",
			"test-service",
			"test-organization",
			"163d417ef83a0539",
			time.Date(2020, time.October, 12, 13, 24, 16, 1337, time.UTC),
			&database.AccessProof{
				ID:               "163d417ef83a0539",
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				CreatedAt:        createdAccessProof.CreatedAt,
				RevokedAt:        time.Date(2020, time.October, 12, 13, 24, 16, 1337, time.UTC),
			},
			nil,
		},
		{
			"access_proof_revoked",
			"test-service",
			"test-organization",
			"163d417ef83a0539",
			time.Date(2020, time.October, 12, 13, 24, 16, 1337, time.UTC),
			nil,
			errors.New("access proof is already revoked"),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()

			accessProof, err := cluster.DB.RevokeAccessProof(ctx, test.organizationName, test.serviceName, test.accessProofID, test.revokeDate)

			assert.Equal(t, test.accessProof, accessProof)

			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			}
		})
	}
}
