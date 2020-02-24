-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.1
-- PostgreSQL version: 10.0
-- Project Site: pgmodeler.io
-- Model Author: ---

SET check_function_bodies = false;
-- ddl-end --

-- object: "nlx-directory" | type: ROLE --
-- DROP ROLE IF EXISTS "nlx-directory";
CREATE ROLE "nlx-directory" WITH 
	LOGIN
	ENCRYPTED PASSWORD 'nlx-directory';
-- ddl-end --


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
	insight_log_endpoint varchar(250),
	insight_irma_endpoint varchar(250),
	CONSTRAINT organizations_pk PRIMARY KEY (id),
	CONSTRAINT organizations_uq_name UNIQUE (name)

);
-- ddl-end --
ALTER TABLE directory.organizations OWNER TO postgres;
-- ddl-end --

-- Appended SQL commands --
GRANT USAGE, SELECT ON SEQUENCE organizations_id_seq TO "nlx-directory";
-- ddl-end --

-- object: directory.services | type: TABLE --
-- DROP TABLE IF EXISTS directory.services CASCADE;
CREATE TABLE directory.services(
	id serial NOT NULL,
	organization_id integer NOT NULL,
	name varchar(250) NOT NULL,
	internal boolean NOT NULL DEFAULT false,
	documentation_url varchar(250),
	api_specification_type varchar(20),
	tech_support_contact varchar(250),
	public_support_contact varchar(250),
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

-- Appended SQL commands --
GRANT USAGE, SELECT ON SEQUENCE services_id_seq TO "nlx-directory";
-- ddl-end --

-- object: directory.inways | type: TABLE --
-- DROP TABLE IF EXISTS directory.inways CASCADE;
CREATE TABLE directory.inways(
	id serial NOT NULL,
	organization_id integer NOT NULL,
	address varchar(100) NOT NULL,
	version varchar(100) NOT NULL DEFAULT 'unknown',
	CONSTRAINT inways_pk PRIMARY KEY (id),
	CONSTRAINT inways_uq_address UNIQUE (organization_id,address)

);
-- ddl-end --
ALTER TABLE directory.inways OWNER TO postgres;
-- ddl-end --

-- Appended SQL commands --
GRANT USAGE, SELECT ON SEQUENCE inways_id_seq TO "nlx-directory";
-- ddl-end --

-- object: directory.availabilities | type: TABLE --
-- DROP TABLE IF EXISTS directory.availabilities CASCADE;
CREATE TABLE directory.availabilities(
	id serial NOT NULL,
	inway_id integer NOT NULL,
	service_id integer NOT NULL,
	healthy bool NOT NULL DEFAULT false,
	unhealthy_since timestamptz,
	last_announced timestamptz NOT NULL DEFAULT NOW(),
	active bool NOT NULL DEFAULT false,
	CONSTRAINT availabilities_pk PRIMARY KEY (id),
	CONSTRAINT availabilities_uq_inway_service UNIQUE (inway_id,service_id)

);
-- ddl-end --
ALTER TABLE directory.availabilities OWNER TO postgres;
-- ddl-end --

-- Appended SQL commands --
GRANT USAGE, SELECT ON SEQUENCE availabilities_id_seq TO "nlx-directory";
-- ddl-end --

-- object: directory.availabilities_verify | type: FUNCTION --
-- DROP FUNCTION IF EXISTS directory.availabilities_verify() CASCADE;
CREATE FUNCTION directory.availabilities_verify ()
	RETURNS trigger
	LANGUAGE plpgsql
	VOLATILE 
	CALLED ON NULL INPUT
	SECURITY INVOKER
	COST 1
	AS $$
DECLARE
	_inway_org_id integer;
	_service_org_id integer;
BEGIN
	SELECT organization_id INTO _inway_org_id
		FROM directory.inways
		WHERE id = NEW.inway_id;
	SELECT organization_id INTO _service_org_id
		FROM directory.services
		WHERE id = NEW.service_id;
	IF _inway_org_id != _service_org_id THEN
		RAISE EXCEPTION 'service organization does not match inway organization';
	END IF;
	RETURN NEW;
END
$$;
-- ddl-end --
ALTER FUNCTION directory.availabilities_verify() OWNER TO postgres;
-- ddl-end --

-- object: availabilities_verify | type: TRIGGER --
-- availabilities_verify ON directory.availabilities CASCADE;
CREATE CONSTRAINT TRIGGER availabilities_verify
	AFTER INSERT OR UPDATE
	ON directory.availabilities
	NOT DEFERRABLE 
	FOR EACH ROW
	EXECUTE PROCEDURE directory.availabilities_verify();
-- ddl-end --

-- object: directory.availabilities_notify_event | type: FUNCTION --
-- DROP FUNCTION IF EXISTS directory.availabilities_notify_event() CASCADE;
CREATE FUNCTION directory.availabilities_notify_event ()
	RETURNS trigger
	LANGUAGE plpgsql
	VOLATILE 
	CALLED ON NULL INPUT
	SECURITY INVOKER
	COST 1
	AS $$
DECLARE 
	data json;	
	notification json;
BEGIN
	IF (TG_OP = 'DELETE') THEN
		data = row_to_json(OLD);
	ELSIF TG_OP = 'INSERT' OR (TG_OP = 'UPDATE' AND OLD.active != NEW.active) THEN
		SELECT row_to_json(t) into data from (
			SELECT
				availabilities.id,
				organizations.name AS organization_name,
				services.name AS service_name,
				inways.address,
				availabilities.active
			FROM directory.availabilities
				INNER JOIN directory.inways
					ON availabilities.inway_id = inways.id
				INNER JOIN directory.services
					ON availabilities.service_id = services.id
				INNER JOIN directory.organizations
					ON services.organization_id = organizations.id
			WHERE 
				availabilities.id = NEW.id
		 ) t;
	END IF;
		
	IF (TG_OP = 'UPDATE' AND OLD.active = NEW.active) = false THEN 
		notification = json_build_object(
					 	 'action', TG_OP,
					  	'availability', data);
			
		PERFORM pg_notify('availabilities',notification::text);
 	END IF;
	
	RETURN NULL;
END;
$$;
-- ddl-end --
ALTER FUNCTION directory.availabilities_notify_event() OWNER TO postgres;
-- ddl-end --

-- object: availabilities_notify | type: TRIGGER --
-- DROP TRIGGER IF EXISTS availabilities_notify ON directory.availabilities CASCADE;
CREATE TRIGGER availabilities_notify
	AFTER INSERT OR DELETE OR UPDATE
	ON directory.availabilities
	FOR EACH ROW
	EXECUTE PROCEDURE directory.availabilities_notify_event();
-- ddl-end --

-- object: inways_organization_id | type: INDEX --
-- DROP INDEX IF EXISTS directory.inways_organization_id CASCADE;
CREATE INDEX inways_organization_id ON directory.inways
	USING btree
	(
	  organization_id
	);
-- ddl-end --

-- object: services_organization_id | type: INDEX --
-- DROP INDEX IF EXISTS directory.services_organization_id CASCADE;
CREATE INDEX services_organization_id ON directory.services
	USING btree
	(
	  organization_id
	);
-- ddl-end --

-- object: availabilities_service_id | type: INDEX --
-- DROP INDEX IF EXISTS directory.availabilities_service_id CASCADE;
CREATE INDEX availabilities_service_id ON directory.availabilities
	USING btree
	(
	  service_id
	);
-- ddl-end --

-- object: availabilities_inway_id | type: INDEX --
-- DROP INDEX IF EXISTS directory.availabilities_inway_id CASCADE;
CREATE INDEX availabilities_inway_id ON directory.availabilities
	USING btree
	(
	  inway_id
	);
-- ddl-end --

-- object: directory.outways | type: TABLE --
-- DROP TABLE IF EXISTS directory.outways CASCADE;
CREATE TABLE directory.outways(
	id serial NOT NULL,
	announced timestamptz NOT NULL DEFAULT NOW(),
	version varchar(100) NOT NULL,
	CONSTRAINT outways_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE directory.outways OWNER TO postgres;
-- ddl-end --

-- Appended SQL commands --
GRANT USAGE, SELECT ON SEQUENCE outways_id_seq TO "nlx-directory";
-- ddl-end --

-- object: outways_announced | type: INDEX --
-- DROP INDEX IF EXISTS directory.outways_announced CASCADE;
CREATE INDEX outways_announced ON directory.outways
	USING btree
	(
	  announced
	);
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

-- object: grant_1ceb2a0d06 | type: PERMISSION --
GRANT SELECT,INSERT,UPDATE,DELETE
   ON TABLE directory.organizations
   TO "nlx-directory";
-- ddl-end --

-- object: grant_3c879d5b54 | type: PERMISSION --
GRANT SELECT,INSERT,UPDATE,DELETE
   ON TABLE directory.inways
   TO "nlx-directory";
-- ddl-end --

-- object: grant_1d00258e74 | type: PERMISSION --
GRANT SELECT,INSERT,UPDATE,DELETE
   ON TABLE directory.services
   TO "nlx-directory";
-- ddl-end --

-- object: grant_baaa0e26cb | type: PERMISSION --
GRANT SELECT,INSERT,UPDATE,DELETE
   ON TABLE directory.availabilities
   TO "nlx-directory";
-- ddl-end --

-- object: grant_0f13af9b64 | type: PERMISSION --
GRANT USAGE
   ON SCHEMA directory
   TO "nlx-directory";
-- ddl-end --

-- object: grant_15e1d02b37 | type: PERMISSION --
GRANT SELECT,INSERT,UPDATE,DELETE
   ON TABLE directory.outways
   TO "nlx-directory";
-- ddl-end --


