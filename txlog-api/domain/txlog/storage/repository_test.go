// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package storage_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"

	pgadapter_test_setup "go.nlx.io/nlx/txlog-api/adapters/storage/postgres/test_setup"
	"go.nlx.io/nlx/txlog-api/domain"
	"go.nlx.io/nlx/txlog-api/domain/txlog/storage"
)

func new(t *testing.T, enableFixtures bool) (storage.Repository, func() error) {
	repo, close := pgadapter_test_setup.New(t)
	if enableFixtures {
		err := loadFixtures(t, repo)
		require.NoError(t, err)
	}

	return repo, close
}

func loadFixtures(t *testing.T, repo storage.Repository) error {
	newRecordsArgs := []*domain.NewRecordArgs{
		{
			Source:        createNewOrganization(t, "0001"),
			Destination:   createNewOrganization(t, "0002"),
			Direction:     domain.IN,
			Service:       createNewService(t, "test-service"),
			Order:         createNewOrder(t, "0003", "test-reference"),
			Data:          []byte(`{"test": "data"}`),
			CreatedAt:     time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			TransactionID: "abcde",
		},
	}

	for _, args := range newRecordsArgs {
		record, err := domain.NewRecord(args)
		if err != nil {
			return err
		}

		err = repo.CreateRecord(context.Background(), record)
		if err != nil {
			return err
		}
	}

	return nil
}

func assertRecordInRepository(t *testing.T, repo storage.Repository, r *domain.Record) {
	require.NotNil(t, r)

	records, err := repo.ListRecords(context.Background(), 100)
	require.NoError(t, err)

	var found bool
	for _, record := range records {
		if record.TransactionID() == r.TransactionID() {
			found = true
			break
		}
	}

	require.Equal(t, true, found)
}

func createNewOrganization(t *testing.T, serialNumber string) *domain.Organization {
	m, err := domain.NewOrganization(serialNumber)
	require.NoError(t, err)

	return m
}

func createNewService(t *testing.T, name string) *domain.Service {
	m, err := domain.NewService(name)
	require.NoError(t, err)

	return m
}

func createNewOrder(t *testing.T, delegator, reference string) *domain.Order {
	m, err := domain.NewOrder(&domain.NewOrderArgs{
		Delegator: delegator,
		Reference: reference,
	})
	require.NoError(t, err)

	return m
}
