-- name: ListRecords :many
SELECT id,
    direction,
    created,
    src_organization,
    dest_organization,
    service_name,
    transaction_id,
    data,
    delegator,
    order_reference
FROM transactionlog.records
ORDER BY created DESC
LIMIT $1;

-- name: CreateRecord :one
INSERT INTO transactionlog.records (
    direction,
    src_organization,
    dest_organization,
    service_name,
    transaction_id,
    data,
    delegator,
    order_reference,
    created
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
returning id;

-- name: CreateDataSubject :exec
INSERT INTO transactionlog.datasubjects (
    record_id,
    key,
    value
)
VALUES ($1, $2, $3);
