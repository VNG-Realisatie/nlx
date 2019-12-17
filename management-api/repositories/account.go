// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package repositories

import (
	uuid "github.com/satori/go.uuid"

	"go.nlx.io/nlx/management-api/daos"
	"go.nlx.io/nlx/management-api/models"
)

// Account provides the interface for an AccountImpl
type Account interface {
	GetByID(id uuid.UUID) (*models.Account, error)
	GetByName(name string) (*models.Account, error)
}

// AccountImpl persists accounts
type AccountImpl struct {
	accountDao daos.Account
}

// NewAccount creates a new Account
func NewAccount(accountDao daos.Account) (*AccountImpl, error) {
	return &AccountImpl{
		accountDao: accountDao,
	}, nil
}

// GetByID finds an AccountImpl in the store and returns it
func (s AccountImpl) GetByID(id uuid.UUID) (*models.Account, error) {
	account, err := s.accountDao.GetByID(id)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// GetByName finds an AccountImpl in the store and returns it
func (s AccountImpl) GetByName(name string) (*models.Account, error) {
	account, err := s.accountDao.GetByName(name)
	if err != nil {
		return nil, err
	}

	return account, nil
}
