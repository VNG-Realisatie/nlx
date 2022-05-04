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

-- name: GetOutway :one
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
    outways.name = $2;

-- name: SelectParticipants :many
select
    serial_number,
    name,
    created_at,
    (select count(id) FROM directory.inways as i where i.organization_id = o.id) as inways,
    (select count(id) FROM directory.outways as ow where ow.organization_id = o.id) as outways,
    (select count(id) FROM directory.services as s where s.organization_id = o.id) as services
from
    directory.organizations as o;

-- name: RegisterInway :exec
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
        updated_at                   = excluded.updated_at;

-- name: RegisterService :one
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
    ) select id from service;

-- name: SelectServices :many
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
    s.name;

-- name: RegisterOutway :exec
with organization as (
    insert into
        directory.organizations
        (
            serial_number,
            name
        )
        values (
            $1,
            $2
        )
    on conflict on constraint organizations_uq_serial_number
        do update set
            serial_number = excluded.serial_number,
            name          = excluded.name
    returning id
)
insert into
    directory.outways
    (
        name,
        organization_id,
        version,
        created_at,
        updated_at
    )
    select
        $3,
        organization.id,
        nullif($4, ''),
        $5,
        $6
    from
        organization
    on conflict (
        name,
        organization_id
    )
    do update set
        name       = excluded.name,
        version    = excluded.version,
        updated_at = excluded.updated_at;
