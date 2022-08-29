BEGIN transaction;

ALTER TABLE transactionlog.records
RENAME COLUMN logrecord_id TO transaction_id;

COMMIT;
