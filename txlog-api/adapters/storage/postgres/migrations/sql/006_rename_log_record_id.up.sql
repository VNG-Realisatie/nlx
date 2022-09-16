-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

ALTER TABLE transactionlog.records
RENAME COLUMN logrecord_id TO transaction_id;

COMMIT;
