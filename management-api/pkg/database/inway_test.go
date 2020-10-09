// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package database_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestListInways(t *testing.T) {
	cluster := newTestCluster(t)
	ctx := context.Background()

	mockInway := &database.Inway{
		Name: "inway42.test",
	}

	anotherMockInway := &database.Inway{
		Name: "inway43.test",
	}

	err := cluster.DB.CreateInway(ctx, mockInway)
	assert.NoError(t, err)

	err = cluster.DB.CreateInway(ctx, anotherMockInway)
	assert.NoError(t, err)

	inways, err := cluster.DB.ListInways(ctx)
	assert.NoError(t, err)

	assert.Equal(t, []*database.Inway{mockInway, anotherMockInway}, inways)
}

func TestCreateGetInway(t *testing.T) {
	cluster := newTestCluster(t)
	ctx := context.Background()

	mockInway := &database.Inway{
		Name: "my-inway",
	}

	err := cluster.DB.CreateInway(ctx, mockInway)
	assert.NoError(t, err)

	service, err := cluster.DB.GetInway(ctx, "my-inway")
	assert.NoError(t, err)
	assert.Equal(t, service, mockInway)
}

func TestUpdateInway(t *testing.T) {
	cluster := newTestCluster(t)
	ctx := context.Background()

	mockInway := &database.Inway{
		Name: "my-inway",
	}

	mockUpdatedInway := &database.Inway{
		Name: "my-inway",
	}

	err := cluster.DB.CreateInway(ctx, mockInway)
	assert.NoError(t, err)

	inway, err := cluster.DB.GetInway(ctx, "my-inway")
	assert.NoError(t, err)
	assert.Equal(t, mockInway, inway)

	err = cluster.DB.UpdateInway(ctx, "my-inway", mockUpdatedInway)
	assert.NoError(t, err)

	inway, err = cluster.DB.GetInway(ctx, "my-inway")
	assert.NoError(t, err)
	assert.Equal(t, inway, mockUpdatedInway)
}

func TestDeleteInway(t *testing.T) {
	cluster := newTestCluster(t)
	ctx := context.Background()

	mockInway := &database.Inway{
		Name: "my-inway",
	}

	err := cluster.DB.CreateInway(ctx, mockInway)
	assert.NoError(t, err)

	inway, err := cluster.DB.GetInway(ctx, "my-inway")
	assert.NoError(t, err)
	assert.Equal(t, inway, mockInway)

	err = cluster.DB.DeleteInway(ctx, "my-inway")
	assert.NoError(t, err)

	inway, err = cluster.DB.GetInway(ctx, "my-inway")
	assert.Error(t, err, database.ErrNotFound)
	assert.Nil(t, inway)
}
