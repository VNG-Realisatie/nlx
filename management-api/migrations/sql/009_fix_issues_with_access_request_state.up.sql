BEGIN transaction;

ALTER TABLE nlx_management.access_grants RENAME COLUMN access_grant_id TO id;
ALTER TABLE nlx_management.access_proofs RENAME COLUMN access_proof_id TO id;
ALTER TABLE nlx_management.access_requests_incoming RENAME COLUMN access_request_incoming_id TO id;
ALTER TABLE nlx_management.access_requests_outgoing RENAME COLUMN access_request_outgoing_id TO id;
ALTER TABLE nlx_management.audit_logs RENAME COLUMN audit_log_id TO id;
ALTER TABLE nlx_management.audit_logs_services RENAME COLUMN audit_log_service_id TO id;
ALTER TABLE nlx_management.inways RENAME COLUMN inway_id TO id;
ALTER TABLE nlx_management.orders RENAME COLUMN order_id TO id;
ALTER TABLE nlx_management.permissions RENAME COLUMN permission_id TO id;
ALTER TABLE nlx_management.roles RENAME COLUMN role_id TO id;
ALTER TABLE nlx_management.services RENAME COLUMN service_id TO id;
ALTER TABLE nlx_management.settings RENAME COLUMN settings_id TO id;
ALTER TABLE nlx_management.users RENAME COLUMN user_id TO id;

COMMIT;
