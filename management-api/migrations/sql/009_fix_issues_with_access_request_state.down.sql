BEGIN transaction;

ALTER TABLE nlx_management.access_grants RENAME COLUMN id TO access_grant_id;
ALTER TABLE nlx_management.access_proofs RENAME COLUMN id TO access_proof_id;
ALTER TABLE nlx_management.access_requests_incoming RENAME COLUMN id TO access_request_incoming_id;
ALTER TABLE nlx_management.access_requests_outgoing RENAME COLUMN id TO access_request_outgoing_id;
ALTER TABLE nlx_management.audit_logs RENAME COLUMN id TO audit_log_id;
ALTER TABLE nlx_management.audit_logs_services RENAME COLUMN id TO audit_log_service_id;
ALTER TABLE nlx_management.inways RENAME COLUMN id TO inway_id;
ALTER TABLE nlx_management.orders RENAME COLUMN id TO order_id;
ALTER TABLE nlx_management.permissions RENAME COLUMN id TO permission_id;
ALTER TABLE nlx_management.roles RENAME COLUMN id TO role_id;
ALTER TABLE nlx_management.services RENAME COLUMN id TO service_id;
ALTER TABLE nlx_management.settings RENAME COLUMN id TO settings_id;
ALTER TABLE nlx_management.users RENAME COLUMN id TO user_id;

COMMIT;
