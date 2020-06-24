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

func TestListServices(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockService := &database.Service{
		Name: "my-service",
	}

	anotherMockService := &database.Service{
		Name: "another-service",
	}

	err := cluster.DB.CreateService(ctx, mockService)
	if err != nil {
		t.Fatal("error creating service", err)
	}

	err = cluster.DB.CreateService(ctx, anotherMockService)
	if err != nil {
		t.Fatal("error creating service", err)
	}

	services, err := cluster.DB.ListServices(ctx)
	if err != nil {
		t.Fatal("error listing services", err)
	}

	assert.Equal(t, []*database.Service{anotherMockService, mockService}, services)
}

func TestCreateGetService(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockService := &database.Service{
		Name: "my-service",
	}

	err := cluster.DB.CreateService(ctx, mockService)
	if err != nil {
		t.Fatal("error creating service", err)
	}

	service, err := cluster.DB.GetService(ctx, "my-service")
	if err != nil {
		t.Fatal("error getting service", err)
	}

	assert.Equal(t, service, mockService)
}

func TestUpdateService(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockService := &database.Service{
		Name:        "my-service",
		EndpointURL: "https://somewhere/",
	}

	updatedMockService := &database.Service{
		Name:        "my-service",
		EndpointURL: "https://somewhere-else/",
	}

	err := cluster.DB.CreateService(ctx, mockService)
	if err != nil {
		t.Fatal("error creating service", err)
	}

	service, err := cluster.DB.GetService(ctx, "my-service")
	if err != nil {
		t.Fatal("error getting service", err)
	}

	assert.Equal(t, service.EndpointURL, mockService.EndpointURL)

	err = cluster.DB.UpdateService(ctx, "my-service", updatedMockService)
	if err != nil {
		t.Fatal("error updating service", err)
	}

	service, err = cluster.DB.GetService(ctx, "my-service")
	if err != nil {
		t.Fatal("error getting service", err)
	}

	assert.Equal(t, service.EndpointURL, updatedMockService.EndpointURL)
}

func TestDeleteService(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockService := &database.Service{
		Name:        "my-service",
		EndpointURL: "https://somewhere/",
	}

	err := cluster.DB.CreateService(ctx, mockService)
	if err != nil {
		t.Fatal("error creating service", err)
	}

	service, err := cluster.DB.GetService(ctx, "my-service")
	if err != nil {
		t.Fatal("error getting service", err)
	}

	assert.Equal(t, service, mockService)

	err = cluster.DB.DeleteService(ctx, "my-service")
	if err != nil {
		t.Fatal("error deleting service", err)
	}

	service, err = cluster.DB.GetService(ctx, "my-service")
	if err != nil {
		t.Fatal("error getting service", err)
	}

	assert.Nil(t, service)
}

func TestFilterServices(t *testing.T) {
	type args struct {
		services []*database.Service
		inway    *database.Inway
	}

	var filterServicesTests = []struct {
		name string
		want []*database.Service
		args args
	}{
		{
			name: "one service",
			args: args{
				services: []*database.Service{{
					Name:   "service1",
					Inways: []string{"inway1"},
				}, {
					Name:   "service2",
					Inways: []string{"inway2"},
				}},
				inway: &database.Inway{
					Name: "inway1",
				}},
			want: []*database.Service{{
				Name:   "service1",
				Inways: []string{"inway1"},
			}},
		},
		{
			name: "two services",
			args: args{
				services: []*database.Service{{
					Name:   "service11",
					Inways: []string{"inway1"},
				}, {
					Name:   "service12",
					Inways: []string{"inway1"},
				}, {
					Name:   "service2",
					Inways: []string{"inway2"},
				}},
				inway: &database.Inway{
					Name: "inway1",
				}},
			want: []*database.Service{{
				Name:   "service11",
				Inways: []string{"inway1"},
			}, {
				Name:   "service12",
				Inways: []string{"inway1"},
			}},
		},
		{
			name: "no services",
			args: args{
				services: []*database.Service{{
					Name:   "service1",
					Inways: []string{"inway1"},
				}},
				inway: &database.Inway{
					Name: "inway2",
				}},
			want: []*database.Service{},
		},
	}

	for _, tt := range filterServicesTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := database.FilterServices(tt.args.services, tt.args.inway)
			assert.Equal(t, tt.want, actual)
		})
	}
}
