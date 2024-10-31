package server

import (
	"chai/database/sqlc"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/go-chi/chi/v5"
)

func (a *App) FlightsHandler(w http.ResponseWriter, r *http.Request) {
	depAirport := r.URL.Query().Get("departure_airport")
	arrAirport := r.URL.Query().Get("arrival_airport")
	if depAirport == "" {
		http.Error(w, "Missing Departure Airport", http.StatusBadRequest)
		return
	}

	if arrAirport == "" {
		http.Error(w, "Missing Arrival Airport", http.StatusBadRequest)
		return
	}

	params := sqlc.GetFlightsParams{
		Arr: arrAirport,
		Dep: depAirport,
	}

	flights, err := a.Queries.GetFlights(context.Background(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(flights)
	_ = err
}

func (a *App) FlightHandler(w http.ResponseWriter, r *http.Request) {
	flightNumber := chi.URLParam(r, "flightNumber")
	if flightNumber == "" {
		http.Error(w, "Invalid Flight Number", http.StatusBadRequest)
		return
	}

	flight, err := a.Queries.GetFlight(context.Background(), flightNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(
		flight)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func timeDiff(firstDate, lastDate string) (int64, error) {
	//test retrieved strings
	fmt.Println(firstDate + " " + lastDate)

	//we need this for some reason, to tell go how date is set up?
	layout := "2006-01-02 15:04:05"

	//parse first time
	t1, err := time.Parse(layout, firstDate)
	if err != nil {
		return 0, err
	}

	//parse second time
	t2, err := time.Parse(layout, lastDate)
	if err != nil {
		return 0, err
	}

	//define diff variable so it's not trapped in if statement
	var diff time.Duration
	//find absolute difference between times
	diff = t2.Sub(t1).Abs() // Use Abs() for absolute difference

	//convert time difference into int64 (not sure about this output variable type)
	outSec := diff.Seconds()
	outInt64 := int64(outSec)
	return outInt64, nil
}
