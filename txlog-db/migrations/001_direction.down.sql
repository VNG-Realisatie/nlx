BEGIN transaction;

DROP TABLE transactionlog.records;
DROP TYPE transactionlog.direction;

DROP SCHEMA transactionlog;

COMMIT;
