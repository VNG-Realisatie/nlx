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
    access_grants.terminated_at,
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

-- name: GetAccessGrantIDOfIncomingAccessRequest :one
select
    access_grants.id
from
    nlx_management.access_grants
where
    access_grants.access_request_incoming_id = $1
ORDER BY
    access_grants.id DESC
LIMIT 1
;

-- name: ListAccessGrantsForService :many
select
    access_grants.id,
    access_grants.created_at,
    access_grants.revoked_at,
    access_grants.terminated_at,
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
    access_grants.terminated_at,
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

-- name: TerminateAccessGrant :exec
update
    nlx_management.access_grants
set
    terminated_at = $1
where
    access_grants.id = $2
;

-- name: UpdateIncomingAccessRequest :execrows
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

-- name: GetAccessProof :one
select
    id, access_request_outgoing_id, created_at, revoked_at, terminated_at
FROM
    nlx_management.access_proofs
WHERE
    id = $1;

-- name: TerminateAccessProof :exec
update
    nlx_management.access_proofs
set
    terminated_at = $1
where
        access_proofs.id = $2
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

-- name: GetUserByEmail :one
select id, email, password, users_roles.role_code
from
    nlx_management.users
left join nlx_management.users_roles on users.id = users_roles.user_id
where email = $1
limit 1;

-- name: ListRolesForUser :many
select
    role_code
from
    nlx_management.users_roles
where user_id = $1;

-- name: ListPermissionsForRole :many
select permission_code
from
    nlx_management.permissions_roles
where role_code = $1;

-- name: CreateUser :one
insert into
    nlx_management.users
    (email, password, created_at, updated_at)
values
    ($1, $2, $3, $4)
returning id;

-- name: CreateUserRoles :exec
insert into
    nlx_management.users_roles
(user_id, role_code, created_at, updated_at)
values
    ($1, $2, $3, $4);

-- name: ListServices :many
select
    id,
    name,
    endpoint_url,
    documentation_url,
    api_specification_url,
    internal,
    tech_support_contact,
    public_support_contact,
    created_at,
    updated_at,
    one_time_costs,
    monthly_costs,
    request_costs
from
    nlx_management.services;

-- name: ListInwaysForService :many
select
    inways.id,
    inways.name,
    inways.self_address,
    inways.version,
    inways.hostname,
    inways.ip_address,
    inways.created_at,
    inways.updated_at
from
    nlx_management.inways_services
join
    nlx_management.inways on inways_services.inway_id = inways.id
where
    inways_services.service_id = $1;

-- name: DeleteOutgoingAccessRequest :exec
delete from
    nlx_management.access_requests_outgoing
where
    access_requests_outgoing.id = $1;

-- name: DeleteIncomingAccessRequest :exec
delete from
    nlx_management.access_requests_incoming
where
        access_requests_incoming.id = $1;

-- name: SetAuditLogAsSucceeded :exec
update
    nlx_management.audit_logs
set
    has_succeeded = true
where
    audit_logs.id = $1;

-- name: UpdateOutgoingAccessRequestState :execrows
update
    nlx_management.access_requests_outgoing
set
    state = $1,
    updated_at = $2
where
    access_requests_outgoing.id = $3;

-- name: ListLatestOutgoingAccessRequests :many
SELECT
    distinct on (
        public_key_fingerprint,
        service_name,
        organization_serial_number
    )
    access_requests_outgoing.id,
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
where
    organization_serial_number = $1 and
    service_name = $2
order by
    organization_serial_number,
    public_key_fingerprint,
    service_name,
    created_at
desc;

-- name: GetOutgoingAccessRequest :one
select
    access_requests_outgoing.id,
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
where
    id = $1;

-- name: GetAccessProofByOutgoingAccessRequest :one
select
    id,
    created_at,
    revoked_at,
    terminated_at
from
    nlx_management.access_proofs
where
    access_request_outgoing_id = $1
limit 1;

-- name: RevokeAccessProof :exec
update
    nlx_management.access_proofs
set
    revoked_at = $1
where
    id = $2
;

-- name: ListIncomingAccessRequests :many
select
    access_requests_incoming.id,
    access_requests_incoming.organization_name,
    access_requests_incoming.organization_serial_number,
    access_requests_incoming.state,
    access_requests_incoming.created_at,
    access_requests_incoming.updated_at
from
    nlx_management.access_requests_incoming
left join
    nlx_management.services on
        access_requests_incoming.service_id = services.id
where
    services.name = $1
;

-- name: GetLatestIncomingAccessRequest :one
select
    access_requests_incoming.id,
    access_requests_incoming.state,
    access_requests_incoming.created_at,
    access_requests_incoming.updated_at
from
    nlx_management.access_requests_incoming
        left join
    nlx_management.services on
            access_requests_incoming.service_id = services.id
where
    access_requests_incoming.organization_serial_number = $1
and
    access_requests_incoming.public_key_fingerprint = $2
and
    services.name = $3
order by
    access_requests_incoming.created_at desc
;

-- name: GetIncomingAccessRequestsByServiceCount :many
select
    count(access_requests_incoming.id),
    services.name
from
    nlx_management.access_requests_incoming
right join
    nlx_management.services on
        access_requests_incoming.service_id = services.id and
        access_requests_incoming.state = 'received'
group by
    services.id
;

-- name: GetIncomingAccessRequest :one
select
    access_requests_incoming.id,
    access_requests_incoming.state,
    access_requests_incoming.organization_name,
    access_requests_incoming.organization_serial_number,
    access_requests_incoming.public_key_fingerprint,
    access_requests_incoming.public_key_pem,
    access_requests_incoming.created_at,
    access_requests_incoming.updated_at,
    services.id as service_id
from
    nlx_management.access_requests_incoming
        join
            nlx_management.services on
                access_requests_incoming.service_id = services.id
where
    access_requests_incoming.id = $1
;

-- name: CreateIncomingAccessRequest :one
insert into
    nlx_management.access_requests_incoming
(
    state,
    organization_name,
    organization_serial_number,
    public_key_fingerprint,
    public_key_pem,
    service_id,
    created_at,
    updated_at
) values (
     $1,
     $2,
     $3,
     $4,
     $5,
     $6,
     $7,
     $8
 )
returning id
;

-- name: CountReceivedOutgoingAccessRequestsForOutway :one
select
    count(*) as count
from
    nlx_management.access_requests_outgoing
where
    access_requests_outgoing.organization_serial_number = $1 and
    access_requests_outgoing.service_name = $2 and
    access_requests_outgoing.public_key_fingerprint = $3 and
    access_requests_outgoing.state = 'received'
;

-- name: CreateOutgoingAccessRequest :one
insert into
    nlx_management.access_requests_outgoing
(
    state,
    organization_name,
    organization_serial_number,
    public_key_fingerprint,
    public_key_pem,
    service_name,
    created_at,
    updated_at
) values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
returning id
;

-- name: GetLatestOutgoingAccessRequest :one
select
    access_requests_outgoing.id,
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
where
    access_requests_outgoing.organization_serial_number = $1 and
    access_requests_outgoing.service_name = $2 and
    access_requests_outgoing.public_key_fingerprint = $3
order by
    access_requests_outgoing.created_at desc
limit 1
;

-- name: DeleteOutgoingAccessRequests :exec
delete from
    nlx_management.access_requests_outgoing
where
    access_requests_outgoing.organization_serial_number = $1 and
    access_requests_outgoing.service_name = $2
;

-- name: CreateAuditLog :one
insert into
    nlx_management.audit_logs
(
    user_name,
    action_type,
    user_agent,
    data,
    created_at,
    has_succeeded
) values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
 )
returning id
;

-- name: CreateAuditLogService :exec
insert into
    nlx_management.audit_logs_services
(
    audit_log_id,
    organization_name,
    organization_serial_number,
    created_at,
    service
) values (
    $1,
    $2,
    $3,
    $4,
    $5
)
;

-- name: ListAuditLogs :many
select
    audit_logs.id,
    audit_logs.user_name,
    audit_logs.action_type,
    audit_logs.user_agent,
    audit_logs.data,
    audit_logs.created_at,
    audit_logs.has_succeeded
from
    nlx_management.audit_logs
order by
    audit_logs.created_at desc
limit $1
;

-- name: ListAuditLogServices :many
select
    audit_logs_services.organization_name,
    audit_logs_services.organization_serial_number,
    audit_logs_services.service,
    audit_logs_services.created_at
from
    nlx_management.audit_logs_services
where
    audit_logs_services.audit_log_id = $1
;

-- name: UpdateInway :exec
update
    nlx_management.inways
set
    name = $1,
    version = $2,
    hostname = $3,
    self_address = $4,
    created_at = $5,
    updated_at = $6
where
    inways.id = $7
;

-- name: RemoveInwayServicesForInway :exec
delete from
    nlx_management.inways_services
where
    inway_id = $1
;

-- name: RemoveInway :exec
delete from
    nlx_management.inways
where
        id = $1
;

-- name: ListInways :many
select
    inways.id,
    inways.name,
    inways.self_address,
    inways.version,
    inways.hostname,
    inways.ip_address,
    inways.created_at,
    inways.updated_at
from
    nlx_management.inways
;

-- name: ListServicesForInway :many
select
    id,
    name,
    endpoint_url,
    documentation_url,
    api_specification_url,
    internal,
    tech_support_contact,
    public_support_contact,
    created_at,
    updated_at,
    one_time_costs,
    monthly_costs,
    request_costs,
    inways_services.inway_id
from
    nlx_management.services
join
    nlx_management.inways_services
        on inways_services.service_id = services.id
where
    inways_services.inway_id = $1
;

-- name: GetInwayByName :one
select
    inways.id,
    inways.name,
    inways.self_address,
    inways.version,
    inways.hostname,
    inways.ip_address,
    inways.created_at,
    inways.updated_at
from
    nlx_management.inways
where
    inways.name = $1
;
