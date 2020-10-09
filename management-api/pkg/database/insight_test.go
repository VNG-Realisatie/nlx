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
	if err != nil {
		t.Fatal("error putting insight configuration", err)
	}

	insightConfig, err := cluster.DB.GetInsightConfiguration(ctx)
	if err != nil {
		t.Fatal("error getting insight configuration", err)
	}

	assert.Equal(t, mockInsightConfiguration, insightConfig)
}
