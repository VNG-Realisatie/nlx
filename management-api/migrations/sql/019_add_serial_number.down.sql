/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */
BEGIN transaction;

ALTER TABLE nlx_management.access_requests_incoming DROP COLUMN organization_serial_number;
ALTER TABLE nlx_management.access_requests_outgoing DROP COLUMN organization_serial_number;

ALTER TABLE nlx_management.incoming_orders_services RENAME COLUMN organization_name TO organization;
ALTER TABLE nlx_management.incoming_orders_services DROP COLUMN organization_serial_number;

ALTER TABLE nlx_management.outgoing_orders_services RENAME COLUMN organization_name TO organization;
ALTER TABLE nlx_management.outgoing_orders_services DROP COLUMN organization_serial_number;

ALTER TABLE nlx_management.audit_logs_services RENAME COLUMN organization_name TO organization;
ALTER TABLE nlx_management.audit_logs_services DROP COLUMN organization_serial_number;

COMMIT;
