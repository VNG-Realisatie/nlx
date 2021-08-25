/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */

BEGIN transaction;

DELETE FROM nlx_management.audit_logs WHERE action_type = 'organization_insight_configuration_update';

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
     'organization_settings_update'
  )
);

COMMIT;
