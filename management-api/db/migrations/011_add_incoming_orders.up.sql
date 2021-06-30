/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */

BEGIN transaction;

CREATE TABLE nlx_management.incoming_orders
(
    order_id        BIGSERIAL PRIMARY KEY,
    description     VARCHAR(100)             NOT NULL,
    public_key_pem  VARCHAR(4096)            NOT NULL,
    delegator       VARCHAR(100)             NOT NULL,
    reference       VARCHAR(100)             NOT NULL,
    valid_from      timestamp with time zone NOT NULL,
    valid_until     timestamp with time zone NOT NULL,
    created_at      timestamp with time zone NOT NULL
);

CREATE INDEX idx_incoming_orders_delegator ON nlx_management.incoming_orders (delegator);
CREATE UNIQUE INDEX idx_incoming_orders_reference ON nlx_management.incoming_orders (reference);

CREATE TABLE nlx_management.incoming_orders_services
(
    order_id     BIGSERIAL    NOT NULL,
    service      VARCHAR(100) NOT NULL,
    organization VARCHAR(100) NOT NULL,

    CONSTRAINT fk_order
        FOREIGN KEY (order_id)
            REFERENCES nlx_management.incoming_orders (order_id)
            ON DELETE RESTRICT
);

CREATE INDEX idx_incoming_orders_service_name ON nlx_management.incoming_orders_services (service);
CREATE INDEX idx_incoming_orders_service_organization_name ON nlx_management.incoming_orders_services (organization);
CREATE UNIQUE INDEX idx_incoming_orders_services ON nlx_management.incoming_orders_services (order_id, service, organization);

COMMIT;
