/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */

BEGIN transaction;

CREATE TABLE nlx_management.orders
(
    order_id        BIGSERIAL PRIMARY KEY,
    description     VARCHAR(100)             NOT NULL,
    public_key_pem  VARCHAR(4096)            NOT NULL,
    delegatee       VARCHAR(100)             NOT NULL,
    reference       VARCHAR(100)             NOT NULL,
    valid_from      timestamp with time zone NOT NULL,
    valid_until     timestamp with time zone NOT NULL,
    created_at      timestamp with time zone NOT NULL
);

CREATE INDEX idx_delegatee ON nlx_management.orders (delegatee);
CREATE UNIQUE INDEX idx_reference ON nlx_management.orders (reference);

CREATE TABLE nlx_management.orders_services
(
    order_id     BIGSERIAL    NOT NULL,
    service_name VARCHAR(100) NOT NULL,

    CONSTRAINT fk_order
        FOREIGN KEY (order_id)
            REFERENCES nlx_management.orders (order_id)
            ON DELETE RESTRICT
);

CREATE INDEX idx_service_name ON nlx_management.orders_services (service_name);
CREATE UNIQUE INDEX idx_orders_services ON nlx_management.orders_services (order_id, service_name);

COMMIT;
