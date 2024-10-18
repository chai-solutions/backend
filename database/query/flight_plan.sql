-- name: CreateFlightPlan :one
WITH new_flight_plan AS (
    INSERT INTO flight_plans (users)
    VALUES (@users)
    RETURNING id
)
INSERT INTO flight_plan_flights (flight_plan, flight)
SELECT new_flight_plan.id, f.id
FROM flights AS f, new_flight_plan
WHERE f.flight_number = @flight_number
RETURNING flight_plan_flights.*
;

-- name: PatchFlightPlan :one
INSERT INTO flight_plan_flights (flight_plan, flight)
SELECT fp.id, f.id
FROM flight_plans AS fp
JOIN flights AS f ON f.flight_number = @flight_number
WHERE fp.id = @flight_plan
RETURNING flight_plan_flights.*
;

-- name: GetFlightPlans :many
SELECT 
    flight_plans.id AS flight_plan_id,
    flights.flight_number,
    dep_airport.name AS dep_airport,
    arr_airport.name AS arr_airport,
    flights.sched_dep_time,
    flights.sched_arr_time,
    flights.actual_dep_time,
    flights.actual_arr_time
FROM flight_plans
JOIN flight_plan_flights 
ON flight_plans.id = flight_plan_flights.flight_plan
JOIN flights 
ON flight_plan_flights.flight = flights.id
JOIN airports AS dep_airport 
ON flights.dep_airport = dep_airport.id
JOIN airports AS arr_airport 
ON flights.arr_airport = arr_airport.id
WHERE flight_plans.users = @users
;

-- name: GetFlightPlan :many
SELECT 
    fp.id,
    f.flight_number,
    departure_airport.name AS dep_airport_name,
    arrival_airport.name AS arr_airport_airport,
    f.sched_dep_time,
    f.sched_arr_time,
    f.actual_dep_time,
    f.actual_arr_time
FROM flight_plans AS fp
JOIN flight_plan_flights AS fpf
ON fp.id = fpf.flight_plan
JOIN flights AS f
ON fpf.flight = f.id
JOIN airports AS departure_airport
ON f.dep_airport = departure_airport.id
JOIN airports AS arrival_airport
ON f.arr_airport = arrival_airport.id
WHERE fp.users = @users AND fp.id = @id
;

-- name: DeleteFlightPlan :many
DELETE FROM flight_plans
WHERE id = @id
RETURNING *;

-- name: DeleteFlightPlanStep :one
DELETE FROM flight_plan_flights
WHERE id = @id
RETURNING *;
