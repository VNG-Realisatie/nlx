
WITH org AS (
    INSERT INTO directory.organizations (name)
        VALUES ('demo-organization')
        RETURNING id
)
INSERT INTO directory.services (organization_id, name)
    SELECT org.id, 'demo-api'
    FROM org;
