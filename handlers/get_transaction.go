package handlers

import (
	"net/http"

	"github.com/likhithkp/banking-ledger-service/helpers"
	"github.com/likhithkp/banking-ledger-service/services"
)

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.JsonEncoder(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	accountID := r.PathValue("id")
	helpers.ValidateID(w, accountID)

	transactionDetails, err := services.GetTransaction(accountID)
	if err != nil {
		helpers.JsonEncoder(w, "Internal server error", http.StatusInternalServerError, nil)
		return
	}

	helpers.JsonEncoder(w, "Transactions fetched", http.StatusOK, &transactionDetails)

}
