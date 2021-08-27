BEGIN transaction;

alter table directory.inways add column created_at timestamp with time zone DEFAULT now() NOT NULL;
alter table directory.inways add column updated_at timestamp with time zone DEFAULT now() NOT NULL;

COMMIT;
