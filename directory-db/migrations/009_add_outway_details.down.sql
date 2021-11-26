BEGIN transaction

DROP TABLE directory.outways;

CREATE TABLE directory.outways(
	id serial NOT NULL,
    announced timestamp with time zone DEFAULT now() NOT NULL,
    version character varying(100) NOT NULL DEFAULT 'unknown',
	CONSTRAINT outways_pk PRIMARY KEY (id)
);

CREATE INDEX outways_announced ON directory.outways USING btree (announced);

COMMIT;