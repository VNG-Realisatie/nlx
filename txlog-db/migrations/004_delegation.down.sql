BEGIN transaction;

ALTER TABLE transactionlog.records DROP COLUMN delegator;
ALTER TABLE transactionlog.records DROP COLUMN order_reference;

COMMIT;
