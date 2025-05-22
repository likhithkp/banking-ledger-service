package services

import (
	"context"
	"log"

	"github.com/likhithkp/banking-ledger-service/db/psql"
	"github.com/likhithkp/banking-ledger-service/shared"
)

func CreateAccount(newAccount *shared.Account) (*shared.Account, error) {
	const query = `INSERT INTO accounts (name, balance) values ($1, $2) RETURNING id, name, balance, created_at`
	row := psql.DB.QueryRow(context.Background(), query, newAccount.Name, newAccount.Balance)

	createdAccount := new(shared.Account)
	if err := row.Scan(&createdAccount.ID, &createdAccount.Name, &createdAccount.Balance, &createdAccount.CreatedAt); err != nil {
		log.Printf("Failed to scan the row: %v", err)
		return nil, err
	}

	return createdAccount, nil
}
