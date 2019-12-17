package models

import (
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID           uuid.UUID `json:"id"           csv:"id"`
	Name         string    `json:"name"         csv:"name"`
	PasswordHash string    `json:"passwordHash" csv:"password_hash"`
	Role         string    `json:"role"         csv:"role"`
}
