BEGIN transaction;

CREATE TABLE nlx_management.outgoing_orders_access_proofs
(
    order_id BIGINT,
    access_proof_id BIGINT,
    CONSTRAINT fk_order
        FOREIGN KEY (order_id)
            REFERENCES nlx_management.outgoing_orders (id)
            ON DELETE RESTRICT,
    CONSTRAINT fk_access_proof
        FOREIGN KEY (access_proof_id)
            REFERENCES nlx_management.access_proofs (id)
            ON DELETE RESTRICT,
    PRIMARY KEY(order_id, access_proof_id)
);

CREATE INDEX idx_outgoing_order_access_proofs_order_id ON nlx_management.outgoing_orders_access_proofs (order_id);
CREATE INDEX idx_outgoing_order_access_proofs_access_proof_id ON nlx_management.outgoing_orders_access_proofs (access_proof_id);

INSERT INTO nlx_management.outgoing_orders_access_proofs (order_id, access_proof_id)
SELECT
    os.outgoing_order_id, ap.id
FROM
    nlx_management.access_proofs ap
        JOIN
    nlx_management.access_requests_outgoing on (nlx_management.access_requests_outgoing.id = ap.access_request_outgoing_id)
        JOIN
    nlx_management.outgoing_orders_services os on (os.service = nlx_management.access_requests_outgoing.service_name AND os.organization_serial_number = nlx_management.access_requests_outgoing.organization_serial_number)
WHERE
    ap.revoked_at IS NULL;

DROP TABLE nlx_management.outgoing_orders_services;

COMMIT;
