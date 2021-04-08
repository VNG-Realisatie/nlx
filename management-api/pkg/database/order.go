// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"time"
)

type Order struct {
	ID           uint `gorm:"primarykey;column:order_id;"`
	Reference    string
	Description  string
	PublicKeyPEM string
	Delegatee    string
	ValidFrom    time.Time
	ValidUntil   time.Time
	Services     []string
	CreatedAt    time.Time
}

func (s *Order) TableName() string {
	return "nlx_management.orders"
}

func (db *PostgresConfigDatabase) CreateOrder(ctx context.Context, order *Order) error {
	return db.DB.
		WithContext(ctx).
		Create(order).Error
}
