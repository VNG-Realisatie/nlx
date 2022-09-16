-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

ALTER TABLE transactionlog.records ADD COLUMN delegator VARCHAR(250) NOT NULL DEFAULT '';
ALTER TABLE transactionlog.records ADD COLUMN order_reference VARCHAR(100) NOT NULL DEFAULT '';

COMMIT;
