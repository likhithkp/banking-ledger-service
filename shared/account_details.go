package shared

import (
	"time"

	"github.com/google/uuid"
)

type AccountWithTransactions struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Balance      float64       `json:"balance"`
	CreatedAt    time.Time     `json:"created_at"`
	Transactions []Transaction `json:"transactions"`
}
