BEGIN transaction;

ALTER TABLE directory.services ADD COLUMN one_time_costs INT NOT NULL DEFAULT 0;
ALTER TABLE directory.services ADD COLUMN monthly_costs INT NOT NULL DEFAULT 0;
ALTER TABLE directory.services ADD COLUMN request_costs INT NOT NULL DEFAULT 0;

COMMIT;
