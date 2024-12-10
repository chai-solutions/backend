-- name: CreateNotifications :copyfrom
INSERT INTO notifications ("user", title, message)
VALUES ($1, $2, $3);
