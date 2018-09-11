
DO
$do$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname='nlx-org-txlog-api') THEN
		CREATE ROLE "nlx-org-txlog-api";
	END IF;
END
$do$;

ALTER ROLE "nlx-org-txlog-api" WITH INHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS ENCRYPTED PASSWORD 'nlx-txlog-api';

REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA public, transactionlog FROM "nlx-org-txlog-api";
GRANT SELECT ON TABLE public.schema_migrations TO "nlx-org-txlog-api";
GRANT USAGE ON SCHEMA transactionlog TO "nlx-org-txlog-api";
GRANT SELECT ON TABLE transactionlog.records TO "nlx-org-txlog-api";


DO
$do$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname='nlx-org-txlog-writer') THEN
		CREATE ROLE "nlx-org-txlog-writer";
	END IF;
END
$do$;

ALTER ROLE "nlx-org-txlog-writer" WITH INHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS ENCRYPTED PASSWORD 'nlx-txlog-writer';

REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA public, transactionlog FROM "nlx-org-txlog-writer";
GRANT SELECT ON TABLE public.schema_migrations TO "nlx-org-txlog-writer";
GRANT USAGE ON SCHEMA transactionlog TO "nlx-org-txlog-writer";
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA transactionlog TO "nlx-org-txlog-writer";
GRANT INSERT ON TABLE transactionlog.records TO "nlx-org-txlog-writer";
