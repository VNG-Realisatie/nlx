BEGIN transaction;

ALTER TABLE nlx_management.outgoing_orders DROP COLUMN revoked_at;
ALTER TABLE nlx_management.incoming_orders DROP COLUMN revoked_at;

DROP INDEX nlx_management.idx_incoming_orders_delegator_reference;
CREATE UNIQUE INDEX idx_incoming_orders_reference ON nlx_management.incoming_orders (reference);

DROP INDEX nlx_management.idx_outgoing_orders_delegator_reference;
CREATE UNIQUE INDEX idx_outgoing_orders_reference ON nlx_management.outgoing_orders (reference);

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
     'order_create',
     'organization_settings_update',
     'organization_insight_configuration_update'
  )
);

COMMIT;
