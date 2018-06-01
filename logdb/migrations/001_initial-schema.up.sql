CREATE SCHEMA transactionlog;

CREATE TYPE transactionlog.direction AS ENUM (
	'in',
	'out'
);

CREATE TABLE transactionlog.records (
	id serial NOT NULL,
	direction transactionlog.direction NOT NULL,
	created timestamp with time zone DEFAULT now() NOT NULL,
	src_orgnization text NOT NULL,
	dest_organization text NOT NULL,
	service_name text NOT NULL,
	request_path text NOT NULL,
	"data" jsonb
);

ALTER TABLE transactionlog.records
	ADD CONSTRAINT records_pk PRIMARY KEY (id);

ALTER TABLE transactionlog.records OWNER TO postgres;
