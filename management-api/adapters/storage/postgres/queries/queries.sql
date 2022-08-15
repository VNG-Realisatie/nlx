-- name: GetSettings :one
select
    settings.organization_email_address,
    inways.name
from
    nlx_management.settings
        left join nlx_management.inways on (
            settings.inway_id = inways.id
        )
order by
    settings.id
limit 1;

-- name: DoesInwayExistByName :one
select
        count(*)>0 as inway_exits
from
    nlx_management.inways
where
    inways.name = sqlc.arg(inway_name)::text;

-- name: UpdateSettings :exec
update
    nlx_management.settings
set
        organization_email_address = sqlc.arg(organization_email_address)::text,
        inway_id = (
            select
                id
            from
                nlx_management.inways
            where
                    inways.name = sqlc.arg(inway_name)::text
        )
;

-- name: CreateAccessGrant :one
insert into
    nlx_management.access_grants
(
    access_request_incoming_id,
    created_at
) values (
      $1,
      $2
)
returning id
;
