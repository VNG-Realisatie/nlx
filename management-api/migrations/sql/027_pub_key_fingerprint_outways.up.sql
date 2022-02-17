BEGIN transaction;

ALTER TABLE nlx_management.outways ADD COLUMN public_key_fingerprint VARCHAR(44);

COMMIT;
