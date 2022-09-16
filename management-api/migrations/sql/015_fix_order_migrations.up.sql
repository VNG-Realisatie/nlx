-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

begin;

DROP INDEX nlx_management.idx_outgoing_orders_reference;
CREATE UNIQUE INDEX idx_outgoing_orders_delegatee_reference ON nlx_management.outgoing_orders (delegatee, reference);

commit;
