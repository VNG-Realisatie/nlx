// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type IncomingOrderService struct {
	IncomingOrderID uint
	Service         string
	Organization    string
}

func (s *IncomingOrderService) TableName() string {
	return "nlx_management.incoming_orders_services"
}

type IncomingOrder struct {
	ID           uint
	Reference    string
	Description  string
	PublicKeyPEM string
	Delegator    string
	CreatedAt    time.Time
	ValidFrom    time.Time
	ValidUntil   time.Time
	Services     []IncomingOrderService `gorm:"foreignKey:incoming_order_id;"`
}

func (o *IncomingOrder) TableName() string {
	return "nlx_management.incoming_orders"
}

func (db *PostgresConfigDatabase) GetIncomingOrderByReference(ctx context.Context, reference string) (*IncomingOrder, error) {
	order := &IncomingOrder{}

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

func (db *PostgresConfigDatabase) ListIncomingOrders(ctx context.Context) ([]*IncomingOrder, error) {
	orders := []*IncomingOrder{}

	if err := db.DB.
		WithContext(ctx).
		Order("valid_until desc").
		Preload("Services").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
