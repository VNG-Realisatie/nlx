/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */
BEGIN transaction;

ALTER TABLE nlx_management.access_grants DROP CONSTRAINT fk_access_grants_incoming_access_request;

ALTER TABLE nlx_management.access_grants ADD CONSTRAINT fk_access_grants_incoming_access_request
    FOREIGN KEY (access_request_incoming_id)
    REFERENCES nlx_management.access_requests_incoming (id)
    ON DELETE CASCADE;

ALTER TABLE nlx_management.access_proofs DROP CONSTRAINT fk_access_proofs_outgoing_access_request;

ALTER TABLE nlx_management.access_proofs ADD CONSTRAINT fk_access_proofs_outgoing_access_request
    FOREIGN KEY (access_request_outgoing_id)
    REFERENCES nlx_management.access_requests_outgoing (id)
    ON DELETE CASCADE;

COMMIT;
