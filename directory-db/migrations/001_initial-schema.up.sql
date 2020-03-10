BEGIN transaction;

CREATE SCHEMA directory;

CREATE TABLE directory.availabilities(
	id serial NOT NULL,
	inway_id integer NOT NULL,
	service_id integer NOT NULL,
	healthy bool NOT NULL DEFAULT false,
	unhealthy_since timestamp with time zone,
	last_announced timestamp with time zone DEFAULT now() NOT NULL,
	active boolean DEFAULT false NOT NULL,
	CONSTRAINT availabilities_pk PRIMARY KEY (id),
	CONSTRAINT availabilities_uq_inway_service UNIQUE(inway_id,service_id)
);

CREATE INDEX availabilities_inway_id ON directory.availabilities USING btree (inway_id);
CREATE INDEX availabilities_service_id ON directory.availabilities USING btree (service_id);

CREATE TABLE directory.organizations(
	id serial NOT NULL,
	name character varying(100) NOT NULL,
	insight_log_endpoint character varying(250),
	insight_irma_endpoint character varying(250),
	CONSTRAINT organizations_pk PRIMARY KEY (id),
	CONSTRAINT organizations_uq_name UNIQUE (name)
);

CREATE TABLE directory.services(
	id serial NOT NULL,
	organization_id integer NOT NULL,
	name character varying(250) NOT NULL,
	documentation_url character varying(250),
	api_specification_type character varying(20),
	internal boolean DEFAULT false NOT NULL,
	tech_support_contact character varying(250),
	public_support_contact character varying(250),
	CONSTRAINT services_pk PRIMARY KEY (id),
	CONSTRAINT services_uq_name UNIQUE (organization_id, name),
	CONSTRAINT services_check_typespec CHECK (
        (api_specification_type IS NULL) OR (
            (api_specification_type::text = 'OpenAPI2'::text)
            OR
            (api_specification_type::text = 'OpenAPI3'::text)
        )
	)
);

CREATE INDEX services_organization_id ON directory.services USING btree (organization_id);

CREATE TABLE directory.inways(
	id serial NOT NULL,
	organization_id integer NOT NULL,
	address character varying(100) NOT NULL,
	version character varying(100) NOT NULL DEFAULT 'unknown',
	CONSTRAINT inways_pk PRIMARY KEY (id),
	CONSTRAINT inways_uq_address UNIQUE (organization_id,address)
);

CREATE INDEX inways_organization_id ON directory.inways USING btree (organization_id);

CREATE TABLE directory.outways(
	id serial NOT NULL,
    announced timestamp with time zone DEFAULT now() NOT NULL,
    version character varying(100) NOT NULL DEFAULT 'unknown',
	CONSTRAINT outways_pk PRIMARY KEY (id)
);

CREATE INDEX outways_announced ON directory.outways USING btree (announced);

ALTER TABLE directory.services ADD CONSTRAINT services_fk_organization FOREIGN KEY (organization_id) REFERENCES directory.organizations (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE directory.inways ADD CONSTRAINT inways_fk_organization FOREIGN KEY (organization_id) REFERENCES directory.organizations (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE directory.availabilities ADD CONSTRAINT availabilities_fk_inway FOREIGN KEY (inway_id) REFERENCES directory.inways (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE directory.availabilities ADD CONSTRAINT availabilities_fk_service FOREIGN KEY (service_id) REFERENCES directory.services (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;

CREATE OR REPLACE FUNCTION directory.availabilities_verify() RETURNS trigger
    LANGUAGE plpgsql COST 1
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

CREATE OR REPLACE FUNCTION directory.availabilities_notify_event() RETURNS trigger
	LANGUAGE plpgsql COST 1
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

CREATE TRIGGER availabilities_notify AFTER INSERT OR UPDATE OR DELETE ON directory.availabilities FOR EACH ROW EXECUTE PROCEDURE directory.availabilities_notify_event();

COMMIT;
