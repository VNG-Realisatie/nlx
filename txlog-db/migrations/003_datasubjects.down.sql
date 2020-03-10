BEGIN transaction;

ALTER TABLE transactionlog.datasubjects DROP CONSTRAINT datasubjects_fk_record;
DROP TABLE transactionlog.datasubjects;

ALTER TABLE transactionlog.records ALTER COLUMN id TYPE int;

COMMIT;