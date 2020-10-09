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
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	err = cluster.DB.CreateInway(ctx, anotherMockInway)
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	inways, err := cluster.DB.ListInways(ctx)
	if err != nil {
		t.Fatal("error listing inways", err)
	}

	assert.Equal(t, []*database.Inway{mockInway, anotherMockInway}, inways)
}

func TestCreateGetInway(t *testing.T) {
	cluster := newTestCluster(t)
	ctx := context.Background()

	mockInway := &database.Inway{
		Name: "my-inway",
	}

	err := cluster.DB.CreateInway(ctx, mockInway)
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	service, err := cluster.DB.GetInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error getting inway", err)
	}

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
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	inway, err := cluster.DB.GetInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error getting inway", err)
	}

	assert.Equal(t, mockInway, inway)

	err = cluster.DB.UpdateInway(ctx, "my-inway", mockUpdatedInway)
	if err != nil {
		t.Fatal("error updating inway", err)
	}

	inway, err = cluster.DB.GetInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error getting inway", err)
	}

	assert.Equal(t, inway, mockUpdatedInway)
}

func TestDeleteInway(t *testing.T) {
	cluster := newTestCluster(t)
	ctx := context.Background()

	mockInway := &database.Inway{
		Name: "my-inway",
	}

	err := cluster.DB.CreateInway(ctx, mockInway)
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	inway, err := cluster.DB.GetInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error getting inway", err)
	}

	assert.Equal(t, inway, mockInway)

	err = cluster.DB.DeleteInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error deleting inway", err)
	}

	inway, err = cluster.DB.GetInway(ctx, "my-inway")
	assert.Error(t, err, database.ErrNotFound)
	assert.Nil(t, inway)
}
