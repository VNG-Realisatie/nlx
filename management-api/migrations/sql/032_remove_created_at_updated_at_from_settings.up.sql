BEGIN transaction;

ALTER TABLE
    nlx_management.settings
DROP COLUMN created_at,
DROP COLUMN updated_at;

COMMIT;
