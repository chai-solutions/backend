// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: flight_plan.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFlightPlan = `-- name: CreateFlightPlan :exec
WITH new_flight_plan AS (
    INSERT INTO flight_plans (users)
    VALUES ($2)
    RETURNING id
)
INSERT INTO flight_plan_flights (flight_plan, flight)
SELECT new_flight_plan.id, f.id
FROM flights AS f, new_flight_plan
WHERE f.flight_number = $1
`

type CreateFlightPlanParams struct {
	FlightNumber string `json:"flight_number"`
	Users        int32  `json:"users"`
}

func (q *Queries) CreateFlightPlan(ctx context.Context, arg CreateFlightPlanParams) error {
	_, err := q.db.Exec(ctx, createFlightPlan, arg.FlightNumber, arg.Users)
	return err
}

const getFlightPlan = `-- name: GetFlightPlan :one
SELECT 
    fp.id,
    f.flight_number,
    departure_airport.name,
    arrival_airport.name,
    f.sched_dep_time,
    f.sched_arr_time,
    f.actual_dep_time,
    f.actual_arr_time
FROM flight_plans AS fp
JOIN flight_plan_flights AS fpf
ON flight_plans.id = fpf.flight_plan
JOIN flights AS f
ON flight_plan_flights.flight = f.id
JOIN airports AS departure_airport
ON flights.dep_airport = d.id
JOIN airports AS arrival_airport
ON flights.arr_airport = a.id
WHERE flight_plans.users = $1 AND fp.id = $2
`

type GetFlightPlanParams struct {
	Users int32 `json:"users"`
	ID    int32 `json:"id"`
}

type GetFlightPlanRow struct {
	ID            int32            `json:"id"`
	FlightNumber  string           `json:"flight_number"`
	Name          string           `json:"name"`
	Name_2        string           `json:"name_2"`
	SchedDepTime  pgtype.Timestamp `json:"sched_dep_time"`
	SchedArrTime  pgtype.Timestamp `json:"sched_arr_time"`
	ActualDepTime pgtype.Timestamp `json:"actual_dep_time"`
	ActualArrTime pgtype.Timestamp `json:"actual_arr_time"`
}

func (q *Queries) GetFlightPlan(ctx context.Context, arg GetFlightPlanParams) (GetFlightPlanRow, error) {
	row := q.db.QueryRow(ctx, getFlightPlan, arg.Users, arg.ID)
	var i GetFlightPlanRow
	err := row.Scan(
		&i.ID,
		&i.FlightNumber,
		&i.Name,
		&i.Name_2,
		&i.SchedDepTime,
		&i.SchedArrTime,
		&i.ActualDepTime,
		&i.ActualArrTime,
	)
	return i, err
}

const getFlightPlans = `-- name: GetFlightPlans :many
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
	FlightPlanID     int32            `json:"flight_plan_id"`
	FlightNumber     string           `json:"flight_number"`
	DepartureAirport string           `json:"departure_airport"`
	ArrivalAirport   string           `json:"arrival_airport"`
	SchedDepTime     pgtype.Timestamp `json:"sched_dep_time"`
	SchedArrTime     pgtype.Timestamp `json:"sched_arr_time"`
	ActualDepTime    pgtype.Timestamp `json:"actual_dep_time"`
	ActualArrTime    pgtype.Timestamp `json:"actual_arr_time"`
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
			&i.DepartureAirport,
			&i.ArrivalAirport,
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

const patchFlightPlan = `-- name: PatchFlightPlan :exec
INSERT INTO flight_plan_flights (flight_plan, flight)
SELECT fp.id, f.id
FROM flight_plans AS fp
JOIN flights AS f ON f.flight_number = $1
WHERE fp.id = $2
`

type PatchFlightPlanParams struct {
	FlightNumber string `json:"flight_number"`
	FlightPlan   int32  `json:"flight_plan"`
}

func (q *Queries) PatchFlightPlan(ctx context.Context, arg PatchFlightPlanParams) error {
	_, err := q.db.Exec(ctx, patchFlightPlan, arg.FlightNumber, arg.FlightPlan)
	return err
}
