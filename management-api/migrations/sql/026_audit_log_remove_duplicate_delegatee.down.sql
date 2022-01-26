BEGIN transaction;

ALTER TABLE nlx_management.audit_logs ADD COLUMN delegatee VARCHAR(250) NULL DEFAULT NULL;

COMMIT;
