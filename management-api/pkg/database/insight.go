// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"encoding/json"
	"path"
)

type InsightConfiguration struct {
	IrmaServerURL string `json:"irmaServerURL,omitempty"`
	InsightAPIURL string `json:"insightAPIURL,omitempty"`
}

// PutInsight sets the insight configuration
func (db ETCDConfigDatabase) PutInsightConfiguration(ctx context.Context, insightConfiguration *InsightConfiguration) error {
	key := path.Join(db.pathPrefix, "insight-configuration")

	data, err := json.Marshal(&insightConfiguration)
	if err != nil {
		return err
	}

	_, err = db.etcdCli.Put(ctx, key, string(data))
	if err != nil {
		return err
	}

	return nil
}

// GetInsight returns the insight configuration
func (db ETCDConfigDatabase) GetInsightConfiguration(ctx context.Context) (*InsightConfiguration, error) {
	key := path.Join(db.pathPrefix, "insight-configuration")

	values, err := db.etcdCli.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if values.Count == 0 {
		return nil, nil
	}

	insightConfiguration := &InsightConfiguration{}
	err = json.Unmarshal(values.Kvs[0].Value, insightConfiguration)

	if err != nil {
		return nil, err
	}

	return insightConfiguration, nil
}
