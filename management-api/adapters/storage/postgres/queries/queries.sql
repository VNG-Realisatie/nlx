-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

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
    access_requests_incoming.id as access_request_incoming_id,
    access_requests_incoming.service_id as access_request_incoming_service_id,
    access_requests_incoming.organization_name as access_request_incoming_organization_name,
    access_requests_incoming.organization_serial_number as access_request_incoming_organization_serial_number,
    access_requests_incoming.state as access_request_incoming_state,
    access_requests_incoming.created_at as access_request_incoming_created_at,
    access_requests_incoming.updated_at as access_request_incoming_updated_at,
    access_requests_incoming.public_key_fingerprint as access_request_incoming_public_key_fingerprint,
    access_requests_incoming.public_key_pem as access_request_incoming_public_key_pem,
    services.id as service_id,
    services.name as service_name,
    services.endpoint_url as service_endpoint_url,
    services.documentation_url as service_documentation_url,
    services.api_specification_url as service_api_specification_url,
    services.internal as service_internal,
    services.tech_support_contact as service_tech_support_contact,
    services.public_support_contact as service_public_support_contact,
    services.one_time_costs as service_one_time_costs,
    services.monthly_costs as service_monthly_costs,
    services.request_costs as service_request_costs,
    services.created_at as service_created_at,
    services.updated_at as service_updated_at
from
    nlx_management.access_grants
        join nlx_management.access_requests_incoming on (
            access_grants.access_request_incoming_id = access_requests_incoming.id
        )
        join nlx_management.services on (
            access_requests_incoming.service_id = services.id
        )
where
        access_grants.id = $1
;

-- name: ListAccessGrantsForService :many
select
    access_grants.id,
    access_grants.created_at,
    access_grants.revoked_at,
    access_grants.access_request_incoming_id,
    access_requests_incoming.service_id as access_request_incoming_service_id,
    access_requests_incoming.organization_name as access_request_incoming_organization_name,
    access_requests_incoming.organization_serial_number as access_request_incoming_organization_serial_number,
    access_requests_incoming.state as access_request_incoming_state,
    access_requests_incoming.created_at as access_request_incoming_created_at,
    access_requests_incoming.updated_at as access_request_incoming_updated_at,
    access_requests_incoming.public_key_fingerprint as access_request_incoming_public_key_fingerprint,
    access_requests_incoming.public_key_pem as access_request_incoming_public_key_pem,
    services.id as service_id,
    services.name as service_name,
    services.endpoint_url as service_endpoint_url,
    services.documentation_url as service_documentation_url,
    services.api_specification_url as service_api_specification_url,
    services.internal as service_internal,
    services.tech_support_contact as service_tech_support_contact,
    services.public_support_contact as service_public_support_contact,
    services.one_time_costs as service_one_time_costs,
    services.monthly_costs as service_monthly_costs,
    services.request_costs as service_request_costs,
    services.created_at as service_created_at,
    services.updated_at as service_updated_at
from
    nlx_management.access_grants
        join nlx_management.access_requests_incoming on (
            access_grants.access_request_incoming_id = access_requests_incoming.id
        )
        join nlx_management.services on (
            access_requests_incoming.service_id = services.id and
            services.name = $1
        )
;

-- name: GetLatestAccessGrantForService :one
select
    access_grants.id,
    access_grants.created_at,
    access_grants.revoked_at,
    access_requests_incoming.id as access_request_incoming_id,
    access_requests_incoming.service_id as access_request_incoming_service_id,
    access_requests_incoming.organization_name as access_request_incoming_organization_name,
    access_requests_incoming.organization_serial_number as access_request_incoming_organization_serial_number,
    access_requests_incoming.state as access_request_incoming_state,
    access_requests_incoming.created_at as access_request_incoming_created_at,
    access_requests_incoming.updated_at as access_request_incoming_updated_at,
    access_requests_incoming.public_key_fingerprint as access_request_incoming_public_key_fingerprint,
    access_requests_incoming.public_key_pem as access_request_incoming_public_key_pem,
    services.id as service_id,
    services.name as service_name,
    services.endpoint_url as service_endpoint_url,
    services.documentation_url as service_documentation_url,
    services.api_specification_url as service_api_specification_url,
    services.internal as service_internal,
    services.tech_support_contact as service_tech_support_contact,
    services.public_support_contact as service_public_support_contact,
    services.one_time_costs as service_one_time_costs,
    services.monthly_costs as service_monthly_costs,
    services.request_costs as service_request_costs,
    services.created_at as service_created_at,
    services.updated_at as service_updated_at
from
    nlx_management.access_grants
        join nlx_management.access_requests_incoming on (
            access_requests_incoming.id = access_grants.access_request_incoming_id and
            access_requests_incoming.organization_serial_number = $1 and
            access_requests_incoming.public_key_fingerprint = $2
        )
        join nlx_management.services on (
            services.id = access_requests_incoming.service_id and
            services.name = sqlc.arg(service_name)::text
        )
order by
    access_grants.created_at desc
limit 1
;

-- name: RevokeAccessGrant :exec
update
    nlx_management.access_grants
set
    revoked_at = $2
where
    access_grants.id = $1
;

-- name: UpdateIncomingAccessRequest :exec
update
    nlx_management.access_requests_incoming
set
    state = $2,
    updated_at = $3
where
    id = $1
;

-- name: CreateAccessProof :one
insert into
    nlx_management.access_proofs
(
    access_request_outgoing_id,
    created_at
) values (
             $1,
             $2
         )
returning id
;

-- name: UpsertInway :exec
insert into nlx_management.inways
    (name, self_address, version, hostname, ip_address, created_at, updated_at)
values
    ($1, $2, $3, $4, $5, $6, $7)
    on conflict
on constraint inways_name_key
    do update
           set
                self_address    = excluded.self_address,
                version         = excluded.version,
                hostname        = excluded.hostname,
                ip_address      = excluded.ip_address,
                updated_at      = excluded.updated_at;

-- name: ListAllLatestOutgoingAccessRequests :many
select
    distinct on (
        public_key_fingerprint,
        service_name,
        organization_serial_number
    ) access_requests_outgoing.id,
      access_requests_outgoing.organization_name,
      access_requests_outgoing.organization_serial_number,
      access_requests_outgoing.service_name,
      access_requests_outgoing.state,
      access_requests_outgoing.reference_id,
      access_requests_outgoing.error_code,
      access_requests_outgoing.error_cause,
      access_requests_outgoing.public_key_fingerprint,
      access_requests_outgoing.public_key_pem,
      access_requests_outgoing.created_at,
      access_requests_outgoing.updated_at
from
    nlx_management.access_requests_outgoing
order by
    organization_serial_number,
    public_key_fingerprint,
    service_name,
    created_at
desc;

-- name: ListTermsOfService :many
select
    id, username, created_at
from
    nlx_management.terms_of_service;

-- name: CreateTermsOfService :exec
insert into
    nlx_management.terms_of_service
(username, created_at)
values
    ($1, $2);

-- name: ListPermissions :many
select code from nlx_management.permissions;
