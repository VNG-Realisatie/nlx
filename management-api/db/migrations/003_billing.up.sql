BEGIN transaction;

ALTER TABLE nlx_management.services ADD COLUMN one_time_costs INT NOT NULL DEFAULT 0;
ALTER TABLE nlx_management.services ADD COLUMN monthly_costs INT NOT NULL DEFAULT 0;
ALTER TABLE nlx_management.services ADD COLUMN request_costs INT NOT NULL DEFAULT 0;

COMMIT;
