CREATE TABLE directory.outways(
	id serial NOT NULL,
    announced timestamp with time zone DEFAULT now() NOT NULL,
    version character varying(100),
	CONSTRAINT outways_pk PRIMARY KEY (id)
);
CREATE INDEX outways_announced ON directory.outways USING btree (announced);
GRANT SELECT, INSERT, UPDATE, DELETE ON directory.outways TO "nlx-directory";
GRANT USAGE, SELECT ON directory.outways_id_seq TO "nlx-directory";
