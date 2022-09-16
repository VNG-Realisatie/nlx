-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

begin transaction;

alter table directory.inways add column management_api_proxy_address character varying(255);

commit;
