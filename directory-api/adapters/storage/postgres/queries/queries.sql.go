// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: queries.sql

package queries

import (
	"context"
	"database/sql"
	"time"
)

const clearOrganizationInway = `-- name: ClearOrganizationInway :execrows
UPDATE directory.organizations
    SET inway_id = null
WHERE serial_number = $1
`

func (q *Queries) ClearOrganizationInway(ctx context.Context, serialNumber string) (int64, error) {
	result, err := q.exec(ctx, q.clearOrganizationInwayStmt, clearOrganizationInway, serialNumber)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getInway = `-- name: GetInway :one
select
    inways.name as name,
    address,
    version as nlx_version,
    inways.created_at as created_at,
    updated_at,
    organizations.serial_number as organization_serial_number,
    organizations.name as organization_name
from
    directory.inways
        join directory.organizations
            on directory.inways.organization_id = directory.organizations.id
where
    organizations.serial_number = $1
  and
    inways.name = $2
`

type GetInwayParams struct {
	SerialNumber string
	Name         string
}

type GetInwayRow struct {
	Name                     string
	Address                  string
	NlxVersion               string
	CreatedAt                time.Time
	UpdatedAt                time.Time
	OrganizationSerialNumber string
	OrganizationName         string
}

func (q *Queries) GetInway(ctx context.Context, arg *GetInwayParams) (*GetInwayRow, error) {
	row := q.queryRow(ctx, q.getInwayStmt, getInway, arg.SerialNumber, arg.Name)
	var i GetInwayRow
	err := row.Scan(
		&i.Name,
		&i.Address,
		&i.NlxVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.OrganizationSerialNumber,
		&i.OrganizationName,
	)
	return &i, err
}

const getOutway = `-- name: GetOutway :one
select
    outways.name as name,
    version as nlx_version,
    outways.created_at as created_at,
    updated_at,
    organizations.serial_number as organization_serial_number,
    organizations.name as organization_name
from
    directory.outways
         join directory.organizations
              on outways.organization_id = organizations.id
where
    organizations.serial_number = $1
  and
    outways.name = $2
`

type GetOutwayParams struct {
	SerialNumber string
	Name         string
}

type GetOutwayRow struct {
	Name                     string
	NlxVersion               string
	CreatedAt                time.Time
	UpdatedAt                time.Time
	OrganizationSerialNumber string
	OrganizationName         string
}

func (q *Queries) GetOutway(ctx context.Context, arg *GetOutwayParams) (*GetOutwayRow, error) {
	row := q.queryRow(ctx, q.getOutwayStmt, getOutway, arg.SerialNumber, arg.Name)
	var i GetOutwayRow
	err := row.Scan(
		&i.Name,
		&i.NlxVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.OrganizationSerialNumber,
		&i.OrganizationName,
	)
	return &i, err
}

const getService = `-- name: GetService :one
select
    services.id as id,
    services.name as name,
    documentation_url,
    api_specification_type,
    internal,
    tech_support_contact,
    public_support_contact,
    organizations.serial_number as organization_serial_number,
    organizations.name as organization_name,
    one_time_costs,
    monthly_costs,
    request_costs
from directory.services
         join directory.organizations
              on services.organization_id = organizations.id
where
        services.id = $1
`

type GetServiceRow struct {
	ID                       int32
	Name                     string
	DocumentationUrl         sql.NullString
	ApiSpecificationType     sql.NullString
	Internal                 bool
	TechSupportContact       sql.NullString
	PublicSupportContact     sql.NullString
	OrganizationSerialNumber string
	OrganizationName         string
	OneTimeCosts             int32
	MonthlyCosts             int32
	RequestCosts             int32
}

func (q *Queries) GetService(ctx context.Context, id int32) (*GetServiceRow, error) {
	row := q.queryRow(ctx, q.getServiceStmt, getService, id)
	var i GetServiceRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DocumentationUrl,
		&i.ApiSpecificationType,
		&i.Internal,
		&i.TechSupportContact,
		&i.PublicSupportContact,
		&i.OrganizationSerialNumber,
		&i.OrganizationName,
		&i.OneTimeCosts,
		&i.MonthlyCosts,
		&i.RequestCosts,
	)
	return &i, err
}

