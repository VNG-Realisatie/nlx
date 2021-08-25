BEGIN TRANSACTION;

ALTER TABLE nlx_management.audit_logs ADD COLUMN delegatee VARCHAR(250) NULL DEFAULT NULL;
ALTER TABLE nlx_management.audit_logs DROP COLUMN organization;
ALTER TABLE nlx_management.audit_logs DROP COLUMN service;
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
     'organization_settings_update',
     'organization_insight_configuration_update'
  )
);

CREATE TABLE nlx_management.audit_logs_services(
  audit_log_service_id BIGSERIAL PRIMARY KEY,
  audit_log_id BIGSERIAL,
  organization VARCHAR(250),
  service VARCHAR(250),
  created_at timestamp with time zone NOT NULL,

  CONSTRAINT fk_audit_logs
    FOREIGN KEY (audit_log_id)
    REFERENCES nlx_management.audit_logs (audit_log_id)
    ON DELETE RESTRICT
);

COMMIT;
