-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

UPDATE nlx_management.access_requests_outgoing SET public_key_pem = '' WHERE public_key_pem IS NULL;
ALTER TABLE nlx_management.access_requests_outgoing ALTER COLUMN public_key_pem SET DEFAULT '';
ALTER TABLE nlx_management.access_requests_outgoing ALTER COLUMN public_key_pem SET NOT NULL;

UPDATE nlx_management.access_requests_incoming SET public_key_pem = '' WHERE public_key_pem IS NULL;
ALTER TABLE nlx_management.access_requests_incoming ALTER COLUMN public_key_pem SET DEFAULT '';
ALTER TABLE nlx_management.access_requests_incoming ALTER COLUMN public_key_pem SET NOT NULL;

COMMIT;
