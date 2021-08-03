// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OutgoingOrderService struct {
	OutgoingOrderID uint
	Service         string
	Organization    string
}

func (s *OutgoingOrderService) TableName() string {
	return "nlx_management.outgoing_orders_services"
}

type OutgoingOrder struct {
	ID           uint
	Reference    string
	Description  string
	PublicKeyPEM string
	Delegatee    string
	RevokedAt    sql.NullTime
	CreatedAt    time.Time
	ValidFrom    time.Time
	ValidUntil   time.Time
	Services     []OutgoingOrderService `gorm:"foreignKey:outgoing_order_id;"`
}

func (o *OutgoingOrder) TableName() string {
	return "nlx_management.outgoing_orders"
}

func (db *PostgresConfigDatabase) CreateOutgoingOrder(ctx context.Context, order *OutgoingOrder) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	if err := dbWithTx.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Create(order).Error; err != nil {
		return err
	}

	orderServices := []OutgoingOrderService{}

	for _, service := range order.Services {
		orderServices = append(orderServices, OutgoingOrderService{
			OutgoingOrderID: order.ID,
			Organization:    service.Organization,
			Service:         service.Service,
		})
	}

	if err := dbWithTx.DB.
		WithContext(ctx).
		Model(OutgoingOrderService{}).
		Create(orderServices).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func (db *PostgresConfigDatabase) GetOutgoingOrderByReference(ctx context.Context, reference string) (*OutgoingOrder, error) {
	order := &OutgoingOrder{}

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

func (db *PostgresConfigDatabase) ListOutgoingOrders(ctx context.Context) ([]*OutgoingOrder, error) {
	orders := []*OutgoingOrder{}

	if err := db.DB.
		WithContext(ctx).
		Order("valid_until desc").
		Preload("Services").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (db *PostgresConfigDatabase) ListOutgoingOrdersByOrganization(ctx context.Context, organizationName string) ([]*OutgoingOrder, error) {
	orders := []*OutgoingOrder{}

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

// nolint dupl: function is not duplicated, difference between incoming and outgoing orders
func (db *PostgresConfigDatabase) RevokeOutgoingOrderByReference(ctx context.Context, delegatee, reference string, revokedAt time.Time) error {
	outgoingOrder := &OutgoingOrder{}

	if err := db.DB.
		WithContext(ctx).
		Where("delegatee = ? AND reference = ?", delegatee, reference).
		First(outgoingOrder).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}

		return err
	}

	if outgoingOrder.RevokedAt.Valid {
		return nil
	}

	outgoingOrder.RevokedAt = sql.NullTime{
		Time:  revokedAt,
		Valid: true,
	}

	return db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Select("revoked_at").
		Save(outgoingOrder).Error
}
