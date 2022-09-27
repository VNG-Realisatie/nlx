-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

-- New unspecified error code is added starting at zero, so all existing error codes need to be incremented by 1
UPDATE nlx_management.access_requests_outgoing SET error_code = error_code + 1 WHERE state = 'failed' AND error_code IS NOT NULL;

COMMIT;
