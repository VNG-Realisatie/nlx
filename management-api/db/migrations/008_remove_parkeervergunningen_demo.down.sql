BEGIN transaction;

ALTER TABLE nlx_management.settings ADD COLUMN insight_api_url VARCHAR(250) NOT NULL DEFAULT '';
ALTER TABLE nlx_management.settings ADD COLUMN irma_server_url VARCHAR(250) NOT NULL DEFAULT '';

COMMIT;
