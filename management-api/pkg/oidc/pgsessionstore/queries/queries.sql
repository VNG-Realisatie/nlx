-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

-- name: CreateSession :exec
INSERT INTO http_sessions.sessions (
    key, 
    data, 
    created_at, 
    updated_at, 
    expires_on
)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateSession :exec
UPDATE http_sessions.sessions 
SET data=$1, 
    updated_at=$2, 
    expires_on=$3 
WHERE key=$4;

-- name: GetSession :one
SELECT id, 
    key, 
    data, 
    created_at, 
    updated_at, 
    expires_on 
FROM http_sessions.sessions 
WHERE key = $1;

-- name: DeleteSession :exec
DELETE
FROM http_sessions.sessions
WHERE key = $1;

-- name: DeleteExpiredSessions :exec
DELETE
FROM http_sessions.sessions 
WHERE expires_on < now();

-- name: CountExpiredSessions :one
SELECT count(*) 
FROM http_sessions.sessions
WHERE expires_on < now();
