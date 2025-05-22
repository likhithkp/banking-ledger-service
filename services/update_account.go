package services

import (
	"context"
	"log"

	"github.com/likhithkp/banking-ledger-service/db/psql"
	"github.com/likhithkp/banking-ledger-service/shared"
)

func UpdateAccount(transaction *shared.Transaction) error {
	query := ""
	if transaction.Type == "CREDIT" {
		query = "UPDATE accounts SET balance = balance + $1 WHERE id = $2"
	} else {
		query = "UPDATE accounts SET balance = balance - $1 WHERE id = $2"
	}

	cmdTag, err := psql.DB.Exec(context.Background(), query, transaction.Amount, transaction.AccountID)
	if err != nil {
		log.Printf("Failed to update the account balance: %v", err)
		return err
	}

	if cmdTag.Insert() {
		log.Println("Transaction success, Account updated successfully")
	}

	return nil
}
