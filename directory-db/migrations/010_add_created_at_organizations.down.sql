begin transaction;

alter table directory.organizations drop column created_at;

commit;
