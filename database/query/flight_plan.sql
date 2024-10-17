-- name: CreateFlightPlan :many
INSERT INTO flight_plans (users)
VALUES (@users)
RETURNING *;

-- name: GetFlightPlans :many
SELECT 
    flight_plans.id AS flight_plan_id,
    flights.flight_number,
    dep_airport.name AS departure_airport,
    arr_airport.name AS arrival_airport,
    flights.sched_dep_time,
    flights.sched_arr_time,
    flights.actual_dep_time,
    flights.actual_arr_time
FROM flight_plans
JOIN flight_plan_flights ON flight_plans.id = flight_plan_flights.flight_plan
JOIN flights ON flight_plan_flights.flight = flights.id
JOIN airports AS dep_airport ON flights.dep_airport = dep_airport.id
JOIN airports AS arr_airport ON flights.arr_airport = arr_airport.id
WHERE flight_plans.users = @users
;

-- name: GetFlightPlan :one
SELECT 
    flight_plans.id AS flight_plan_id,
    flights.flight_number,
    dep_airport.name AS departure_airport,
    arr_airport.name AS arrival_airport,
    flights.sched_dep_time,
    flights.sched_arr_time,
    flights.actual_dep_time,
    flights.actual_arr_time
FROM flight_plans
JOIN flight_plan_flights ON flight_plans.id = flight_plan_flights.flight_plan
JOIN flights ON flight_plan_flights.flight = flights.id
JOIN airports AS dep_airport ON flights.dep_airport = dep_airport.id
JOIN airports AS arr_airport ON flights.arr_airport = arr_airport.id
WHERE flight_plans.users = @users AND flight_plans.id = @id
;