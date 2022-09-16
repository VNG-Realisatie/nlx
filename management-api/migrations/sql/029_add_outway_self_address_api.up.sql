-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

ALTER TABLE nlx_management.outways ADD COLUMN self_address_api VARCHAR(255) NOT NULL DEFAULT '';

COMMIT;
