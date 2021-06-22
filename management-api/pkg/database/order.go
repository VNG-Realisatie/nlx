// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order struct {
	ID           uint
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
	OrderID      uint
	Service      string
	Organization string
}

func (s *OrderService) TableName() string {
	return "nlx_management.orders_services"
}

func (s *Order) TableName() string {
	return "nlx_management.orders"
}

func (db *PostgresConfigDatabase) CreateOrder(ctx context.Context, order *Order) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	if err := dbWithTx.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Create(order).Error; err != nil {
		return err
	}

	orderServices := []OrderService{}

	for _, service := range order.Services {
		orderServices = append(orderServices, OrderService{
			OrderID:      order.ID,
			Organization: service.Organization,
			Service:      service.Service,
		})
	}

	if err := dbWithTx.DB.
		WithContext(ctx).
		Model(OrderService{}).
		Create(orderServices).Error; err != nil {
		return err
	}

	return tx.Commit().Error
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

func (db *PostgresConfigDatabase) ListIssuedOrders(ctx context.Context) ([]*Order, error) {
	orders := []*Order{}

	if err := db.DB.
		WithContext(ctx).
		Order("valid_until desc").
		Preload("Services").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (db *PostgresConfigDatabase) ListOrdersByOrganization(ctx context.Context, organizationName string) ([]*Order, error) {
	orders := []*Order{}

	if err := db.DB.
		WithContext(ctx).
		Where("delegatee = ?", organizationName).
		Order("valid_until desc").
		Preload("Services").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
