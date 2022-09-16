-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

ALTER TABLE directory.organizations DROP COLUMN insight_log_endpoint;
ALTER TABLE directory.organizations DROP COLUMN insight_irma_endpoint;

COMMIT;
