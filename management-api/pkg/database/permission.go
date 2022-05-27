// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package database

import "context"

type Permission struct {
	Code string `gorm:"primaryKey"`
}

func (s *Permission) TableName() string {
	return "nlx_management.permissions"
}

func (db *PostgresConfigDatabase) ListPermissions(ctx context.Context) ([]Permission, error) {
	permissions := &[]Permission{}

	res := db.DB.WithContext(ctx).Find(permissions)
	if res.Error != nil {
		return nil, res.Error
	}

	return *permissions, nil
}
