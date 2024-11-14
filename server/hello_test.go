package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Create a new instance of the app (make sure this is correct for your app structure)
	app := &App{}

	// Create a new request to test the HelloHandler
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Create a new ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the HelloHandler with the request and response recorder
	handler := http.HandlerFunc(app.HelloHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	// Check the response body
	var res helloRes
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Errorf("could not decode response body: %v", err)
	}

	// Verify that the "message" field is correct
	expectedMessage := "Hello, world! The server is alive."
	if res.Message != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, res.Message)
	}

	// Verify that the "number" field is a non-negative integer (for this test, just check if it's >= 0)
	if res.Number < 0 {
		t.Errorf("expected number to be non-negative, got %d", res.Number)
	}
}
