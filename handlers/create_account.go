package handlers

import (
	"net/http"

	"github.com/likhithkp/banking-ledger-service/helpers"
	"github.com/likhithkp/banking-ledger-service/services"
	"github.com/likhithkp/banking-ledger-service/shared"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.JsonEncoder(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	newAccount := new(shared.Account)
	helpers.JsonDecoder(r, w, &newAccount)
	if newAccount.Name == "" || newAccount.Balance == 0 {
		helpers.JsonEncoder(w, "Missing field/value", http.StatusBadRequest, nil)
		return
	}

	createdAccount, err := services.CreateAccount(newAccount)
	if err != nil {
		helpers.JsonEncoder(w, "Internal server error", http.StatusInternalServerError, nil)
		return
	}

	helpers.JsonEncoder(w, "Account created", http.StatusOK, &createdAccount)
}
