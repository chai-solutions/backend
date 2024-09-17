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
		a.insertAccountHandler(w, r)
	case http.MethodGet:
		a.selectAccountHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *App) insertAccountHandler(w http.ResponseWriter, r *http.Request) {
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

func (a *App) selectAccountHandler(w http.ResponseWriter, r *http.Request) {
	accountIDstr := r.URL.Query().Get("id")
	if accountIDstr == "" {
		http.Error(w, "Missing account ID", http.StatusBadRequest)
		return
	}
	accountID, err := strconv.ParseInt(accountIDstr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)

	}

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
