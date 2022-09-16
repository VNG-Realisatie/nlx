-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN TRANSACTION;

UPDATE nlx_management.audit_logs
SET data = jsonb_build_object('delegatee', "delegatee")
WHERE action_type = 'order_create';

ALTER TABLE nlx_management.audit_logs DROP COLUMN delegatee;

COMMIT;
