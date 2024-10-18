package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

const OverpassAPIURL = "http://overpass-api.de/api/interpreter"

type MapResponse struct {
	MapURL string `json:"map_url"`
}

func (a *App) AirportsHandler(w http.ResponseWriter, _ *http.Request) {
	airports, err := a.Queries.GetAllAirports(context.Background())
	if err != nil {
		log.Error().AnErr("AirportsHandler", err).Msg("Failed to retieve airport")
		http.Error(w, "Failed to get airports", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(airports); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (a *App) AirportMapHandler(w http.ResponseWriter, r *http.Request) {
	iata := chi.URLParam(r, "iata")
	if iata == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	airport, err := a.Queries.GetAirport(context.Background(), iata)
	if err != nil {
		log.Error().AnErr("AirportMapHandler", err).Msg("Failed to retrieve airport")
		http.Error(w, "Failed to retrieve airport", http.StatusInternalServerError)
		return
	}

	mapboxToken := os.Getenv("MAPBOX_PUBLIC")
	if mapboxToken == "" {
		http.Error(w, "Mapbox token not set", http.StatusInternalServerError)
		return
	}

	mapURL := fmt.Sprintf(
		"https://api.mapbox.com/styles/v1/mapbox/streets-v11/static/pin-s-l+000(%f,%f)/%f,%f,18/1170x1280?access_token=%s",
		airport.Longitude, airport.Latitude, airport.Longitude, airport.Latitude, mapboxToken,
	)

	resp, err := http.Get(mapURL)
	if err != nil {
		log.Error().AnErr("AirportMapHandler", err).Msg("Failed to retrieve map from Mapbox")
		http.Error(w, "Failed to retrieve map from Mapbox", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to retrieve map from Mapbox", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, "Failed to write map image", http.StatusInternalServerError)
	}
}
