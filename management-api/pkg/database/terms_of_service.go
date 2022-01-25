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

func (db *PostgresConfigDatabase) AcceptTermsOfService(ctx context.Context, username string, createdAt time.Time) error {
	if createdAt.After(time.Now()) {
		return ErrInvalidDate
	}

	return db.DB.
		WithContext(ctx).
		Where("id IS NOT NULL").
		FirstOrCreate(TermsOfServiceStatus{
			Username:  username,
			CreatedAt: createdAt,
		}).Error
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
