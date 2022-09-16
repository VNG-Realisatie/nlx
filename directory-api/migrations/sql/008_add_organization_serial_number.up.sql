-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

begin transaction;

truncate table directory.organizations cascade;

alter table directory.organizations
    add column serial_number varchar(20) not null;

create unique index organizations_uq_serial_number
    on directory.organizations (serial_number);

alter table directory.organizations
    add constraint organizations_uq_serial_number unique
    using index organizations_uq_serial_number;

alter table directory.organizations
    drop constraint organizations_uq_name;

commit;
