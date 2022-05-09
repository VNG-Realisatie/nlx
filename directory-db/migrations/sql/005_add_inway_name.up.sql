BEGIN transaction;

-- delete oldest inways of an organization if there are multiple
delete from directory.availabilities where inway_id not in (select max(id) from directory.inways group by directory.inways.organization_id);
delete from directory.inways where id not in (select max(id) from directory.inways group by directory.inways.organization_id);

-- now add the name column. only a single inway of an organization can have an empty name, therefore we had to
-- remove duplicates in the query above
alter table directory.inways add column name character varying(100) not null default '';
create unique index inways_name_organization_id ON directory.inways (name, organization_id);

COMMIT;
