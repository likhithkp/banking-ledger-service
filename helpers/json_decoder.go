package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonDecoder(r *http.Request, w http.ResponseWriter, dest any) {
	if err := json.NewDecoder(r.Body).Decode(&dest); err != nil {
		log.Printf("Failed to decode the request body: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
