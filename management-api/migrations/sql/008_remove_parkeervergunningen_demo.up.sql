-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

ALTER TABLE nlx_management.settings DROP COLUMN insight_api_url;
ALTER TABLE nlx_management.settings DROP COLUMN irma_server_url;

COMMIT;
