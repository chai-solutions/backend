// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: flight_plan.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFlightPlan = `-- name: CreateFlightPlan :one
WITH new_flight_plan AS (
    INSERT INTO flight_plans (users)
    VALUES ($2)
    RETURNING id
)
INSERT INTO flight_plan_flights (flight_plan, flight)
SELECT new_flight_plan.id, f.id
FROM flights AS f, new_flight_plan
WHERE f.flight_number = $1
RETURNING flight_plan_flights.id, flight_plan_flights.flight_plan, flight_plan_flights.flight
`

type CreateFlightPlanParams struct {
	Flightnumber string `json:"flightnumber"`
	Users        int32  `json:"users"`
}

func (q *Queries) CreateFlightPlan(ctx context.Context, arg CreateFlightPlanParams) (FlightPlanFlight, error) {
	row := q.db.QueryRow(ctx, createFlightPlan, arg.Flightnumber, arg.Users)
	var i FlightPlanFlight
	err := row.Scan(&i.ID, &i.FlightPlan, &i.Flight)
	return i, err
}

const deleteFlightPlan = `-- name: DeleteFlightPlan :exec
DELETE FROM flight_plans
WHERE id = $1
RETURNING id, users
`

func (q *Queries) DeleteFlightPlan(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteFlightPlan, id)
	return err
}

const deleteFlightPlanStep = `-- name: DeleteFlightPlanStep :exec
DELETE FROM flight_plan_flights
WHERE id = $1
RETURNING id, flight_plan, flight
`

func (q *Queries) DeleteFlightPlanStep(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteFlightPlanStep, id)
	return err
}

const getFlightPlan = `-- name: GetFlightPlan :many
SELECT 
    fp.id,
    f.flight_number,
    departure_airport.name AS dep_airport_name,
    arrival_airport.name AS arr_airport_name,
    arrival_airport.iata AS arr_iata,
    departure_airport.iata AS dep_iata,
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
WHERE fp.users = $1 AND fp.id = $2
`

type GetFlightPlanParams struct {
	Users int32 `json:"users"`
	ID    int32 `json:"id"`
}

type GetFlightPlanRow struct {
	ID             int32            `json:"id"`
	FlightNumber   string           `json:"flight_number"`
	DepAirportName string           `json:"dep_airport_name"`
	ArrAirportName string           `json:"arr_airport_name"`
	ArrIata        string           `json:"arr_iata"`
	DepIata        string           `json:"dep_iata"`
	SchedDepTime   pgtype.Timestamp `json:"sched_dep_time"`
	SchedArrTime   pgtype.Timestamp `json:"sched_arr_time"`
	ActualDepTime  pgtype.Timestamp `json:"actual_dep_time"`
	ActualArrTime  pgtype.Timestamp `json:"actual_arr_time"`
}

func (q *Queries) GetFlightPlan(ctx context.Context, arg GetFlightPlanParams) ([]GetFlightPlanRow, error) {
	rows, err := q.db.Query(ctx, getFlightPlan, arg.Users, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFlightPlanRow
	for rows.Next() {
		var i GetFlightPlanRow
		if err := rows.Scan(
			&i.ID,
			&i.FlightNumber,
			&i.DepAirportName,
			&i.ArrAirportName,
			&i.ArrIata,
			&i.DepIata,
			&i.SchedDepTime,
			&i.SchedArrTime,
			&i.ActualDepTime,
			&i.ActualArrTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFlightPlans = `-- name: GetFlightPlans :many
SELECT 
    flight_plans.id AS flight_plan_id,
    flights.flight_number,
    dep_airport.name AS dep_airport,
    arr_airport.name AS arr_airport,
    dep_airport.iata,
    arr_airport.iata,
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
WHERE flight_plans.users = $1
`

type GetFlightPlansRow struct {
	FlightPlanID  int32            `json:"flight_plan_id"`
	FlightNumber  string           `json:"flight_number"`
	DepAirport    string           `json:"dep_airport"`
	ArrAirport    string           `json:"arr_airport"`
	Iata          string           `json:"iata"`
	Iata_2        string           `json:"iata_2"`
	SchedDepTime  pgtype.Timestamp `json:"sched_dep_time"`
	SchedArrTime  pgtype.Timestamp `json:"sched_arr_time"`
	ActualDepTime pgtype.Timestamp `json:"actual_dep_time"`
	ActualArrTime pgtype.Timestamp `json:"actual_arr_time"`
}

func (q *Queries) GetFlightPlans(ctx context.Context, users int32) ([]GetFlightPlansRow, error) {
	rows, err := q.db.Query(ctx, getFlightPlans, users)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFlightPlansRow
	for rows.Next() {
		var i GetFlightPlansRow
		if err := rows.Scan(
			&i.FlightPlanID,
			&i.FlightNumber,
			&i.DepAirport,
			&i.ArrAirport,
			&i.Iata,
			&i.Iata_2,
			&i.SchedDepTime,
			&i.SchedArrTime,
			&i.ActualDepTime,
			&i.ActualArrTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersByFlightNumber = `-- name: GetUsersByFlightNumber :many
SELECT f.flight_number, f.status, u.public_id
FROM USERS AS u
JOIN flight_plans AS fp
ON fp.id = u.id
JOIN flight_plan_flights AS fpf
ON fpf.flight_plan = fp.id
JOIN flights AS f
ON f.id = fpf.flight
WHERE f.flight_number = $1
`

type GetUsersByFlightNumberRow struct {
	FlightNumber string      `json:"flight_number"`
	Status       string      `json:"status"`
	PublicID     pgtype.UUID `json:"public_id"`
}

func (q *Queries) GetUsersByFlightNumber(ctx context.Context, flightNumber string) ([]GetUsersByFlightNumberRow, error) {
	rows, err := q.db.Query(ctx, getUsersByFlightNumber, flightNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersByFlightNumberRow
	for rows.Next() {
		var i GetUsersByFlightNumberRow
		if err := rows.Scan(&i.FlightNumber, &i.Status, &i.PublicID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const patchFlightPlan = `-- name: PatchFlightPlan :one
INSERT INTO flight_plan_flights (flight_plan, flight)
SELECT fp.id, f.id
FROM flight_plans AS fp
JOIN flights AS f ON f.flight_number = $1
WHERE fp.id = $2
RETURNING flight_plan_flights.id, flight_plan_flights.flight_plan, flight_plan_flights.flight
`

type PatchFlightPlanParams struct {
	FlightNumber string `json:"flight_number"`
	FlightPlan   int32  `json:"flight_plan"`
}

func (q *Queries) PatchFlightPlan(ctx context.Context, arg PatchFlightPlanParams) (FlightPlanFlight, error) {
	row := q.db.QueryRow(ctx, patchFlightPlan, arg.FlightNumber, arg.FlightPlan)
	var i FlightPlanFlight
	err := row.Scan(&i.ID, &i.FlightPlan, &i.Flight)
	return i, err
}
