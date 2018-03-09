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


CREATE CONSTRAINT TRIGGER availabilities_verify
	AFTER INSERT OR UPDATE
	ON directory.availabilities
	NOT DEFERRABLE INITIALLY IMMEDIATE
	FOR EACH ROW
	EXECUTE PROCEDURE directory.availabilities_verify();
