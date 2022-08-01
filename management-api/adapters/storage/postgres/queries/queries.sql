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

-- name: CountInwaysByName :one
select
    count(*)
from
    nlx_management.inways
where inways.name = sqlc.arg(inway_name)::text;

-- name: DeleteSettings :exec
delete from nlx_management.settings;

-- name: CreateSettings :exec
insert into
    nlx_management.settings
        (organization_email_address, inway_id)
    VALUES (
               sqlc.arg(organization_email_address)::text,
            (
                select
                    id
                from
                    nlx_management.inways
                where
                        inways.name = sqlc.arg(inway_name)::text
            )
    )
;
