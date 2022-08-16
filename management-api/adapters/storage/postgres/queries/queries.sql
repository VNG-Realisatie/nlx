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

-- name: GetAccessGrant :one
select
    access_grants.id,
    access_grants.created_at,
    access_grants.revoked_at,
    access_request_incoming_id,
    access_requests_incoming.id as ari_id,
    access_requests_incoming.service_id as ari_service_id,
    access_requests_incoming.organization_name as ari_organization_name,
    access_requests_incoming.organization_serial_number as ari_organization_serial_number,
    access_requests_incoming.state as ari_state,
    access_requests_incoming.created_at as ari_created_at,
    access_requests_incoming.updated_at as ari_updated_at,
    access_requests_incoming.public_key_fingerprint as ari_public_key_fingerprint,
    access_requests_incoming.public_key_pem as ari_public_key_pem,
    services.id as s_id,
    services.name as s_name,
    services.endpoint_url as s_endpoint_url,
    services.documentation_url as s_documentation_url,
    services.api_specification_url as s_api_specification_url,
    services.internal as s_internal,
    services.tech_support_contact as s_tech_support_contact,
    services.public_support_contact as s_public_support_contact,
    services.one_time_costs as s_one_time_costs,
    services.monthly_costs as s_monthly_costs,
    services.request_costs as s_request_costs,
    services.created_at as s_created_at,
    services.updated_at as s_updated_at
from
    nlx_management.access_grants
        left join nlx_management.access_requests_incoming on (
            access_grants.access_request_incoming_id = access_requests_incoming.id
        )
        left join nlx_management.services on (
            access_requests_incoming.service_id = services.id
        )
where
        access_grants.id = $1
;
