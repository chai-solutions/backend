-- name: CreateFlightPlan :many
INSERT INTO flight_plans (users)
VALUES (@users)
RETURNING *;