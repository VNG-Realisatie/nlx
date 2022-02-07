BEGIN transaction;

ALTER TABLE nlx_management.outways DROP COLUMN public_key_fingerprint VARCHAR(44) NOT NULL;

COMMIT;
