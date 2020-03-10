BEGIN transaction;

DROP TRIGGER availabilities_notify ON directory.availabilities;
DROP FUNCTION directory.availabilities_verify();
DROP FUNCTION directory.availabilities_notify_event();

ALTER TABLE directory.services DROP CONSTRAINT services_fk_organization;
ALTER TABLE directory.inways DROP CONSTRAINT inways_fk_organization;
ALTER TABLE directory.availabilities DROP CONSTRAINT availabilities_fk_inway;
ALTER TABLE directory.availabilities DROP CONSTRAINT availabilities_fk_service;

DROP TABLE directory.availabilities;
DROP TABLE directory.organizations;
DROP TABLE directory.services;
DROP TABLE directory.inways;
DROP TABLE directory.outways;

DROP SCHEMA directory;

COMMIT;
