begin transaction;

drop index directory."inways_name_organization_id";
alter table directory.inways drop column name;

commit;
