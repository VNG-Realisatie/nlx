-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.1
-- PostgreSQL version: 10.0
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: nlx_logdb | type: DATABASE --
-- -- DROP DATABASE IF EXISTS nlx_logdb;
-- CREATE DATABASE nlx_logdb;
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

-- object: transactionlog.records | type: TABLE --
-- DROP TABLE IF EXISTS transactionlog.records CASCADE;
CREATE TABLE transactionlog.records(
	id serial NOT NULL,
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


