// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package database

import (
	"context"

	"go.nlx.io/nlx/management-api/domain"
)

type Permission struct {
	Code string `gorm:"primaryKey"`
}

func (s *Permission) TableName() string {
	return "nlx_management.permissions"
}

func (db *PostgresConfigDatabase) ListPermissions(ctx context.Context) ([]*domain.Permission, error) {
	permissions, err := db.queries.ListPermissions(ctx)
	if err != nil {
		return nil, err
	}

	var result = make([]*domain.Permission, len(permissions))

	for i, permission := range permissions {
		result[i] = &domain.Permission{
			Code: permission,
		}
	}

	return result, nil
}
