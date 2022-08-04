BEGIN transaction;

ALTER TABLE
    nlx_management.settings
DROP COLUMN created_at,
DROP COLUMN updated_at;

ALTER TABLE
    nlx_management.settings
RENAME CONSTRAINT
    "fk_organization_settings_inway"
TO
    "fk_settings_inway_id_inways_id";

--- the lines below will ensure that there is a single row with settings available for the organization
INSERT INTO nlx_management.settings (inway_id, organization_email_address) VALUES (NULL, '');
DELETE FROM nlx_management.settings WHERE id != (SELECT MIN(id) FROM nlx_management.settings);

COMMIT;
