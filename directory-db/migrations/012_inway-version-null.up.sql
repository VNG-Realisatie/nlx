UPDATE directory.inways SET version = 'unknown' WHERE version IS NULL;

ALTER TABLE
    directory.inways
ALTER COLUMN version SET DEFAULT 'unknown';
 
ALTER TABLE
    directory.inways
ALTER COLUMN version SET NOT NULL;

ALTER TABLE
    directory.outways
ALTER COLUMN version SET DEFAULT 'unknown'
