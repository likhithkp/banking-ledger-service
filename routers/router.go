package routers

import (
	"net/http"

	"github.com/likhithkp/banking-ledger-service/handlers"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc("POST /accounts", handlers.CreateAccount)
	mux.HandleFunc("POST /accounts/{id}/transactions", handlers.CreateTransaction)
	mux.HandleFunc("GET /accounts/{id}", handlers.GetAccountDetails)
	mux.HandleFunc("GET /accounts/{id}/transactions", handlers.GetTransaction)
}
