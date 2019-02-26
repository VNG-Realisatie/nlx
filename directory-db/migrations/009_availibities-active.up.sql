DROP TRIGGER availabilities_notify ON directory.availabilities;

ALTER TABLE directory.availabilities
	ADD COLUMN unhealthy_since timestamp with time zone,
	ADD COLUMN last_announced timestamp with time zone DEFAULT now() NOT NULL,
	ADD COLUMN active boolean DEFAULT false NOT NULL;

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

CREATE INDEX availabilities_inway_id ON directory.availabilities USING btree (inway_id);

CREATE INDEX availabilities_service_id ON directory.availabilities USING btree (service_id);

CREATE INDEX inways_organization_id ON directory.inways USING btree (organization_id);

CREATE INDEX services_organization_id ON directory.services USING btree (organization_id);

CREATE TRIGGER availabilities_notify
	AFTER INSERT OR UPDATE OR DELETE ON directory.availabilities
	FOR EACH ROW
	EXECUTE PROCEDURE directory.availabilities_notify_event();
