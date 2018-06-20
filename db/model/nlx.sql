-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.1
-- PostgreSQL version: 10.0
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: nlx | type: DATABASE --
-- -- DROP DATABASE IF EXISTS nlx;
-- CREATE DATABASE nlx;
-- -- ddl-end --
-- 

-- object: directory | type: SCHEMA --
-- DROP SCHEMA IF EXISTS directory CASCADE;
CREATE SCHEMA directory;
-- ddl-end --
ALTER SCHEMA directory OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,directory;
-- ddl-end --

-- object: directory.organizations | type: TABLE --
-- DROP TABLE IF EXISTS directory.organizations CASCADE;
CREATE TABLE directory.organizations(
	id serial NOT NULL,
	name varchar(100) NOT NULL,
	CONSTRAINT organizations_pk PRIMARY KEY (id),
	CONSTRAINT organizations_uq_name UNIQUE (name)

);
-- ddl-end --
ALTER TABLE directory.organizations OWNER TO postgres;
-- ddl-end --

-- object: directory.services | type: TABLE --
-- DROP TABLE IF EXISTS directory.services CASCADE;
CREATE TABLE directory.services(
	id serial NOT NULL,
	organization_id integer NOT NULL,
	name varchar(100) NOT NULL,
	documentation_url varchar(250),
	api_specification_type varchar(20),
	CONSTRAINT services_pk PRIMARY KEY (id),
	CONSTRAINT services_uq_name UNIQUE (organization_id,name),
	CONSTRAINT services_check_typespec CHECK (api_specification_type IS NULL OR (
	api_specification_type = 'OpenAPI2'
	OR
	api_specification_type = 'OpenAPI3'
))

);
-- ddl-end --
ALTER TABLE directory.services OWNER TO postgres;
-- ddl-end --

-- object: directory.inways | type: TABLE --
-- DROP TABLE IF EXISTS directory.inways CASCADE;
CREATE TABLE directory.inways(
	id serial NOT NULL,
	organization_id integer NOT NULL,
	address varchar(100) NOT NULL,
	CONSTRAINT inways_pk PRIMARY KEY (id),
	CONSTRAINT inways_uq_address UNIQUE (organization_id,address)

);
-- ddl-end --
ALTER TABLE directory.inways OWNER TO postgres;
-- ddl-end --

-- object: directory.availabilities | type: TABLE --
-- DROP TABLE IF EXISTS directory.availabilities CASCADE;
CREATE TABLE directory.availabilities(
	id serial NOT NULL,
	inway_id integer NOT NULL,
	service_id integer NOT NULL,
	healthy bool NOT NULL DEFAULT false,
	CONSTRAINT availabilities_pk PRIMARY KEY (id),
	CONSTRAINT availabilities_uq_inway_service UNIQUE (inway_id,service_id)

);
-- ddl-end --
ALTER TABLE directory.availabilities OWNER TO postgres;
-- ddl-end --

-- object: services_fk_organization | type: CONSTRAINT --
-- ALTER TABLE directory.services DROP CONSTRAINT IF EXISTS services_fk_organization CASCADE;
ALTER TABLE directory.services ADD CONSTRAINT services_fk_organization FOREIGN KEY (organization_id)
REFERENCES directory.organizations (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: inways_fk_organization | type: CONSTRAINT --
-- ALTER TABLE directory.inways DROP CONSTRAINT IF EXISTS inways_fk_organization CASCADE;
ALTER TABLE directory.inways ADD CONSTRAINT inways_fk_organization FOREIGN KEY (organization_id)
REFERENCES directory.organizations (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: availabilities_fk_inway | type: CONSTRAINT --
-- ALTER TABLE directory.availabilities DROP CONSTRAINT IF EXISTS availabilities_fk_inway CASCADE;
ALTER TABLE directory.availabilities ADD CONSTRAINT availabilities_fk_inway FOREIGN KEY (inway_id)
REFERENCES directory.inways (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: availabilities_fk_service | type: CONSTRAINT --
-- ALTER TABLE directory.availabilities DROP CONSTRAINT IF EXISTS availabilities_fk_service CASCADE;
ALTER TABLE directory.availabilities ADD CONSTRAINT availabilities_fk_service FOREIGN KEY (service_id)
REFERENCES directory.services (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


