// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: flight.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getFlight = `-- name: GetFlight :one
SELECT 
    f.id, f.flight_number, f.sched_dep_time, f.actual_arr_time, f.actual_dep_time,
    a.iata AS arrival_iata, a.name AS arrival_name,
    d.iata AS dep_iata, d.name AS dep_name
FROM flights AS f
INNER JOIN airports AS d
    ON d.id = f.dep_airport
INNER JOIN airports AS a
    ON a.id = f.arr_airport
WHERE f.flight_number = $1
`

type GetFlightRow struct {
	ID            int32            `json:"id"`
	FlightNumber  string           `json:"flight_number"`
	SchedDepTime  pgtype.Timestamp `json:"sched_dep_time"`
	ActualArrTime pgtype.Timestamp `json:"actual_arr_time"`
	ActualDepTime pgtype.Timestamp `json:"actual_dep_time"`
	ArrivalIata   string           `json:"arrival_iata"`
	ArrivalName   string           `json:"arrival_name"`
	DepIata       string           `json:"dep_iata"`
	DepName       string           `json:"dep_name"`
}

func (q *Queries) GetFlight(ctx context.Context, flightNumber string) (GetFlightRow, error) {
	row := q.db.QueryRow(ctx, getFlight, flightNumber)
	var i GetFlightRow
	err := row.Scan(
		&i.ID,
		&i.FlightNumber,
		&i.SchedDepTime,
		&i.ActualArrTime,
		&i.ActualDepTime,
		&i.ArrivalIata,
		&i.ArrivalName,
		&i.DepIata,
		&i.DepName,
	)
	return i, err
}

const getFlights = `-- name: GetFlights :many
SELECT 
    f.id, f.flight_number, f.sched_dep_time, f.actual_arr_time, f.actual_dep_time,
    a.iata AS arrival_iata, a.name AS arrival_name,
    d.iata AS dep_iata, d.name AS dep_name
FROM flights AS f
INNER JOIN airports AS d
    ON d.id = f.dep_airport
INNER JOIN airports AS a
    ON a.id = f.arr_airport
WHERE d.iata = $1
AND a.iata = $2
`

type GetFlightsParams struct {
	Dep string `json:"dep"`
	Arr string `json:"arr"`
}

type GetFlightsRow struct {
	ID            int32            `json:"id"`
	FlightNumber  string           `json:"flight_number"`
	SchedDepTime  pgtype.Timestamp `json:"sched_dep_time"`
	ActualArrTime pgtype.Timestamp `json:"actual_arr_time"`
	ActualDepTime pgtype.Timestamp `json:"actual_dep_time"`
	ArrivalIata   string           `json:"arrival_iata"`
	ArrivalName   string           `json:"arrival_name"`
	DepIata       string           `json:"dep_iata"`
	DepName       string           `json:"dep_name"`
}

func (q *Queries) GetFlights(ctx context.Context, arg GetFlightsParams) ([]GetFlightsRow, error) {
	rows, err := q.db.Query(ctx, getFlights, arg.Dep, arg.Arr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFlightsRow
	for rows.Next() {
		var i GetFlightsRow
		if err := rows.Scan(
			&i.ID,
			&i.FlightNumber,
			&i.SchedDepTime,
			&i.ActualArrTime,
			&i.ActualDepTime,
			&i.ArrivalIata,
			&i.ArrivalName,
			&i.DepIata,
			&i.DepName,
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
