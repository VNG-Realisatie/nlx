ALTER TABLE directory.availabilities
	ADD COLUMN inway_version character varying(100);

ALTER TABLE directory.services
	ALTER COLUMN name TYPE character varying(250) USING name::character varying(250)
