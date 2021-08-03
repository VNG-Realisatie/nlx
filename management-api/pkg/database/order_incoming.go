// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	RevokedAt    sql.NullTime
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

func (db *PostgresConfigDatabase) SynchronizeOrders(ctx context.Context, orders []*IncomingOrder) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	for _, order := range orders {
		existingOrder, err := db.GetIncomingOrderByReference(ctx, order.Reference)
		if errors.Is(err, ErrNotFound) {
			if createOrderErr := dbWithTx.DB.
				WithContext(ctx).
				Omit(clause.Associations).
				Create(order).
				Error; createOrderErr != nil {
				return createOrderErr
			}

			orderServices := []IncomingOrderService{}

			for _, service := range order.Services {
				orderServices = append(orderServices, IncomingOrderService{
					IncomingOrderID: order.ID,
					Organization:    service.Organization,
					Service:         service.Service,
				})
			}

			if createServicesErr := dbWithTx.DB.
				WithContext(ctx).
				Model(IncomingOrderService{}).
				Create(orderServices).Error; createServicesErr != nil {
				return createServicesErr
			}

			continue
		} else if err != nil {
			return fmt.Errorf("failed to get order by reference: %w", err)
		}

		order.ID = existingOrder.ID

		// if nothing changed skip it
		if isOrderEqual(existingOrder, order) {
			continue
		}

		existingOrder.ValidUntil = order.ValidUntil
		existingOrder.Description = order.Description
		existingOrder.RevokedAt = order.RevokedAt

		if err := dbWithTx.DB.
			WithContext(ctx).
			Omit(clause.Associations).
			Save(existingOrder).
			Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error
}

// nolint dupl: function is not duplicated, difference between incoming and outgoing orders
func (db *PostgresConfigDatabase) RevokeIncomingOrderByReference(ctx context.Context, delegator, reference string, revokedAt time.Time) error {
	incomingOrder := &IncomingOrder{}

	if err := db.DB.
		WithContext(ctx).
		Where("delegator = ? AND reference = ?", delegator, reference).
		First(incomingOrder).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}

		return err
	}

	if incomingOrder.RevokedAt.Valid {
		return nil
	}

	incomingOrder.RevokedAt = sql.NullTime{
		Time:  revokedAt,
		Valid: true,
	}

	return db.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Select("revoked_at").
		Save(incomingOrder).Error
}

func isOrderEqual(a, b *IncomingOrder) bool {
	return a.ValidUntil.Equal(b.ValidUntil) &&
		a.Description == b.Description &&
		a.RevokedAt == b.RevokedAt
}
