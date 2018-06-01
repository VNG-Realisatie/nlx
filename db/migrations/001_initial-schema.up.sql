
CREATE SCHEMA directory;
ALTER SCHEMA directory OWNER TO postgres;

CREATE TABLE directory.organizations(
	id serial NOT NULL,
	name varchar(100) NOT NULL,
	CONSTRAINT organizations_pk PRIMARY KEY (id),
	CONSTRAINT organizations_uq_name UNIQUE (name)

);
ALTER TABLE directory.organizations OWNER TO postgres;

CREATE TABLE directory.services(
	id serial NOT NULL,
	organization_id integer NOT NULL,
	name varchar(100) NOT NULL,
	documentation_url character varying(250),
	CONSTRAINT services_pk PRIMARY KEY (id),
	CONSTRAINT services_uq_name UNIQUE (organization_id,name)

);
ALTER TABLE directory.services OWNER TO postgres;

CREATE TABLE directory.inways(
	id serial NOT NULL,
	organization_id integer NOT NULL,
	address varchar(100) NOT NULL,
	CONSTRAINT inways_pk PRIMARY KEY (id),
	CONSTRAINT inways_uq_address UNIQUE (organization_id,address)

);
ALTER TABLE directory.inways OWNER TO postgres;

CREATE TABLE directory.availabilities(
	id serial NOT NULL,
	inway_id integer NOT NULL,
	service_id integer NOT NULL,
	healthy bool NOT NULL DEFAULT false,
	CONSTRAINT availabilities_pk PRIMARY KEY (id),
	CONSTRAINT availabilities_uq_inway_service UNIQUE(inway_id,service_id)

);
ALTER TABLE directory.availabilities OWNER TO postgres;

ALTER TABLE directory.services ADD CONSTRAINT services_fk_organization FOREIGN KEY (organization_id)
REFERENCES directory.organizations (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE directory.inways ADD CONSTRAINT inways_fk_organization FOREIGN KEY (organization_id)
REFERENCES directory.organizations (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE directory.availabilities ADD CONSTRAINT availabilities_fk_inway FOREIGN KEY (inway_id)
REFERENCES directory.inways (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE directory.availabilities ADD CONSTRAINT availabilities_fk_service FOREIGN KEY (service_id)
REFERENCES directory.services (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
