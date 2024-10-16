-- name: CreateAccount :one
INSERT INTO users (name, email, password)
VALUES (@name, @email, @password_hash)
RETURNING name, email, created_at;
