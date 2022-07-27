// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package queries

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
)

type NlxManagementAccessGrant struct {
	ID                      int32
	AccessRequestIncomingID int32
	CreatedAt               time.Time
	RevokedAt               sql.NullTime
}

type NlxManagementAccessProof struct {
	ID                      int32
	AccessRequestOutgoingID int32
	CreatedAt               time.Time
	RevokedAt               sql.NullTime
}

type NlxManagementAccessRequestsIncoming struct {
	ID                       int32
	ServiceID                int32
	OrganizationName         string
	PublicKeyFingerprint     string
	State                    string
	CreatedAt                time.Time
	UpdatedAt                time.Time
	PublicKeyPem             sql.NullString
	OrganizationSerialNumber string
}

type NlxManagementAccessRequestsOutgoing struct {
	ID                       int32
	OrganizationName         string
	ServiceName              string
	State                    string
	PublicKeyFingerprint     string
	ReferenceID              int32
	ErrorCode                int32
	ErrorCause               sql.NullString
	ErrorStackTrace          []byte
	LockID                   uuid.NullUUID
	LockExpiresAt            sql.NullTime
	CreatedAt                time.Time
	UpdatedAt                time.Time
	PublicKeyPem             sql.NullString
	OrganizationSerialNumber string
	SynchronizeAt            time.Time
}

type NlxManagementAuditLog struct {
	ID         int64
	UserName   sql.NullString
	ActionType string
	UserAgent  string
	Data       pqtype.NullRawMessage
	CreatedAt  time.Time
}

type NlxManagementAuditLogsService struct {
	ID                       int64
	AuditLogID               sql.NullInt64
	OrganizationName         sql.NullString
	Service                  sql.NullString
	CreatedAt                time.Time
	OrganizationSerialNumber string
}

type NlxManagementIncomingOrder struct {
	ID           int64
	Description  string
	PublicKeyPem string
	Delegator    string
	Reference    string
	ValidFrom    time.Time
	ValidUntil   time.Time
	CreatedAt    time.Time
	RevokedAt    sql.NullTime
}

type NlxManagementIncomingOrdersService struct {
	IncomingOrderID          int64
	Service                  string
	OrganizationName         string
	OrganizationSerialNumber string
}

type NlxManagementInway struct {
	ID          int32
	Name        string
	SelfAddress string
	Version     string
	Hostname    string
	IpAddress   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NlxManagementInwaysService struct {
	InwayID   int32
	ServiceID int32
}

type NlxManagementOutgoingOrder struct {
	ID           int64
	Description  string
	PublicKeyPem string
	Delegatee    string
	Reference    string
	ValidFrom    time.Time
	ValidUntil   time.Time
	CreatedAt    time.Time
	RevokedAt    sql.NullTime
}

type NlxManagementOutgoingOrdersAccessProof struct {
	OutgoingOrderID int64
	AccessProofID   int64
}

type NlxManagementOutway struct {
	ID                   int32
	Name                 string
	PublicKeyPem         string
	Version              string
	IpAddress            pqtype.Inet
	CreatedAt            time.Time
	UpdatedAt            time.Time
	PublicKeyFingerprint sql.NullString
	SelfAddressApi       string
}

type NlxManagementPermission struct {
	Code string
}

type NlxManagementPermissionsRole struct {
	RoleCode       string
	PermissionCode string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type NlxManagementRole struct {
	Code string
}

type NlxManagementService struct {
	ID                   int32
	Name                 string
	EndpointUrl          string
	DocumentationUrl     string
	ApiSpecificationUrl  string
	Internal             bool
	TechSupportContact   string
	PublicSupportContact string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	OneTimeCosts         int32
	MonthlyCosts         int32
	RequestCosts         int32
}

type NlxManagementSetting struct {
	ID                       int32
	InwayID                  sql.NullInt32
	CreatedAt                time.Time
	UpdatedAt                time.Time
	OrganizationEmailAddress sql.NullString
}

type NlxManagementTermsOfService struct {
	ID        int32
	Username  string
	CreatedAt time.Time
}

type NlxManagementUser struct {
	ID        int32
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Password  sql.NullString
}

type NlxManagementUsersRole struct {
	UserID    int32
	RoleCode  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
