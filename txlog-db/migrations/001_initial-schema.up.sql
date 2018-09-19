CREATE SCHEMA transactionlog;

CREATE TYPE transactionlog.direction AS ENUM (
	'in',
	'out'
);

CREATE SEQUENCE transactionlog.records_id_seq
	INCREMENT BY 1
	MINVALUE 0
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE transactionlog.records_id_seq OWNER TO postgres;

CREATE TABLE transactionlog.records (
	id integer NOT NULL DEFAULT nextval('transactionlog.records_id_seq'),
	direction transactionlog.direction NOT NULL,
	created timestamp with time zone DEFAULT now() NOT NULL,
	src_organization text NOT NULL,
	dest_organization text NOT NULL,
	service_name text NOT NULL,
	logrecord_id text NOT NULL,
	"data" jsonb
);

ALTER SEQUENCE transactionlog.records_id_seq OWNED BY transactionlog.records.id;

ALTER TABLE transactionlog.records
	ADD CONSTRAINT records_pk PRIMARY KEY (id);

ALTER TABLE transactionlog.records OWNER TO postgres;
