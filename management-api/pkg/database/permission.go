package database

import "time"

type Permission struct {
	ID        uint
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Permission) TableName() string {
	return "nlx_management.permissions"
}
