package server

import (
	"chai/database/sqlc"
	"context"
	"encoding/json"
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

// parameters: Flight numbers for 2 different flights
func (a *App) timeDiff(firstFlightNum, secondFlightNum string) (time.Duration, error) {

	firstFlight, err := a.Queries.GetFlight(context.Background(), firstFlightNum)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return -1, err
	}
	secondFlight, err := a.Queries.GetFlight(context.Background(), secondFlightNum)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return -1, err
	}

	//get arrival times for flights
	firstFlightArr := firstFlight.ActualArrTime.Time
	secondFlightArr := secondFlight.ActualArrTime.Time

	//if secondflight comes first, swap flights for readability
	if firstFlightArr.Compare(secondFlightArr) == 1 {
		firstFlight, secondFlight = secondFlight, firstFlight
		//swap arrival time variables
		firstFlightArr = firstFlight.ActualArrTime.Time
		secondFlightArr = secondFlight.ActualArrTime.Time
	}

	firstFlightDep := firstFlight.ActualDepTime.Time

	//find layover between first and second flight
	var diff time.Duration
	//find difference between times
	diff = secondFlightArr.Sub(firstFlightDep)

	return diff, nil
}
