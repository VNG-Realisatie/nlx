-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

CREATE SCHEMA http_sessions;

CREATE TABLE http_sessions.sessions (
    id BIGSERIAL PRIMARY KEY,
    key BYTEA,
    data BYTEA,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ,
    expires_on TIMESTAMPTZ
);

CREATE INDEX http_sessions_sessions_expiry_idx ON http_sessions.sessions (expires_on);
CREATE INDEX http_sessions_sessions_key_idx ON http_sessions.sessions (key);

COMMIT;
