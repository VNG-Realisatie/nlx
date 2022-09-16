-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

begin transaction;

alter table nlx_management.settings add column organization_email_address character varying(250);

commit;
