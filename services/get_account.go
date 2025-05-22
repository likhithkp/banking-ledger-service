package services

import (
	"context"
	"log"

	"github.com/likhithkp/banking-ledger-service/db/psql"
	"github.com/likhithkp/banking-ledger-service/shared"
)

func GetAccountDetails(accountID string) (*shared.Account, error) {
	accountDetails := new(shared.Account)

	const accQuery = `SELECT id, name, balance, created_at FROM accounts WHERE id = $1`
	err := psql.DB.QueryRow(context.Background(), accQuery, accountID).Scan(
		&accountDetails.ID,
		&accountDetails.Name,
		&accountDetails.Balance,
		&accountDetails.CreatedAt,
	)

	if err != nil {
		log.Printf("Failed to fetch account details: %v", err)
		return nil, err
	}

	return accountDetails, nil
}
