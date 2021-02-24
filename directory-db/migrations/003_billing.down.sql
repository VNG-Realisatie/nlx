BEGIN transaction;

ALTER TABLE directory.services DROP COLUMN one_time_costs;
ALTER TABLE directory.services DROP COLUMN monthly_costs;
ALTER TABLE directory.services DROP COLUMN request_costs;

COMMIT;
