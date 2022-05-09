begin transaction;

alter table directory.inways add column created_at timestamp with time zone default now() not null;
alter table directory.inways add column updated_at timestamp with time zone default now() not null;

commit;
