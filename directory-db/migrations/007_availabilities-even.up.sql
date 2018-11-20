CREATE OR REPLACE FUNCTION directory.availabilities_notify_event() RETURNS trigger
    LANGUAGE plpgsql COST 1
    AS $$
    DECLARE 
        data json;
        notification json;
    BEGIN
        IF (TG_OP = 'DELETE') THEN
            data = row_to_json(OLD);
        ELSE
             SELECT row_to_json(t) into data from (
                SELECT
                    availabilities.id,
                    organizations.name AS organization_name,
                    services.name AS service_name,
                    inways.address
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

      
        notification = json_build_object(
                          'action', TG_OP,
                          'availability', data);
                
        PERFORM pg_notify('availabilities',notification::text);
   
        RETURN NULL; 
    END;
$$;

CREATE TRIGGER availabilities_notify
	AFTER INSERT OR DELETE ON directory.availabilities
	FOR EACH ROW
	EXECUTE PROCEDURE directory.availabilities_notify_event();
