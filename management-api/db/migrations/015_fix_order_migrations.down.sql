begin;

DROP INDEX nlx_management.idx_outgoing_orders_delegatee_reference;
CREATE UNIQUE INDEX idx_outgoing_orders_reference ON nlx_management.outgoing_orders (reference);

commit;
