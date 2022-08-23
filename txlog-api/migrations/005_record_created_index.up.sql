BEGIN transaction;

CREATE INDEX records_index_created ON transactionlog.records USING btree (created);

COMMIT:
