package helpers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func ValidateID(w http.ResponseWriter, ID string) error {
	if ID == "" {
		JsonEncoder(w, "Account ID is required", http.StatusBadRequest, nil)
		err := errors.New("account ID is required")
		return err
	}

	_, err := uuid.Parse(ID)
	if err != nil {
		JsonEncoder(w, "Invalid account ID", http.StatusBadRequest, nil)
		return err
	}

	return nil
}
