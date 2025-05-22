package handlers

import (
	"net/http"

	"github.com/likhithkp/banking-ledger-service/helpers"
	"github.com/likhithkp/banking-ledger-service/services"
)

func GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.JsonEncoder(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	accountID := r.PathValue("id")
	helpers.ValidateID(w, accountID)

	accountDetails, err := services.GetAccountDetails(accountID)
	if err != nil {
		helpers.JsonEncoder(w, "Internal server error", http.StatusInternalServerError, nil)
		return
	}

	helpers.JsonEncoder(w, "Account details fetched", http.StatusOK, &accountDetails)
}