const registerInway = `-- name: RegisterInway :exec
with organization as (
    insert into directory.organizations
                (serial_number, name)
         values ($1, $2)
    on conflict
        on constraint organizations_uq_serial_number
        do update
            set
                serial_number = excluded.serial_number,
                name          = excluded.name
        returning id
)
insert into
    directory.inways (
      name,
      organization_id,
      address,
      management_api_proxy_address,
      version,
      created_at,
      updated_at
    )
    select
        $3,
        organization.id,
        $4,
        $5,
        nullif($6, ''),
        $7,
        $8
    from
        organization
    on conflict (
        name,
        organization_id
    ) do update set
        name                         = excluded.name,
        address                      = excluded.address,
        management_api_proxy_address = excluded.management_api_proxy_address,
        version                      = excluded.version,
        updated_at                   = excluded.updated_at
`

type RegisterInwayParams struct {
	SerialNumber              string
	Name                      string
	Name_2                    string
	Address                   string
	ManagementApiProxyAddress sql.NullString
	Column6                   interface{}
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

func (q *Queries) RegisterInway(ctx context.Context, arg *RegisterInwayParams) error {
	_, err := q.exec(ctx, q.registerInwayStmt, registerInway,
		arg.SerialNumber,
		arg.Name,
		arg.Name_2,
		arg.Address,
		arg.ManagementApiProxyAddress,
		arg.Column6,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const registerService = `-- name: RegisterService :one
with organization as (
    select
        id
    from
        directory.organizations
    where
        organizations.serial_number = $1
    ),
    inway as (
        select
            inways.id
        from
            directory.inways,
            organization
        where
            organization_id = organization.id
    ),
    service as (
        insert into
            directory.services
            (
                organization_id,
                name,
                internal,
                documentation_url,
                api_specification_type,
                public_support_contact,
                tech_support_contact,
                request_costs,
                monthly_costs,
                one_time_costs
            )
            select
                organization.id,
                $2,
                $3,
                nullif($4, ''),
                nullif($5, ''),
                nullif($6, ''),
                nullif($7, ''),
                $8,
                $9,
                $10
            from
                organization
            on conflict on constraint services_uq_name
                do update set
                    internal = excluded.internal,
                    documentation_url = excluded.documentation_url,
                    api_specification_type = excluded.api_specification_type,
                    public_support_contact = excluded.public_support_contact,
                    tech_support_contact = excluded.tech_support_contact,
                    request_costs = excluded.request_costs,
                    monthly_costs = excluded.monthly_costs,
                    one_time_costs = excluded.one_time_costs
            returning id
    ),
    availabilities as (
        insert into
            directory.availabilities (
                inway_id,
                service_id,
                last_announced
            )
            select
                inway.id,
                service.id,
                now()
            from
                inway,
                service
        on conflict on constraint availabilities_uq_inway_service
        do update set
            last_announced = now(),
            active = true
    ) select id from service
`

type RegisterServiceParams struct {
	SerialNumber string
	Name         string
	Internal     bool
	Column4      interface{}
	Column5      interface{}
	Column6      interface{}
	Column7      interface{}
	RequestCosts int32
	MonthlyCosts int32
	OneTimeCosts int32
}

func (q *Queries) RegisterService(ctx context.Context, arg *RegisterServiceParams) (int32, error) {
	row := q.queryRow(ctx, q.registerServiceStmt, registerService,
		arg.SerialNumber,
		arg.Name,
		arg.Internal,
		arg.Column4,
		arg.Column5,
		arg.Column6,
		arg.Column7,
		arg.RequestCosts,
		arg.MonthlyCosts,
		arg.OneTimeCosts,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const selectInwayByAddress = `-- name: SelectInwayByAddress :one
select
    i.id as inway_id,
    i.organization_id
from
    directory.inways as i
         inner join directory.organizations o
             on o.id = i.organization_id
where
    i.address = $1
  and
    o.serial_number = $2
`

type SelectInwayByAddressParams struct {
	Address      string
	SerialNumber string
}

type SelectInwayByAddressRow struct {
	InwayID        int32
	OrganizationID int32
}

func (q *Queries) SelectInwayByAddress(ctx context.Context, arg *SelectInwayByAddressParams) (*SelectInwayByAddressRow, error) {
	row := q.queryRow(ctx, q.selectInwayByAddressStmt, selectInwayByAddress, arg.Address, arg.SerialNumber)
	var i SelectInwayByAddressRow
	err := row.Scan(&i.InwayID, &i.OrganizationID)
	return &i, err
}

const selectOrganizationInwayAddress = `-- name: SelectOrganizationInwayAddress :one
select
    i.address
from
    directory.organizations o
left join
        directory.inways i
            on
                o.inway_id = i.id
where
    o.serial_number = $1
`

func (q *Queries) SelectOrganizationInwayAddress(ctx context.Context, serialNumber string) (sql.NullString, error) {
	row := q.queryRow(ctx, q.selectOrganizationInwayAddressStmt, selectOrganizationInwayAddress, serialNumber)
	var address sql.NullString
	err := row.Scan(&address)
	return address, err
}

const selectOrganizationInwayManagementAPIProxyAddress = `-- name: SelectOrganizationInwayManagementAPIProxyAddress :one
select
    i.management_api_proxy_address
from
    directory.organizations as o
left join
        directory.inways i
            on
                o.inway_id = i.id
where
    o.serial_number = $1
`

func (q *Queries) SelectOrganizationInwayManagementAPIProxyAddress(ctx context.Context, serialNumber string) (sql.NullString, error) {
	row := q.queryRow(ctx, q.selectOrganizationInwayManagementAPIProxyAddressStmt, selectOrganizationInwayManagementAPIProxyAddress, serialNumber)
	var management_api_proxy_address sql.NullString
	err := row.Scan(&management_api_proxy_address)
	return management_api_proxy_address, err
}

const selectOrganizations = `-- name: SelectOrganizations :many
select
    serial_number,
    name
from
    directory.organizations
order by
    name
`

type SelectOrganizationsRow struct {
	SerialNumber string
	Name         string
}

func (q *Queries) SelectOrganizations(ctx context.Context) ([]*SelectOrganizationsRow, error) {
	rows, err := q.query(ctx, q.selectOrganizationsStmt, selectOrganizations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*SelectOrganizationsRow{}
	for rows.Next() {
		var i SelectOrganizationsRow
		if err := rows.Scan(&i.SerialNumber, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectParticipants = `-- name: SelectParticipants :many
select
    serial_number,
    name,
    created_at,
    (select count(id) FROM directory.inways as i where i.organization_id = o.id) as inways,
    (select count(id) FROM directory.outways as ow where ow.organization_id = o.id) as outways,
    (select count(id) FROM directory.services as s where s.organization_id = o.id) as services
from
    directory.organizations as o
`

type SelectParticipantsRow struct {
	SerialNumber string
	Name         string
	CreatedAt    time.Time
	Inways       int64
	Outways      int64
	Services     int64
}

func (q *Queries) SelectParticipants(ctx context.Context) ([]*SelectParticipantsRow, error) {
	rows, err := q.query(ctx, q.selectParticipantsStmt, selectParticipants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*SelectParticipantsRow{}
	for rows.Next() {
		var i SelectParticipantsRow
		if err := rows.Scan(
			&i.SerialNumber,
			&i.Name,
			&i.CreatedAt,
			&i.Inways,
			&i.Outways,
			&i.Services,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectServices = `-- name: SelectServices :many
select
    o.serial_number as organization_serial_number,
    o.name AS organization_name,
    s.name AS name,
    s.internal as internal,
    s.one_time_costs as one_time_costs,
    s.monthly_costs as monthly_costs,
    s.request_costs as request_costs,
    array_remove(array_agg(i.address), NULL) as inway_addresses,
    coalesce(s.documentation_url, '') as documentation_url,
    coalesce(s.api_specification_type, '') as api_specification_type,
    coalesce(s.public_support_contact, '') as public_support_contact,
    array_remove(array_agg(a.healthy), NULL) as healthy_statuses
from
    directory.services s
         inner join directory.availabilities a on a.service_id = s.id
         inner join directory.organizations o on o.id = s.organization_id
         inner join directory.inways i on i.id = a.inway_id
where (
    internal = false
    or (
        internal = true and
        o.serial_number = $1
   )
)
group by
    s.id,
    o.id
order by
    o.name,
    s.name
`

type SelectServicesRow struct {
	OrganizationSerialNumber string
	OrganizationName         string
	Name                     string
	Internal                 bool
	OneTimeCosts             int32
	MonthlyCosts             int32
	RequestCosts             int32
	InwayAddresses           interface{}
	DocumentationUrl         string
	ApiSpecificationType     string
	PublicSupportContact     string
	HealthyStatuses          interface{}
}

func (q *Queries) SelectServices(ctx context.Context, serialNumber string) ([]*SelectServicesRow, error) {
	rows, err := q.query(ctx, q.selectServicesStmt, selectServices, serialNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*SelectServicesRow{}
	for rows.Next() {
		var i SelectServicesRow
		if err := rows.Scan(
			&i.OrganizationSerialNumber,
			&i.OrganizationName,
			&i.Name,
			&i.Internal,
			&i.OneTimeCosts,
			&i.MonthlyCosts,
			&i.RequestCosts,
			&i.InwayAddresses,
			&i.DocumentationUrl,
			&i.ApiSpecificationType,
			&i.PublicSupportContact,
			&i.HealthyStatuses,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectVersionStatistics = `-- name: SelectVersionStatistics :many
select
    'outway' AS type,
    version,
    count(*) as amount
from
    directory.outways
group by
    version
union
select
    'inway' AS type,
    version,
    count(*) as amount
from
    directory.inways
group by
    version
order by
    type,
    version
desc
`

type SelectVersionStatisticsRow struct {
	Type    interface{}
	Version string
	Amount  int64
}

func (q *Queries) SelectVersionStatistics(ctx context.Context) ([]*SelectVersionStatisticsRow, error) {
	rows, err := q.query(ctx, q.selectVersionStatisticsStmt, selectVersionStatistics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*SelectVersionStatisticsRow{}
	for rows.Next() {
		var i SelectVersionStatisticsRow
		if err := rows.Scan(&i.Type, &i.Version, &i.Amount); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setOrganizationEmail = `-- name: SetOrganizationEmail :exec
insert into
    directory.organizations
    (serial_number, name, email_address)
values
    ($1, $2, $3)
    on conflict
on constraint organizations_uq_serial_number
    do update set
        serial_number = excluded.serial_number,
        name 		  = excluded.name,
        email_address = excluded.email_address
    returning id
`

type SetOrganizationEmailParams struct {
	SerialNumber string
	Name         string
	EmailAddress sql.NullString
}

func (q *Queries) SetOrganizationEmail(ctx context.Context, arg *SetOrganizationEmailParams) error {
	_, err := q.exec(ctx, q.setOrganizationEmailStmt, setOrganizationEmail, arg.SerialNumber, arg.Name, arg.EmailAddress)
	return err
}

const setOrganizationInway = `-- name: SetOrganizationInway :exec
update
    directory.organizations
set
    inway_id = $1
where
    serial_number = $2
`

type SetOrganizationInwayParams struct {
	InwayID      sql.NullInt32
	SerialNumber string
}

func (q *Queries) SetOrganizationInway(ctx context.Context, arg *SetOrganizationInwayParams) error {
	_, err := q.exec(ctx, q.setOrganizationInwayStmt, setOrganizationInway, arg.InwayID, arg.SerialNumber)
	return err
}
