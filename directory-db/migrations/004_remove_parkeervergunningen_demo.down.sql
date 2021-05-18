BEGIN transaction;

ALTER TABLE directory.organizations ADD COLUMN insight_log_endpoint character varying(250);
ALTER TABLE directory.organizations ADD COLUMN insight_irma_endpoint character varying(250);

COMMIT;
