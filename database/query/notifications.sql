-- name: CreateNotification :one
INSERT INTO notifications (event_type, flight_number, created_at)
VALUES ($1, $2, NOW())
RETURNING id;
