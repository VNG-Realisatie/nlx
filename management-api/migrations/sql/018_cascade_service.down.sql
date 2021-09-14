/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */
BEGIN transaction;

ALTER TABLE nlx_management.incoming_orders_services DROP CONSTRAINT fk_incoming_orders_services_incoming_order;

ALTER TABLE nlx_management.incoming_orders_services ADD CONSTRAINT fk_order
    FOREIGN KEY (incoming_order_id)
        REFERENCES nlx_management.incoming_orders (id)
        ON DELETE RESTRICT
        
COMMIT;
