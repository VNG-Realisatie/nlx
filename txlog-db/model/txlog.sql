-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.1
-- PostgreSQL version: 10.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- object: "nlx-org-txlog-api" | type: ROLE --
-- DROP ROLE IF EXISTS "nlx-org-txlog-api";
CREATE ROLE "nlx-org-txlog-api" WITH 
	LOGIN
	ENCRYPTED PASSWORD 'nlx-txlog-api';
-- ddl-end --

-- object: "nlx-org-txlog-writer" | type: ROLE --
-- DROP ROLE IF EXISTS "nlx-org-txlog-writer";
CREATE ROLE "nlx-org-txlog-writer" WITH 
	LOGIN
	ENCRYPTED PASSWORD 'nlx-txlog-writer';
-- ddl-end --


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: txlog | type: DATABASE --
-- -- DROP DATABASE IF EXISTS txlog;
-- CREATE DATABASE txlog;
-- -- ddl-end --
-- 

-- object: transactionlog | type: SCHEMA --
-- DROP SCHEMA IF EXISTS transactionlog CASCADE;
CREATE SCHEMA transactionlog;
-- ddl-end --
ALTER SCHEMA transactionlog OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,transactionlog;
-- ddl-end --

-- object: transactionlog.direction | type: TYPE --
-- DROP TYPE IF EXISTS transactionlog.direction CASCADE;
CREATE TYPE transactionlog.direction AS
 ENUM ('in','out');
-- ddl-end --
ALTER TYPE transactionlog.direction OWNER TO postgres;
-- ddl-end --

-- object: transactionlog.records_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS transactionlog.records_id_seq CASCADE;
CREATE SEQUENCE transactionlog.records_id_seq
	INCREMENT BY 1
	MINVALUE 0
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE transactionlog.records_id_seq OWNER TO postgres;
-- ddl-end --

-- object: transactionlog.records | type: TABLE --
-- DROP TABLE IF EXISTS transactionlog.records CASCADE;
CREATE TABLE transactionlog.records(
	id bigint NOT NULL DEFAULT nextval('records_id_seq'),
	direction transactionlog.direction NOT NULL,
	created timestamptz NOT NULL DEFAULT now(),
	src_organization text NOT NULL,
	dest_organization text NOT NULL,
	service_name text NOT NULL,
	logrecord_id text NOT NULL,
	data jsonb,
	CONSTRAINT records_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE transactionlog.records OWNER TO postgres;
-- ddl-end --

-- Appended SQL commands --
ALTER SEQUENCE transactionlog.records_id_seq OWNED BY records.id;

-- ddl-end --

-- object: transactionlog.datasubjects_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS transactionlog.datasubjects_id_seq CASCADE;
CREATE SEQUENCE transactionlog.datasubjects_id_seq
	INCREMENT BY 1
	MINVALUE 0
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE transactionlog.datasubjects_id_seq OWNER TO postgres;
-- ddl-end --

-- object: transactionlog.datasubjects | type: TABLE --
-- DROP TABLE IF EXISTS transactionlog.datasubjects CASCADE;
CREATE TABLE transactionlog.datasubjects(
	id bigint NOT NULL DEFAULT nextval('datasubjects_id_seq'),
	record_id integer NOT NULL,
	key varchar(100) NOT NULL,
	value varchar(100) NOT NULL,
	CONSTRAINT datasubjects_pk PRIMARY KEY (id),
	CONSTRAINT datasubjects_uq_key UNIQUE (record_id,key)

);
-- ddl-end --
ALTER TABLE transactionlog.datasubjects OWNER TO postgres;
-- ddl-end --

-- object: datasubjects_index_keyvalue | type: INDEX --
-- DROP INDEX IF EXISTS transactionlog.datasubjects_index_keyvalue CASCADE;
CREATE INDEX datasubjects_index_keyvalue ON transactionlog.datasubjects
	USING btree
	(
	  key,
	  value ASC NULLS LAST
	);
-- ddl-end --

-- object: datasubjects_fk_record | type: CONSTRAINT --
-- ALTER TABLE transactionlog.datasubjects DROP CONSTRAINT IF EXISTS datasubjects_fk_record CASCADE;
ALTER TABLE transactionlog.datasubjects ADD CONSTRAINT datasubjects_fk_record FOREIGN KEY (record_id)
REFERENCES transactionlog.records (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: grant_44d1d56c5a | type: PERMISSION --
GRANT SELECT
   ON TABLE transactionlog.records
   TO "nlx-org-txlog-api";
-- ddl-end --

-- object: grant_c011d34c83 | type: PERMISSION --
GRANT SELECT,INSERT
   ON TABLE transactionlog.records
   TO "nlx-org-txlog-writer";
-- ddl-end --

-- object: grant_51c82c7170 | type: PERMISSION --
GRANT USAGE
   ON SCHEMA transactionlog
   TO "nlx-org-txlog-api";
-- ddl-end --

-- object: grant_04e581e244 | type: PERMISSION --
GRANT USAGE
   ON SCHEMA transactionlog
   TO "nlx-org-txlog-writer";
-- ddl-end --

-- object: grant_aa562a9360 | type: PERMISSION --
GRANT SELECT,USAGE
   ON SEQUENCE transactionlog.records_id_seq
   TO "nlx-org-txlog-writer";
-- ddl-end --

-- object: grant_de75a7700e | type: PERMISSION --
GRANT SELECT
   ON TABLE transactionlog.datasubjects
   TO "nlx-org-txlog-api";
-- ddl-end --

-- object: grant_2fba446269 | type: PERMISSION --
GRANT INSERT
   ON TABLE transactionlog.datasubjects
   TO "nlx-org-txlog-writer";
-- ddl-end --

-- object: grant_5e6736b63c | type: PERMISSION --
GRANT SELECT,USAGE
   ON SEQUENCE transactionlog.datasubjects_id_seq
   TO "nlx-org-txlog-writer";
-- ddl-end --


