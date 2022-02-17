// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreateOutgoingOrder struct {
	ID             uint64
	Reference      string
	Description    string
	PublicKeyPEM   string
	Delegatee      string
	CreatedAt      time.Time
	ValidFrom      time.Time
	ValidUntil     time.Time
	AccessProofIds []uint64 `gorm:"-"`
}

type UpdateOutgoingOrder struct {
	ID             uint64
	Reference      string
	Description    string
	PublicKeyPEM   string
	ValidFrom      time.Time
	ValidUntil     time.Time
	AccessProofIds []uint64 `gorm:"-"`
}

type OutgoingOrder struct {
	ID                        uint64
	Reference                 string
	Description               string
	PublicKeyPEM              string
	Delegatee                 string
	RevokedAt                 sql.NullTime
	CreatedAt                 time.Time
	ValidFrom                 time.Time
	ValidUntil                time.Time
	OutgoingOrderAccessProofs []*OutgoingOrderAccessProof `gorm:"foreignKey:outgoing_order_id"`
}

type OutgoingOrderAccessProof struct {
	OutgoingOrderID uint64
	AccessProofID   uint64
	AccessProof     *AccessProof `gorm:"foreignKey:access_proof_id;"`
}

func (o *OutgoingOrder) TableName() string {
	return "nlx_management.outgoing_orders"
}

func (o *CreateOutgoingOrder) TableName() string {
	return "nlx_management.outgoing_orders"
}

func (o *UpdateOutgoingOrder) TableName() string {
	return "nlx_management.outgoing_orders"
}

func (o *OutgoingOrderAccessProof) TableName() string {
	return "nlx_management.outgoing_orders_access_proofs"
}

type OutgoingOrderServiceOrganization struct {
	Name         string
	SerialNumber string
}

var ErrDuplicateOutgoingOrder = errors.New("duplicate outgoing order")

func (db *PostgresConfigDatabase) CreateOutgoingOrder(ctx context.Context, order *CreateOutgoingOrder) error {
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

	orderAccessProofs := make([]*OutgoingOrderAccessProof, len(order.AccessProofIds))
	for i, id := range order.AccessProofIds {
		orderAccessProofs[i] = &OutgoingOrderAccessProof{
			OutgoingOrderID: order.ID,
			AccessProofID:   id,
		}
	}

	if err := dbWithTx.DB.
		WithContext(ctx).
		Model(OutgoingOrderAccessProof{}).
		Create(orderAccessProofs).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func (db *PostgresConfigDatabase) GetOutgoingOrderByReference(ctx context.Context, reference string) (*OutgoingOrder, error) {
	order := &OutgoingOrder{}

	if err := db.DB.
		WithContext(ctx).
		Preload("OutgoingOrderAccessProofs").
		Preload("OutgoingOrderAccessProofs.AccessProof").
		Preload("OutgoingOrderAccessProofs.AccessProof.OutgoingAccessRequest").
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
		Preload("OutgoingOrderAccessProofs").
		Preload("OutgoingOrderAccessProofs.AccessProof").
		Preload("OutgoingOrderAccessProofs.AccessProof.OutgoingAccessRequest").
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
		Preload("OutgoingOrderAccessProofs").
		Preload("OutgoingOrderAccessProofs.AccessProof").
		Preload("OutgoingOrderAccessProofs.AccessProof.OutgoingAccessRequest").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (db *PostgresConfigDatabase) UpdateOutgoingOrder(ctx context.Context, order *UpdateOutgoingOrder) error {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	duplicateDelegateWithReferenceSQLError := "duplicate key value violates unique constraint \"idx_outgoing_orders_delegatee_reference\""

	// Update the outgoing order table for the modified order, uses order.ID as pk
	if err := dbWithTx.DB.
		WithContext(ctx).
		Model(&order).
		Where("reference=?", order.Reference).
		Updates(order).Error; err != nil {
		if strings.Contains(err.Error(), duplicateDelegateWithReferenceSQLError) {
			return ErrDuplicateOutgoingOrder
		}

		return err
	}

	// Delete all access proofs for the updated order (note that we are in a transaction)
	if err := dbWithTx.DB.
		WithContext(ctx).
		Where(fmt.Sprintf("outgoing_order_id = (select id FROM %s WHERE reference = ? LIMIT 1)", order.TableName()), order.Reference).
		Delete(&OutgoingOrderAccessProof{}).Error; err != nil {
		return err
	}

	// Add all access proofs to the order
	orderAccessProofs := make([]*OutgoingOrderAccessProof, len(order.AccessProofIds))
	for i, id := range order.AccessProofIds {
		orderAccessProofs[i] = &OutgoingOrderAccessProof{
			OutgoingOrderID: order.ID,
			AccessProofID:   id,
		}
	}

	if len(orderAccessProofs) > 0 {
		if err := dbWithTx.DB.
			WithContext(ctx).
			Model(OutgoingOrderAccessProof{}).
			Create(orderAccessProofs).Error; err != nil {
			return err
		}
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
