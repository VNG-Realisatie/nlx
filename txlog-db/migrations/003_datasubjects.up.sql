
CREATE SEQUENCE transactionlog.datasubjects_id_seq
	INCREMENT BY 1
	MINVALUE 0
	NO MAXVALUE
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE transactionlog.datasubjects_id_seq OWNER TO postgres;

CREATE TABLE transactionlog.datasubjects (
	id bigint NOT NULL DEFAULT nextval('transactionlog.datasubjects_id_seq'),
	record_id integer NOT NULL,
	"key" character varying(100) NOT NULL,
	"value" character varying(100) NOT NULL
);

ALTER SEQUENCE transactionlog.datasubjects_id_seq OWNED BY transactionlog.datasubjects.id;

ALTER TABLE transactionlog.datasubjects OWNER TO postgres;

ALTER TABLE transactionlog.datasubjects
	ADD CONSTRAINT datasubjects_pk PRIMARY KEY (id);

ALTER TABLE transactionlog.datasubjects
	ADD CONSTRAINT datasubjects_uq_key UNIQUE (record_id, key);

ALTER TABLE transactionlog.datasubjects
	ADD CONSTRAINT datasubjects_fk_record FOREIGN KEY (record_id) REFERENCES transactionlog.records(id) MATCH FULL;

CREATE INDEX datasubjects_index_keyvalue ON transactionlog.datasubjects USING btree (key, value);

GRANT SELECT ON TABLE transactionlog.datasubjects TO "nlx-org-txlog-api";
GRANT INSERT ON TABLE transactionlog.datasubjects TO "nlx-org-txlog-writer";
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA transactionlog TO "nlx-org-txlog-writer";

-- Upgrade records.id to bigint with no max value on the sequence
ALTER TABLE transactionlog.records
	ALTER COLUMN id TYPE bigint USING id::bigint /* TYPE change - table: records original: integer new: bigint */;

ALTER SEQUENCE transactionlog.records_id_seq
	NO MAXVALUE;

-- writer now uses `RETURNING id`, so needs to be granted SELECT as well.
REVOKE ALL ON TABLE transactionlog.records FROM "nlx-org-txlog-writer";
GRANT SELECT, INSERT ON TABLE transactionlog.records TO "nlx-org-txlog-writer";
