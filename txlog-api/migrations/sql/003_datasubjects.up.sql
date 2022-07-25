BEGIN transaction;

ALTER TABLE transactionlog.records ALTER COLUMN id TYPE bigint;

CREATE TABLE transactionlog.datasubjects (
	id bigserial NOT NULL,
	record_id bigint NOT NULL,
	"key" character varying(100) NOT NULL,
	"value" character varying(100) NOT NULL,
    CONSTRAINT datasubjects_pk PRIMARY KEY (id),
    CONSTRAINT datasubjects_uq_key UNIQUE (record_id, key)
);
CREATE INDEX datasubjects_index_keyvalue ON transactionlog.datasubjects USING btree (key, value);

ALTER TABLE transactionlog.datasubjects ADD CONSTRAINT datasubjects_fk_record FOREIGN KEY (record_id) REFERENCES transactionlog.records(id) MATCH FULL;

COMMIT;
