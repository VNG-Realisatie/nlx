-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

INSERT INTO nlx_management.permissions (code) VALUES ('permissions.outgoing_access_request.withdraw');
INSERT INTO nlx_management.permissions (code) VALUES ('permissions.access.terminate');
INSERT INTO nlx_management.permissions_roles (role_code, permission_code)
VALUES
    ('admin', 'permissions.outgoing_access_request.withdraw'),
    ('admin', 'permissions.access.terminate');

ALTER TABLE nlx_management.audit_logs DROP CONSTRAINT audit_log_ck_action_type;
ALTER TABLE nlx_management.audit_logs ADD CONSTRAINT audit_log_ck_action_type
    CHECK (action_type IN (
                           'login_success',
                           'login_fail',
                           'logout_success',
                           'incoming_access_request_accept',
                           'incoming_access_request_reject',
                           'access_grant_revoke',
                           'outgoing_access_request_create',
                           'outgoing_access_request_withdraw',
                           'access_terminate',
                           'service_create',
                           'service_update',
                           'service_delete',
                           'order_create',
                           'order_outgoing_revoke',
                           'organization_settings_update',
                           'inway_delete',
                           'order_outgoing_update',
                           'accept_terms_of_service',
                           'outway_delete'
        )
        );

ALTER TABLE nlx_management.audit_logs ADD COLUMN has_succeeded BOOL NOT NULL DEFAULT false;

UPDATE nlx_management.audit_logs SET has_succeeded = true;

UPDATE nlx_management.access_requests_outgoing
SET
    state = 'approved'
WHERE
    state = 'revoked';

UPDATE nlx_management.access_requests_incoming
SET
    state = 'approved'
WHERE
    state = 'revoked';

ALTER TABLE nlx_management.access_requests_outgoing DROP CONSTRAINT access_requests_outgoing_state_check;

ALTER TABLE nlx_management.access_requests_outgoing ADD CONSTRAINT access_requests_outgoing_state_check
    CHECK (state IN (
                           'failed',
                           'received',
                           'approved',
                           'rejected',
                           'withdrawn'
        ) );

ALTER TABLE nlx_management.access_requests_incoming DROP CONSTRAINT access_requests_incoming_state_check;

ALTER TABLE nlx_management.access_requests_incoming ADD CONSTRAINT access_requests_incoming_state_check
    CHECK (state IN (
                     'received',
                     'approved',
                     'rejected',
                     'withdrawn'
        ) );

ALTER TABLE nlx_management.access_grants ADD COLUMN terminated_at timestamp with time zone NULL DEFAULT NULL;
ALTER TABLE nlx_management.access_proofs ADD COLUMN terminated_at timestamp with time zone NULL DEFAULT NULL;

COMMIT;
