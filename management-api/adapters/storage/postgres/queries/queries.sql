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
