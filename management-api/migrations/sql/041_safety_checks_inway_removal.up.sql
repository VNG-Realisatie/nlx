-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

ALTER TABLE nlx_management.settings DROP CONSTRAINT fk_settings_inway_id_inways_id;

ALTER TABLE nlx_management.settings ADD
    CONSTRAINT fk_settings_inway
        FOREIGN KEY (inway_id)
            REFERENCES nlx_management.inways (id)
            ON DELETE RESTRICT;

COMMIT;
