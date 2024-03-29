// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package queries

import (
	"database/sql"
	"time"
)

type DirectoryAvailability struct {
	ID             int32
	InwayID        int32
	ServiceID      int32
	Healthy        bool
	UnhealthySince sql.NullTime
	LastAnnounced  time.Time
	Active         bool
}

type DirectoryInway struct {
	ID                        int32
	OrganizationID            int32
	Address                   string
	Version                   string
	Name                      string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	ManagementApiProxyAddress sql.NullString
}

type DirectoryOrganization struct {
	ID           int32
	Name         string
	InwayID      sql.NullInt32
	SerialNumber string
	CreatedAt    time.Time
	EmailAddress sql.NullString
}

type DirectoryOutway struct {
	ID             int32
	OrganizationID int32
	Name           string
	Version        string
	UpdatedAt      time.Time
	CreatedAt      time.Time
}

type DirectoryService struct {
	ID                   int32
	OrganizationID       int32
	Name                 string
	DocumentationUrl     sql.NullString
	ApiSpecificationType sql.NullString
	Internal             bool
	TechSupportContact   sql.NullString
	PublicSupportContact sql.NullString
	OneTimeCosts         int32
	MonthlyCosts         int32
	RequestCosts         int32
}
