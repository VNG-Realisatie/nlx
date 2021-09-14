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

	"go.nlx.io/nlx/management-api/domain"
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

func (db *PostgresConfigDatabase) ListIncomingOrders(ctx context.Context) ([]*domain.IncomingOrder, error) {
	orders := []*IncomingOrder{}
	if err := db.DB.
		WithContext(ctx).
		Order("valid_until desc").
		Preload("Services").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	incomingOrders := make([]*domain.IncomingOrder, len(orders))

	for i, order := range orders {
		services := make([]domain.IncomingOrderService, len(order.Services))

		for i, service := range order.Services {
			services[i] = domain.NewIncomingOrderService(service.Service, service.Organization)
		}

		var revokedAt *time.Time
		if order.RevokedAt.Valid {
			revokedAt = &order.RevokedAt.Time
		}

		incomingOrder, err := domain.NewIncomingOrder(order.Reference, order.Description, order.Delegator, revokedAt, order.ValidFrom, order.ValidUntil, services)
		if err != nil {
			return nil, fmt.Errorf("error converting incoming order: %w", err)
		}

		incomingOrders[i] = incomingOrder
	}

	return incomingOrders, nil
}

func (db *PostgresConfigDatabase) SynchronizeOrders(ctx context.Context, orders []*IncomingOrder) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	organizations := getUniqueOrganizations(orders)

	err := dbWithTx.DB.
		Debug().
		WithContext(ctx).
		Omit(clause.Associations).
		Delete(&IncomingOrder{}, "delegator IN ?", organizations).
		Error
	if err != nil {
		return err
	}

	for _, order := range orders {
		if createOrderErr := dbWithTx.DB.
			WithContext(ctx).
			Omit(clause.Associations).
			Create(order).
			Error; createOrderErr != nil {
			return createOrderErr
		}

		if len(order.Services) == 0 {
			continue
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

func getUniqueOrganizations(orders []*IncomingOrder) []string {
	uniqueOrgs := make(map[string]interface{})
	for _, order := range orders {
		uniqueOrgs[order.Delegator] = nil
	}

	var i int

	stringOrgs := make([]string, len(uniqueOrgs))
	for org := range uniqueOrgs {
		stringOrgs[i] = org
		i++
	}

	return stringOrgs
}
