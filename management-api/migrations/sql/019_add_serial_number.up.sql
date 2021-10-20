/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */
BEGIN transaction;

ALTER TABLE nlx_management.access_requests_incoming ADD COLUMN organization_serial_number VARCHAR(20) NOT NULL;
ALTER TABLE nlx_management.access_requests_outgoing ADD COLUMN organization_serial_number VARCHAR(20) NOT NULL;

-- delete access requests, since we dont have the organization serial number in the old records
DELETE FROM nlx_management.access_requests_incoming;
DELETE FROM nlx_management.access_requests_outgoing;

-- delete existing orders, since we dont have the serial numbers of the organization for every service
DELETE FROM nlx_management.outgoing_orders;
DELETE FROM nlx_management.incoming_orders;

ALTER TABLE nlx_management.incoming_orders_services RENAME COLUMN organization TO organization_name;
ALTER TABLE nlx_management.incoming_orders_services ADD COLUMN organization_serial_number VARCHAR(20) NOT NULL;

ALTER TABLE nlx_management.outgoing_orders_services RENAME COLUMN organization TO organization_name;
ALTER TABLE nlx_management.outgoing_orders_services ADD COLUMN organization_serial_number VARCHAR(20) NOT NULL;

ALTER TABLE nlx_management.audit_logs_services RENAME COLUMN organization TO organization_name;
ALTER TABLE nlx_management.audit_logs_services ADD COLUMN organization_serial_number VARCHAR(20) NOT NULL;

COMMIT;
