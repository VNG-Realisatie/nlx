-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

CREATE TABLE nlx_management.outways (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL UNIQUE,
  public_key_pem VARCHAR(4096) NOT NULL,
  version VARCHAR(100) NOT NULL,
  ip_address INET NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL);

COMMIT;
