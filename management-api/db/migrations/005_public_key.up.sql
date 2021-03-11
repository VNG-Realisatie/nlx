BEGIN transaction;

ALTER TABLE nlx_management.access_requests_incoming ADD COLUMN public_key_pem VARCHAR(4096) NULL DEFAULT NULL;
ALTER TABLE nlx_management.access_requests_outgoing ADD COLUMN public_key_pem VARCHAR(4096) NULL DEFAULT NULL;

COMMIT;
