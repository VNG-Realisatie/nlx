BEGIN transaction;

ALTER TABLE nlx_management.services DROP COLUMN one_time_costs;
ALTER TABLE nlx_management.services DROP COLUMN monthly_costs;
ALTER TABLE nlx_management.services DROP COLUMN request_costs;

COMMIT;
