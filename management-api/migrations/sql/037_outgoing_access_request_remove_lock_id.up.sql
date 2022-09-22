-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

ALTER TABLE
    nlx_management.access_requests_outgoing
    DROP COLUMN lock_id,
    DROP COLUMN lock_expires_at,
    DROP COLUMN synchronize_at;

COMMIT;
