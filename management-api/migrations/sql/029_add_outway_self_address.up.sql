BEGIN transaction;

ALTER TABLE nlx_management.outways ADD COLUMN self_address VARCHAR(100) NOT NULL DEFAULT '';

COMMIT;
