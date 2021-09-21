begin transaction;

alter table directory.organizations drop constraint organizations_uq_serial_number;
alter table directory.organizations drop column serial_number;

create unique index organizations_uq_name
    on directory.organizations (name);
alter table directory.organizations
    add constraint organizations_uq_name unique
    using index organizations_uq_name;

commit;
