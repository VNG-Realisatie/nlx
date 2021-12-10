-- name: ListRecords :many
SELECT id,
    direction,
    created,
    src_organization,
    dest_organization,
    service_name,
    logrecord_id,
    data,
    delegator,
    order_reference
FROM transactionlog.records
ORDER BY created DESC
LIMIT $1;

-- name: CreateRecord :exec
INSERT INTO transactionlog.records (
    direction,
    created,
    src_organization,
    dest_organization,
    service_name,
    logrecord_id,
    data,
    delegator,
    order_reference
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
