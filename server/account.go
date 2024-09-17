package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"chai/database/sqlc"
)

type accountRes struct {
	Message string       `json:"message"`
	Data    sqlc.Account `json:"accounts"`
}

func (a *App) AccountHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		a.InsertAccountHandler(w, r)
	case http.MethodGet:
		a.SelectAccountHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *App) InsertAccountHandler(w http.ResponseWriter, r *http.Request) {
	var params sqlc.InsertAccountParams
	params.Owner = r.URL.Query().Get("name")
	params.Currency = r.URL.Query().Get("currency")
	params.Balance = 0
	newAccount, err := a.Queries.InsertAccount(context.Background(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(
		accountRes{
			Message: "Account created successfully",
			Data:    newAccount,
		})

}

func (a *App) SelectAccountHandler(w http.ResponseWriter, r *http.Request) {
	accountIDstr := r.URL.Query().Get("id")
	if accountIDstr == "" {
		http.Error(w, "Missing account ID", http.StatusBadRequest)
		return
	}
	accountID, err := strconv.ParseInt(accountIDstr, 10, 64)

	account, err := a.Queries.SelectAccountByID(context.Background(), accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(
		accountRes{
			Message: "Account retrieved successfully",
			Data:    account,
		})

}
