// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
)

type InsightConfiguration struct {
	IrmaServerURL string `json:"irmaServerURL,omitempty"`
	InsightAPIURL string `json:"insightAPIURL,omitempty"`
}

const insightConfigurationKey = "insight-configuration"

// PutInsight sets the insight configuration
func (db ETCDConfigDatabase) PutInsightConfiguration(ctx context.Context, insightConfiguration *InsightConfiguration) error {
	if err := db.put(ctx, insightConfigurationKey, &insightConfiguration); err != nil {
		return err
	}

	return nil
}

// GetInsight returns the insight configuration
func (db ETCDConfigDatabase) GetInsightConfiguration(ctx context.Context) (*InsightConfiguration, error) {
	insightConfiguration := &InsightConfiguration{}

	if err := db.get(ctx, insightConfigurationKey, insightConfiguration); err != nil {
		return nil, err
	}

	return insightConfiguration, nil
}
