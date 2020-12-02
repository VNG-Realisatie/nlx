BEGIN transaction;

ALTER TABLE nlx_management.access_requests_incoming DROP CONSTRAINT fk_access_requests_incoming_service;
ALTER TABLE nlx_management.access_grants DROP CONSTRAINT fk_access_grants_incoming_access_request;
ALTER TABLE nlx_management.access_proofs DROP CONSTRAINT fk_access_proofs_outgoing_access_request;
ALTER TABLE nlx_management.inways_services DROP CONSTRAINT fk_inway_services_inway;
ALTER TABLE nlx_management.inways_services DROP CONSTRAINT fk_inway_services_service;
ALTER TABLE nlx_management.settings DROP CONSTRAINT fk_organization_settings_inway;

DROP TABLE nlx_management.services;
DROP TABLE nlx_management.access_requests_incoming;
DROP TABLE nlx_management.access_requests_outgoing;
DROP TABLE nlx_management.access_grants;
DROP TABLE nlx_management.access_proofs;
DROP TABLE nlx_management.inways;
DROP TABLE nlx_management.inways_services;
DROP TABLE nlx_management.settings;

DROP SCHEMA nlx_management;

COMMIT;
