package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/likhithkp/banking-ledger-service/db/psql"
	"github.com/likhithkp/banking-ledger-service/shared"
)

func ValidateAccount(ID string, newTransaction *shared.Transaction) error {
	var balance float64
	const query = `SELECT balance FROM accounts WHERE id = $1`

	err := psql.DB.QueryRow(context.Background(), query, ID).Scan(&balance)
	if err == pgx.ErrNoRows {
		return fmt.Errorf("account doesn't exist")
	} else if err != nil {
		return err
	}

	if newTransaction.Type == "DEBIT" && balance < newTransaction.Amount {
		return fmt.Errorf("insufficient balance")
	}
	return nil
}
