// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"go.nlx.io/nlx/management-api/domain"
)

type TermsOfServiceStatus struct {
	Username  string
	CreatedAt time.Time
}

var ErrInvalidDate = errors.New("date cannot be in the future")

func (i *TermsOfServiceStatus) TableName() string {
	return "nlx_management.terms_of_service"
}

func (db *PostgresConfigDatabase) AcceptTermsOfService(ctx context.Context, username string, createdAt time.Time) (bool, error) {
	if createdAt.After(time.Now()) {
		return false, ErrInvalidDate
	}

	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	var count int64
	tos := dbWithTx.DB.
		Debug().
		WithContext(ctx).
		Model(&TermsOfServiceStatus{}).
		Count(&count)

	if tos.Error != nil {
		return false, tos.Error
	}

	if count > 0 {
		return true, nil
	}

	create := dbWithTx.DB.
		Debug().
		WithContext(ctx).
		Create(&TermsOfServiceStatus{
			Username:  username,
			CreatedAt: createdAt,
		})
	if create.Error != nil {
		return false, create.Error
	}

	dbWithTx.DB.Commit()

	return false, nil
}

func (db *PostgresConfigDatabase) GetTermsOfServiceStatus(ctx context.Context) (*domain.TermsOfServiceStatus, error) {
	tos := &TermsOfServiceStatus{}

	if err := db.DB.
		WithContext(ctx).
		First(tos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	model, err := domain.NewTermsOfServiceStatus(&domain.NewTermsOfServiceStatusArgs{
		Username:  tos.Username,
		CreatedAt: tos.CreatedAt,
	})
	if err != nil {
		return nil, err
	}

	return model, nil
}
