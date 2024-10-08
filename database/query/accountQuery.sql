-- name: InsertAccount :one
INSERT INTO accounts (owner, balance, currency)
VALUES ($1, $2, $3)
RETURNING id, owner, balance, currency, created_at;

-- name: SelectAccountByID :one
SELECT id, owner, balance, currency, created_at
FROM accounts
WHERE id = $1;
