begin transaction;

alter table directory.organizations add column email_address character varying(250);

commit;
