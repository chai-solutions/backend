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

// struct required by sql to pass 2 flight numbers into params
type GetTwoFlightsParams struct {
	FlightNumbers [2]string // Array to hold both flight numbers
}

// parameters: Flight numbers for 2 different flights
func (a *App) TimeDiff(firstFlightNum, secondFlightNum string) (time.Duration, error) {

	//firstFlight, err := a.Queries.GetFlight(context.Background(), firstFlightNum)
	// if err != nil {
	// 	//http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return 0, err
	// }
	// secondFlight, err := a.Queries.GetFlight(context.Background(), secondFlightNum)
	// if err != nil {
	// 	//http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return 0, err
	// }
	// var flightsNumArray [2]string
	// flightsNumArray[0] = firstFlightNum
	// flightsNumArray[1] = secondFlightNum

	//parameter for query
	var flightsQueryParam sqlc.GetTwoFlightsParams

	flightsQueryParam.FlightNumber1 = firstFlightNum
	flightsQueryParam.FlightNumber2 = secondFlightNum

	//storage for query response
	var flightsArray []sqlc.GetTwoFlightsRow

	flightsArray, err := a.Queries.GetTwoFlights(context.Background(), flightsQueryParam)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return 0, err
	}

	//get arrival times for flights
	firstFlightArr := flightsArray[0].ActualArrTime.Time
	secondFlightArr := flightsArray[1].ActualArrTime.Time

	//if secondFlight comes first, swap flights for readability
	if firstFlightArr.Compare(secondFlightArr) == 1 {
		flightsArray[0], flightsArray[1] = flightsArray[0], flightsArray[1]
		//swap arrival time variables
		firstFlightArr = flightsArray[0].ActualArrTime.Time
		secondFlightArr = flightsArray[1].ActualArrTime.Time
	}

	firstFlightDep := flightsArray[0].ActualDepTime.Time

	//find layover between first and second flight
	var diff time.Duration
	//find difference between times
	diff = secondFlightArr.Sub(firstFlightDep)

	return diff, nil
}
