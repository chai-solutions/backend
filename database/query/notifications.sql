-- name: CreateNotification :one
INSERT INTO notifications (event_type, flight_number)
VALUES ($1, $2)
RETURNING id;
