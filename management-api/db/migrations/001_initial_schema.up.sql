BEGIN transaction; 

CREATE SCHEMA nlx_management;

CREATE TABLE nlx_management.services (
  service_id SERIAL PRIMARY KEY,
  name VARCHAR(250) NOT NULL UNIQUE,
  endpoint_url VARCHAR(250) NOT NULL,
  documentation_url VARCHAR(250) NOT NULL,
  api_specification_url VARCHAR(250) NOT NULL,
  internal boolean NOT NULL DEFAULT false,
  tech_support_contact VARCHAR(250) NOT NULL,
  public_support_contact VARCHAR(250) NOT NULL,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL);


CREATE TABLE nlx_management.access_requests_incoming (
  access_request_incoming_id SERIAL PRIMARY KEY,
  service_id INT NOT NULL,
  organization_name VARCHAR(250) NOT NULL,
  public_key_fingerprint VARCHAR(44) NOT NULL,
  state VARCHAR(50) CHECK (state IN ('received', 'approved', 'rejected', 'revoked')) NOT NULL DEFAULT 'received',
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL,
  CONSTRAINT fk_access_requests_incoming_service
    FOREIGN KEY (service_id)
    REFERENCES nlx_management.services (service_id) 
    ON DELETE RESTRICT
);

CREATE INDEX fk_access_requests_incoming_service ON nlx_management.access_requests_incoming (service_id);


CREATE TABLE nlx_management.access_grants (
  access_grant_id SERIAL PRIMARY KEY,
  access_request_incoming_id INT NOT NULL,
  created_at timestamp with time zone NOT NULL,
  revoked_at timestamp with time zone NULL DEFAULT NULL,
  CONSTRAINT fk_access_grants_incoming_access_request
    FOREIGN KEY (access_request_incoming_id)
    REFERENCES nlx_management.access_requests_incoming (access_request_incoming_id)
    ON DELETE RESTRICT
);

CREATE INDEX fk_access_grants_incoming_access_request ON nlx_management.access_grants (access_request_incoming_id);


CREATE TABLE nlx_management.access_requests_outgoing (
  access_request_outgoing_id SERIAL PRIMARY KEY,
  organization_name VARCHAR(100) NOT NULL,
  service_name VARCHAR(250) NOT NULL,
  state VARCHAR(50) CHECK (state IN ('created', 'received', 'approved', 'rejected', 'revoked')) NOT NULL DEFAULT 'created',
  public_key_fingerprint VARCHAR(44) NOT NULL,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL
  );

CREATE INDEX idx_organization_service ON nlx_management.access_requests_outgoing (organization_name, service_name);


CREATE TABLE nlx_management.access_proofs (
  access_proof_id SERIAL PRIMARY KEY,
  access_request_outgoing_id INT  NOT NULL,
  created_at timestamp with time zone NOT NULL,
  revoked_at timestamp with time zone NULL DEFAULT NULL,
  CONSTRAINT fk_access_proofs_outgoing_access_request
    FOREIGN KEY (access_request_outgoing_id)
    REFERENCES nlx_management.access_requests_outgoing (access_request_outgoing_id)
    ON DELETE RESTRICT
);

CREATE INDEX fk_access_proofs_outgoing_access_request ON nlx_management.access_proofs (access_request_outgoing_id);


CREATE TABLE nlx_management.inways ( 
  inway_id SERIAL PRIMARY KEY,
  name VARCHAR(250) NOT NULL UNIQUE,
  self_address VARCHAR(100) NOT NULL,
  version VARCHAR(100) NOT NULL,
  hostname VARCHAR(250) NOT NULL,
  ip_address VARCHAR(45) NOT NULL,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL);


CREATE TABLE nlx_management.inways_services (
  inway_id SERIAL NOT NULL,
  service_id SERIAL NOT NULL,
  PRIMARY KEY (inway_id, service_id),
  CONSTRAINT fk_inway_services_inway
    FOREIGN KEY (inway_id)
    REFERENCES nlx_management.inways (inway_id)
    ON DELETE RESTRICT,
  CONSTRAINT fk_inway_services_service
    FOREIGN KEY (service_id)
    REFERENCES nlx_management.services (service_id)
    ON DELETE RESTRICT
);


CREATE TABLE nlx_management.settings (
  inway_id INT NULL DEFAULT NULL,
  insight_api_url VARCHAR(250) NOT NULL,
  irma_server_url VARCHAR(250) NOT NULL,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL,
  CONSTRAINT fk_organization_settings_inway
    FOREIGN KEY (inway_id)
    REFERENCES nlx_management.inways (inway_id)
);

CREATE INDEX fk_organization_settings_inway ON nlx_management.settings (inway_id);


COMMIT;
