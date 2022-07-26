/*
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

BEGIN transaction;

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

INSERT INTO nlx_management.permissions (code) VALUES ('permissions.outway.delete');
INSERT INTO nlx_management.permissions_roles (role_code, permission_code) VALUES ('admin', 'permissions.outway.delete');

COMMIT;
