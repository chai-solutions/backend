// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: login.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING user_id, token
`

type CreateSessionParams struct {
	UserID     pgtype.Int4      `json:"user_id"`
	Token      string           `json:"token"`
	ExpiryTime pgtype.Timestamp `json:"expiry_time"`
}

type CreateSessionRow struct {
	UserID pgtype.Int4 `json:"user_id"`
	Token  string      `json:"token"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (CreateSessionRow, error) {
	row := q.db.QueryRow(ctx, createSession, arg.UserID, arg.Token, arg.ExpiryTime)
	var i CreateSessionRow
	err := row.Scan(&i.UserID, &i.Token)
	return i, err
}

const selectAccountByEmail = `-- name: SelectAccountByEmail :many
SELECT id, created_at, name, email, password
FROM users
WHERE email = $1
`

func (q *Queries) SelectAccountByEmail(ctx context.Context, email string) ([]User, error) {
	rows, err := q.db.Query(ctx, selectAccountByEmail, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Email,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
