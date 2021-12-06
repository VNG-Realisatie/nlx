begin transaction;

alter table nlx_management.settings drop column organization_email_address;

commit;
