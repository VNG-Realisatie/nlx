BEGIN transaction;

CREATE SCHEMA transactionlog;

CREATE TYPE transactionlog.direction AS ENUM (
    'in',
    'out'
);

CREATE TABLE transactionlog.records (
    id serial NOT NULL,
    direction transactionlog.direction NOT NULL,
    created timestamp with time zone DEFAULT now() NOT NULL,
    src_organization text NOT NULL,
    dest_organization text NOT NULL,
    service_name text NOT NULL,
    logrecord_id text NOT NULL,
    "data" jsonb,
    CONSTRAINT records_pk PRIMARY KEY (id)
);

COMMIT;
