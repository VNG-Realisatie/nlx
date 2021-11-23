// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

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

type OutgoingOrderService struct {
	OutgoingOrderID uint
	Service         string
	Organization    OutgoingOrderServiceOrganization `gorm:"embedded;embeddedPrefix:organization_"`
}

func (s *OutgoingOrderService) TableName() string {
	return "nlx_management.outgoing_orders_services"
}

type OutgoingOrderServiceOrganization struct {
	Name         string
	SerialNumber string
}

var ErrDuplicateOutgoingOrder = errors.New("duplicate outgoing order")

func (db *PostgresConfigDatabase) CreateOutgoingOrder(ctx context.Context, order *OutgoingOrder) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	duplicateDelegateWithReferenceSQLError := "duplicate key value violates unique constraint \"idx_outgoing_orders_delegatee_reference\""

	if err := dbWithTx.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Create(order).Error; err != nil {
		if strings.Contains(err.Error(), duplicateDelegateWithReferenceSQLError) {
			return ErrDuplicateOutgoingOrder
		}

		return err
	}

	orderServices := []OutgoingOrderService{}

	for _, service := range order.Services {
		orderServices = append(orderServices, OutgoingOrderService{
			OutgoingOrderID: order.ID,
			Organization: OutgoingOrderServiceOrganization{
				SerialNumber: service.Organization.SerialNumber,
				Name:         service.Organization.Name,
			},
			Service: service.Service,
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

func (db *PostgresConfigDatabase) ListOutgoingOrdersByOrganization(ctx context.Context, organizationSerialNumber string) ([]*OutgoingOrder, error) {
	orders := []*OutgoingOrder{}

	if err := db.DB.
		WithContext(ctx).
		Where("delegatee = ?", organizationSerialNumber).
		Order("valid_until desc").
		Preload("Services").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (db *PostgresConfigDatabase) UpdateOutgoingOrder(ctx context.Context, order *OutgoingOrder) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	duplicateDelegateWithReferenceSQLError := "duplicate key value violates unique constraint \"idx_outgoing_orders_delegatee_reference\""

	// Update (whitelisted) fields in the outgoing order table for the modified order
	if err := dbWithTx.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Select([]string{"description", "public_key_pem", "valid_from", "valid_until"}).
		Save(order).Error; err != nil {
		if strings.Contains(err.Error(), duplicateDelegateWithReferenceSQLError) {
			return ErrDuplicateOutgoingOrder
		}

		return err
	}

	// Delete all services for the updated order (note that we are in a transaction)
	err := dbWithTx.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Where("outgoing_order_id = ?", order.ID).
		Delete(&OutgoingOrderService{}).
		Error
	if err != nil {
		return err
	}

	// Add all services to the order
	orderServices := []OutgoingOrderService{}

	for _, service := range order.Services {
		orderServices = append(orderServices, OutgoingOrderService{
			OutgoingOrderID: order.ID,
			Organization: OutgoingOrderServiceOrganization{
				SerialNumber: service.Organization.SerialNumber,
				Name:         service.Organization.Name,
			},
			Service: service.Service,
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
