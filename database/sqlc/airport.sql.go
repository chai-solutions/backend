// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: airport.sql

package sqlc

import (
	"context"
)

const getAirportByIATACode = `-- name: GetAirportByIATACode :one
SELECT id, iata, name, latitude, longitude
FROM airports AS a
WHERE a.iata = $1
`

func (q *Queries) GetAirportByIATACode(ctx context.Context, iata string) (Airport, error) {
	row := q.db.QueryRow(ctx, getAirportByIATACode, iata)
	var i Airport
	err := row.Scan(
		&i.ID,
		&i.Iata,
		&i.Name,
		&i.Latitude,
		&i.Longitude,
	)
	return i, err
}

const getAirportByID = `-- name: GetAirportByID :one
SELECT id, iata, name, latitude, longitude
FROM airports AS a
WHERE a.id = $1
`

func (q *Queries) GetAirportByID(ctx context.Context, id int32) (Airport, error) {
	row := q.db.QueryRow(ctx, getAirportByID, id)
	var i Airport
	err := row.Scan(
		&i.ID,
		&i.Iata,
		&i.Name,
		&i.Latitude,
		&i.Longitude,
	)
	return i, err
}

const getAllAirports = `-- name: GetAllAirports :many
SELECT id, iata, name, latitude, longitude
FROM airports
`

func (q *Queries) GetAllAirports(ctx context.Context) ([]Airport, error) {
	rows, err := q.db.Query(ctx, getAllAirports)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Airport
	for rows.Next() {
		var i Airport
		if err := rows.Scan(
			&i.ID,
			&i.Iata,
			&i.Name,
			&i.Latitude,
			&i.Longitude,
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
