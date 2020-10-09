// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database_test

import (
	"context"
	"testing"

	"go.nlx.io/nlx/management-api/pkg/database"

	"github.com/stretchr/testify/assert"
)

func TestPutGetInsight(t *testing.T) {
	cluster := newTestCluster(t)
	ctx := context.Background()

	mockInsightConfiguration := &database.InsightConfiguration{
		IrmaServerURL: "http://irma-url.com",
		InsightAPIURL: "http://insight-url.com",
	}

	err := cluster.DB.PutInsightConfiguration(ctx, mockInsightConfiguration)
	assert.NoError(t, err)

	insightConfig, err := cluster.DB.GetInsightConfiguration(ctx)
	assert.NoError(t, err)
	assert.Equal(t, mockInsightConfiguration, insightConfig)
}
