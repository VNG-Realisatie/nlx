ALTER TABLE directory.organizations ADD inway_id integer NULL;

ALTER TABLE directory.organizations ADD CONSTRAINT organizations_fk_inway FOREIGN KEY (inway_id) REFERENCES directory.inways(id) ON DELETE SET NULL;
