BEGIN transaction;

ALTER TABLE nlx_management.orders RENAME TO outgoing_orders;
ALTER INDEX nlx_management.idx_delegatee RENAME TO idx_outgoing_orders_delegatee;
ALTER INDEX nlx_management.idx_reference RENAME TO idx_outgoing_orders_reference;

ALTER TABLE nlx_management.orders_services RENAME TO outgoing_orders_services;

ALTER TABLE nlx_management.outgoing_orders_services RENAME COLUMN order_id TO outgoing_order_id;

DROP INDEX nlx_management.idx_orders_services;
CREATE UNIQUE INDEX idx_outgoing_orders_services ON nlx_management.outgoing_orders_services (outgoing_order_id, service, organization);

ALTER INDEX nlx_management.idx_service_name RENAME TO idx_outgoing_orders_services_name;
ALTER INDEX nlx_management.idx_organization_name RENAME TO idx_outgoing_orders_organization_name;

COMMIT;
