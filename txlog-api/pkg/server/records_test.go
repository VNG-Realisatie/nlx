// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/txlog-api/domain"
	mock_txlog "go.nlx.io/nlx/txlog-api/domain/txlog/storage/mock"
	"go.nlx.io/nlx/txlog-api/pkg/server"
)

var fixedTestClockTime = time.Now()

type testClock struct {
	timeToReturn time.Time
}

func (c *testClock) Now() time.Time {
	return c.timeToReturn
}

func newStorageRepository(t *testing.T) (s *server.TXLogService, m *mock_txlog.MockRepository) {
	logger := zaptest.Logger(t)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	m = mock_txlog.NewMockRepository(ctrl)

	clock := &testClock{
		timeToReturn: fixedTestClockTime,
	}

	s = server.NewTXLogService(logger, m, clock)

	return
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
