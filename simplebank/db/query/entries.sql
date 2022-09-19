-- name: CreateEntry :one
INSERT INTO entries (
    account_id,
    amount
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetEntriesByAccountId :many
SELECT * FROM entries 
WHERE account_id = $1 LIMIT $2; 

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: ListEntriesByAccountId :many 
SELECT * FROM entries 
WHERE account_id = $1
ORDER BY id
LIMIT $2 
OFFSET $3; 

-- name: UpdateEntriesAmount :one
UPDATE entries
set amount = $2
WHERE id = $1
RETURNING *; 

-- name: DeleteEntries :exec 
DELETE FROM entries WHERE id =$1;