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
	MAXVALUE 2147483647
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
	id integer NOT NULL DEFAULT nextval('records_id_seq'),
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

-- object: grant_e3a699271d | type: PERMISSION --
GRANT SELECT
   ON TABLE transactionlog.records
   TO "nlx-org-txlog-api";
-- ddl-end --

-- object: grant_09ea658de9 | type: PERMISSION --
GRANT INSERT
   ON TABLE transactionlog.records
   TO "nlx-org-txlog-writer";
-- ddl-end --

-- object: grant_a080fa2047 | type: PERMISSION --
GRANT USAGE
   ON SCHEMA transactionlog
   TO "nlx-org-txlog-api";
-- ddl-end --

-- object: grant_43c252f634 | type: PERMISSION --
GRANT USAGE
   ON SCHEMA transactionlog
   TO "nlx-org-txlog-writer";
-- ddl-end --

-- object: grant_1c54d8fa2f | type: PERMISSION --
GRANT SELECT,USAGE
   ON SEQUENCE transactionlog.records_id_seq
   TO "nlx-org-txlog-writer";
-- ddl-end --


