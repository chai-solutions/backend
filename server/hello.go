package server

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
)

type helloRes struct {
	Message string `json:"message"`
	Number  int    `json:"number"`
}

func (a *App) HelloHandler(w http.ResponseWriter, _ *http.Request) {
	number := rand.IntN(69_420_000)

	err := json.NewEncoder(w).Encode(helloRes{
		Message: "Hello, world! The server is alive.",
		Number:  number,
	})
	_ = err
}
