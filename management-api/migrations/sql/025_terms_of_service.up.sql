begin transaction;

CREATE TABLE nlx_management.terms_of_service (
  id SERIAL PRIMARY KEY,
  username VARCHAR(250) NOT NULL,
  created_at timestamp with time zone NOT NULL
);

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
     'accept_terms_of_service'
  )
);

commit;
