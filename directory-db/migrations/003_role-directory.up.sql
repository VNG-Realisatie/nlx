
DO
$do$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname='nlx-directory') THEN
		CREATE ROLE "nlx-directory";
	END IF;
END
$do$;

ALTER ROLE "nlx-directory" WITH INHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS ENCRYPTED PASSWORD 'nlx-directory';

REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA directory FROM "nlx-directory";
GRANT SELECT ON TABLE public.schema_migrations TO "nlx-directory";
GRANT USAGE ON SCHEMA directory TO "nlx-directory";
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA directory TO "nlx-directory";
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA directory TO "nlx-directory";
