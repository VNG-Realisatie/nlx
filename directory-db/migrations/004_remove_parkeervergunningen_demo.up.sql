BEGIN transaction;

ALTER TABLE directory.organizations DROP COLUMN insight_log_endpoint;
ALTER TABLE directory.organizations DROP COLUMN insight_irma_endpoint;

COMMIT;
