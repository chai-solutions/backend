-- name: GetAllAirports :many
SELECT *
FROM airports;

-- name: GetAirport :one
SELECT a.id, a.iata, a.name, a.latitude, a.longitude
FROM airports AS a
WHERE a.iata = @iata
;
