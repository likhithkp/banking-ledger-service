package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/likhithkp/banking-ledger-service/shared"
)

func JsonEncoder(w http.ResponseWriter, message string, statusCode uint, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))

	err := json.NewEncoder(w).Encode(&shared.Response{
		Message:    message,
		StatusCode: statusCode,
		Data:       &data,
	})

	if err != nil {
		log.Printf("Failed to encode the response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
