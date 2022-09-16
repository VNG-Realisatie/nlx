-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

begin transaction;

alter table directory.organizations add column created_at timestamp with time zone default now() not null;

commit;
