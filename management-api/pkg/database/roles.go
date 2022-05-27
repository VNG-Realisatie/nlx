// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package database

type Role struct {
	Code        string       `gorm:"primaryKey"`
	Permissions []Permission `gorm:"many2many:nlx_management.permissions_roles;"`
}

func (s *Role) TableName() string {
	return "nlx_management.roles"
}
