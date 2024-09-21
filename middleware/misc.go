package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// Always use the Content-Type header "application/json".
func JSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Log HTTP requests made to the console using zerolog.
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This retrieves some information from the response that would otherwise
		// not be available: the status code, and the amount of time the request took
		ew := extendResponseWriter(w)

		logEntry := log.Info().Str("method", r.Method).Str("url", r.URL.Path)
		t1 := time.Now()

		defer func() {
			ew.Done()
			logEntry.Int("status", ew.StatusCode).Str("duration", time.Since(t1).String()).Send()
		}()

		next.ServeHTTP(w, r)
	})
}

type customResponseWriter struct {
	responseWriter http.ResponseWriter
	StatusCode     int
}

func extendResponseWriter(w http.ResponseWriter) *customResponseWriter {
	return &customResponseWriter{w, 0}
}

func (w *customResponseWriter) Write(b []byte) (int, error) {
	return w.responseWriter.Write(b)
}

func (w *customResponseWriter) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *customResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.responseWriter.WriteHeader(statusCode)
}

func (w *customResponseWriter) Done() {
	if w.StatusCode == 0 {
		w.StatusCode = http.StatusOK
	}
}
