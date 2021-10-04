/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */
BEGIN transaction;

ALTER TABLE nlx_management.access_requests_incoming ADD COLUMN organization_serial_number CHAR(20) NOT NULL;

COMMIT;
