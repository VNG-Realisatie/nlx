package daos

import (
	uuid "github.com/satori/go.uuid"

	"go.nlx.io/nlx/management-api/models"
)

// Account interface
type Account interface {
	GetByID(id uuid.UUID) (*models.Account, error)
	GetByName(name string) (*models.Account, error)
}
