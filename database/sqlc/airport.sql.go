// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: airport.sql

package sqlc

import (
	"context"
)

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