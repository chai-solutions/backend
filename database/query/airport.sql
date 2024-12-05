-- name: GetAllAirports :many
SELECT *
FROM airports;

-- name: GetAirportByIATACode :one
SELECT *
FROM airports AS a
WHERE a.iata = @iata;

-- name: GetAirportByID :one
SELECT *
FROM airports AS a
WHERE a.id = @id;
