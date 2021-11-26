BEGIN transaction;

DROP TABLE directory.outways;

CREATE TABLE directory.outways(
	id serial NOT NULL,
	organization_id integer NOT NULL,
	name character varying(100) not null default '',
	version character varying(100) NOT NULL DEFAULT 'unknown',
	updated_at timestamp with time zone default now() not null,
	created_at timestamp with time zone default now() not null,
	CONSTRAINT outways_pk PRIMARY KEY (id),
	CONSTRAINT outways_uq_name UNIQUE (organization_id, name)
);

CREATE INDEX outways_organization_id ON directory.outways USING btree (organization_id);
ALTER TABLE directory.outways ADD CONSTRAINT outways_fk_organization FOREIGN KEY (organization_id) REFERENCES directory.organizations (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;

COMMIT;
