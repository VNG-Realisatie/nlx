// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package grpc_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	common_testing "go.nlx.io/nlx/testing/testingutils"
	txlog_mock "go.nlx.io/nlx/txlog-api/domain/record/mock"
	"go.nlx.io/nlx/txlog-api/ports/grpc"
	"go.nlx.io/nlx/txlog-api/service"
)

var fixedTestClockTime = time.Now()

type testClock struct {
	timeToReturn time.Time
}

func (c *testClock) Now() time.Time {
	return c.timeToReturn
}

func newStorageRepository(t *testing.T) (s *grpc.Server, m *txlog_mock.MockRepository) {
	logger := zaptest.Logger(t)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	m = txlog_mock.NewMockRepository(ctrl)

	clock := &testClock{
		timeToReturn: fixedTestClockTime,
	}

	app, err := service.NewApplication(&service.NewApplicationArgs{
		Context:    context.Background(),
		Logger:     logger,
		Repository: m,
	})
	assert.NoError(t, err)

	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	internalCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.NLXTestInternal)
	require.NoError(t, err)

	s = grpc.New(logger, clock, app, internalCert)

	return
}
