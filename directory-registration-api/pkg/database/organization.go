// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"fmt"
)

type Organization struct {
	Name                string
	InsightIrmaEndpoint string
	InsightLogEndpoint  string
}

// SetInsightConfiguration updates the insight configuration options for an organization
func (db PostgreSQLDirectoryDatabase) SetInsightConfiguration(ctx context.Context, organizationName, insightAPIURL, irmaServerURL string) error {
	_, err := db.setInsightConfigurationStatement.Exec(organizationName, insightAPIURL, irmaServerURL)
	if err != nil {
		return fmt.Errorf("failed to execute setInsightConfigurationStatement: %v", err)
	}

	return nil
}
