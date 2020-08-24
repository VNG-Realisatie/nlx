// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
)

type Settings struct {
	InwayNameForManagementAPITraffic string `json:"inwayNameForManagementAPITraffic"`
}

const settingsKey = "settings"

func (db ETCDConfigDatabase) GetSettings(ctx context.Context) (*Settings, error) {
	r := &Settings{}

	err := db.get(ctx, settingsKey, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db ETCDConfigDatabase) UpdateSettings(ctx context.Context, settings *Settings) error {
	err := db.put(ctx, settingsKey, &settings)
	if err != nil {
		return err
	}

	return nil
}
