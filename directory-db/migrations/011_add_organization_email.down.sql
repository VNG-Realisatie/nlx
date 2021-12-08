begin transaction;

alter table directory.organizations drop column email_address;

commit;
