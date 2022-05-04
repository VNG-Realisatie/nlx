-- name: ClearOrganizationInway :execrows
UPDATE directory.organizations
    SET inway_id = null
WHERE serial_number = $1;

-- name: GetInway :one
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
    inways.name = $2;
;

-- name: GetService :one
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
        services.id = $1;

-- name: SelectInwayByAddress :one
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
    o.serial_number = $2;

-- name: SetOrganizationInway :exec
update
    directory.organizations
set
    inway_id = $1
where
    serial_number = $2;

-- name: SetOrganizationEmail :exec
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
    returning id;

-- name: SelectOrganizationInwayAddress :one
select
    i.address
from
    directory.organizations o
left join
        directory.inways i
            on
                o.inway_id = i.id
where
    o.serial_number = $1;

-- name: SelectOrganizationInwayManagementAPIProxyAddress :one
select
    i.management_api_proxy_address
from
    directory.organizations as o
left join
        directory.inways i
            on
                o.inway_id = i.id
where
    o.serial_number = $1;

-- name: SelectVersionStatistics :many
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
desc;

-- name: SelectOrganizations :many
select
    serial_number,
    name
from
    directory.organizations
order by
    name;
