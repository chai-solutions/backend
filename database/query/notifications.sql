-- name: CreateNotifications :copyfrom
INSERT INTO notifications ("user", title, message)
VALUES ($1, $2, $3);

-- name: GetUsersNotifications :many
SELECT * FROM notifications
WHERE "user" = $1;
