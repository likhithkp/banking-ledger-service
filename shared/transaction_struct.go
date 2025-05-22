package shared

import (
	"time"
)

type Transaction struct {
	TransactionID string    `json:"transaction_id"`
	AccountID     string    `json:"account_id"`
	Type          string    `json:"type"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
