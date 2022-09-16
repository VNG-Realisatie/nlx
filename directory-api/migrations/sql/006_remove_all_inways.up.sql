-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

begin transaction;

delete from directory.availabilities;
delete from directory.inways;

commit;
