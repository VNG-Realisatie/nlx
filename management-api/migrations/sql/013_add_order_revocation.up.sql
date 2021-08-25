BEGIN transaction;

ALTER TABLE nlx_management.outgoing_orders ADD COLUMN revoked_at timestamp with time zone NULL DEFAULT NULL;
ALTER TABLE nlx_management.incoming_orders ADD COLUMN revoked_at timestamp with time zone NULL DEFAULT NULL;

DROP INDEX nlx_management.idx_incoming_orders_reference;
CREATE UNIQUE INDEX idx_incoming_orders_delegator_reference ON nlx_management.incoming_orders (delegator, reference);

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
     'order_incoming_revoke',
     'organization_settings_update',
     'organization_insight_configuration_update'
  )
);

COMMIT;
