BEGIN transaction;

ALTER TABLE nlx_management.access_requests_incoming DROP COLUMN public_key_pem;
ALTER TABLE nlx_management.access_requests_outgoing DROP COLUMN public_key_pem;

COMMIT;
