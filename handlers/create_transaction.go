package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/likhithkp/banking-ledger-service/helpers"
	"github.com/likhithkp/banking-ledger-service/services"
	"github.com/likhithkp/banking-ledger-service/shared"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.JsonEncoder(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	accountID := r.PathValue("id")
	helpers.ValidateID(w, accountID)

	newTransaction := new(shared.Transaction)
	helpers.JsonDecoder(r, w, &newTransaction)

	err := services.ValidateAccount(accountID, newTransaction)
	if err != nil {
		helpers.JsonEncoder(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	if newTransaction.Amount == 0 {
		helpers.JsonEncoder(w, "Transaction amount cannot be 0", http.StatusBadRequest, nil)
		return
	}

	if newTransaction.Type != "CREDIT" && newTransaction.Type != "DEBIT" {
		helpers.JsonEncoder(w, "Invalid transaction type(CREDIT/DEBIT)", http.StatusBadRequest, nil)
		return
	}

	txnID := uuid.New().String()
	newTransaction.TransactionID = txnID
	newTransaction.AccountID = accountID
	newTransaction.CreatedAt = time.Now().UTC()

	bytes, err := json.Marshal(newTransaction)
	if err != nil {
		helpers.JsonEncoder(w, "Internal server error", http.StatusInternalServerError, nil)
		log.Printf("Error while marshalling: %v\n", err)
		return
	}

	err = services.PublishTransaction("transactions", "transaction", bytes, "kafka:9092")
	if err != nil {
		log.Printf("Failed to publish transaction: %v\n", err)
		helpers.JsonEncoder(w, "Internal server error", http.StatusInternalServerError, nil)
		return
	}
}
