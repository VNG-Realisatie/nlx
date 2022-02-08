// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
)

func (db *PostgresConfigDatabase) GetFingerprintOfPublicKeys(ctx context.Context) ([]string, error) {
	outway := &Outway{}
	fingerPrints := &[]string{}
	err := db.DB.
		WithContext(ctx).
		Model(outway).
		Select("public_key_fingerprint").
		Group("public_key_fingerprint").
		Find(fingerPrints).Error

	if err != nil {
		return []string{}, err
	}

	return *fingerPrints, nil
}
