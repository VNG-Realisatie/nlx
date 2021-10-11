/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */
BEGIN transaction;

ALTER TABLE nlx_management.access_requests_incoming DROP COLUMN organization_serial_number;
ALTER TABLE nlx_management.access_requests_outgoing DROP COLUMN organization_serial_number;

COMMIT;
