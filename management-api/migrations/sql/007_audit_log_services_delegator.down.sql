BEGIN TRANSACTION;

ALTER TABLE nlx_management.audit_logs DROP COLUMN delegatee;
ALTER TABLE nlx_management.audit_logs ADD COLUMN service VARCHAR(250) NULL DEFAULT NULL;
ALTER TABLE nlx_management.audit_logs ADD COLUMN organization VARCHAR(250) NULL DEFAULT NULL;

ALTER TABLE nlx_management.audit_logs DROP CONSTRAINT audit_log_ck_action_type;
ALTER TABLE nlx_management.audit_logs ADD CONSTRAINT audit_log_ck_action_type
CHECK (
	 action_type IN (
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
     'organization_settings_update',
     'organization_insight_configuration_update'
  )
);

DROP TABLE nlx_management.audit_logs_services;

COMMIT;
