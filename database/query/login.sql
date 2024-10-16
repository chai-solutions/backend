-- name: SelectAccountByEmail :many
SELECT *
FROM users
WHERE email = @email;

-- name: CreateSession :one
INSERT INTO sessions (user_id, token, expires_at)
VALUES (@user_id, @token, @expiry_time)
RETURNING user_id, token;
