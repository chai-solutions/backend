// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accountQuery.sql

package sqlc

import (
	"context"
)

const insertAccount = `-- name: InsertAccount :one
INSERT INTO accounts (owner, balance, currency)
VALUES ($1, $2, $3)
RETURNING id, owner, balance, currency, created_at
`

type InsertAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) InsertAccount(ctx context.Context, arg InsertAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, insertAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const selectAccountByID = `-- name: SelectAccountByID :one
SELECT id, owner, balance, currency, created_at
FROM accounts
WHERE id = $1
`

func (q *Queries) SelectAccountByID(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRow(ctx, selectAccountByID, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}