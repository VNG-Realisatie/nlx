CREATE TABLE nlx_management.audit_logs(
  audit_log_id BIGSERIAL PRIMARY KEY,
  user_name VARCHAR(250),
  action_type VARCHAR(250) NOT NULL,
  user_agent VARCHAR(250) NOT NULL,
  organization VARCHAR(250),
  service VARCHAR(250),
  data JSON,
  created_at timestamp with time zone NOT NULL
);

CREATE INDEX audit_log_idx_organization ON nlx_management.audit_logs (organization);
CREATE INDEX audit_log_idx_service ON nlx_management.audit_logs (service);

ALTER TABLE nlx_management.audit_logs 
ADD CONSTRAINT audit_log_ck_action_type
CHECK (
	 action_type = 'login_success'
	OR action_type = 'login_fail'
	OR action_type = 'logout_success'
	OR action_type = 'incoming_access_request_accept'
	OR action_type = 'incoming_access_request_reject'
	OR action_type = 'access_grant_revoke'
	OR action_type = 'outgoing_access_request_create'
	OR action_type = 'service_create'
	OR action_type = 'service_update'
	OR action_type = 'service_delete'
	OR action_type = 'organization_settings_update'
	OR action_type = 'organization_insight_configuration_update'
);
