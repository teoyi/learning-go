-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
 ) VALUES (
    $1, $2, $3
 )
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTranfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: UpdateTransfers :one
UPDATE transfers
set amount = $2
WHERE id = $1
RETURNING *; 

-- name: DeleteTransfers :exec 
DELETE FROM transfers WHERE id =$1;