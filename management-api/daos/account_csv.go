// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package daos

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	uuid "github.com/satori/go.uuid"

	"go.nlx.io/nlx/management-api/models"
)

// AccountCSV implements the queries of the Account dao on a csv file
type AccountCSV struct {
	accounts []*models.Account
}

// AccountCSVOptions contains go-flags fields which can be used to configure this dao
type AccountCSVOptions struct {
	CsvFileName string `long:"csv-filename" env:"CSV_FILENAME" required:"true" description:"Name of the file that contains the user accounts"`
}

// NewAccountCSV sets up a new CSV DAO for the session resource
func NewAccountCSV(fileName string) (*AccountCSV, error) {
	// Open the file
	csvfile, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}

	defer csvfile.Close()

	accounts := []*models.Account{}
	if err := gocsv.UnmarshalFile(csvfile, &accounts); err != nil {
		return nil, err
	}

	return &AccountCSV{accounts: accounts}, nil
}

// GetByID retrieves an account by its ID
func (dao AccountCSV) GetByID(id uuid.UUID) (*models.Account, error) {
	for _, account := range dao.accounts {
		if id == account.ID {
			return account, nil
		}
	}

	return nil, errors.New("account not found")
}

// GetByName retrieves an account by its Name
func (dao AccountCSV) GetByName(name string) (*models.Account, error) {
	for _, account := range dao.accounts {
		if name == account.Name {
			return account, nil
		}
	}

	return nil, errors.New("account not found")
}
