// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           uint `gorm:"primarykey;column:order_id;"`
	Reference    string
	Description  string
	PublicKeyPEM string
	Delegatee    string
	ValidFrom    time.Time
	ValidUntil   time.Time
	Services     []OrderService `gorm:"foreignKey:order_id;"`
	CreatedAt    time.Time
}

type OrderService struct {
	OrderID     uint
	ServiceName string
}

func (s *Order) TableName() string {
	return "nlx_management.orders"
}

func (db *PostgresConfigDatabase) CreateOrder(ctx context.Context, order *Order) error {
	if err := db.DB.
		WithContext(ctx).
		Create(order).Error; err != nil {
		return err
	}

	return nil
}

func (db *PostgresConfigDatabase) GetOrderByReference(ctx context.Context, reference string) (*Order, error) {
	order := &Order{}

	if err := db.DB.
		WithContext(ctx).
		Preload("Services").
		Where("reference = ?", reference).
		First(order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return order, nil
}
