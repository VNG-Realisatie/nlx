BEGIN transaction;

ALTER TABLE nlx_management.outgoing_orders RENAME TO orders;
ALTER TABLE nlx_management.outgoing_orders_services RENAME TO orders_services;

ALTER INDEX nlx_management.idx_outgoing_orders_delegatee RENAME TO idx_delegatee;
ALTER INDEX nlx_management.idx_outgoing_orders_reference RENAME TO idx_reference;

ALTER TABLE nlx_management.orders_services RENAME COLUMN outgoing_order_id TO order_id;

DROP INDEX nlx_management.idx_outgoing_orders_services;
CREATE UNIQUE INDEX idx_orders_services ON nlx_management.orders_services (order_id, service, organization);

ALTER INDEX nlx_management.idx_outgoing_orders_service_name RENAME TO idx_services_name;
ALTER INDEX nlx_management.idx_outgoing_orders_organization_name RENAME TO idx_organization_name;

COMMIT;
