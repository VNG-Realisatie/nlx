BEGIN transaction;

ALTER TABLE nlx_management.outways ADD COLUMN listen_address VARCHAR(255) NOT NULL DEFAULT '';

COMMIT;
