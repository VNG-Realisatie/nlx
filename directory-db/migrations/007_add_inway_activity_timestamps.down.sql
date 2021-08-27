BEGIN transaction;

alter table directory.inways drop column created_at;
alter table directory.inways drop column updated_at;

COMMIT;
