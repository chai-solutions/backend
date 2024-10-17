-- name: GetFlights :many
SELECT 
    f.id, f.flight_number, f.sched_dep_time, f.actual_arr_time, f.actual_dep_time,
    a.iata AS arrival_iata, a.name AS arrival_name,
    d.iata AS dep_iata, d.name AS dep_name
FROM flights AS f
INNER JOIN airports AS d
    ON d.id = f.dep_airport
INNER JOIN airports AS a
    ON a.id = f.arr_airport
WHERE d.iata = @dep
AND a.iata = @arr
;

-- name: GetFlight :one
SELECT 
    f.id, f.flight_number, f.sched_dep_time, f.actual_arr_time, f.actual_dep_time,
    a.iata AS arrival_iata, a.name AS arrival_name,
    d.iata AS dep_iata, d.name AS dep_name
FROM flights AS f
INNER JOIN airports AS d
    ON d.id = f.dep_airport
INNER JOIN airports AS a
    ON a.id = f.arr_airport
WHERE f.id = @id
;