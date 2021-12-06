begin transaction;

alter table directory.organizations add column created_at timestamp with time zone default now() not null;

commit;
