package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/likhithkp/banking-ledger-service/db/mongo"
	"github.com/likhithkp/banking-ledger-service/shared"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTransaction(accountID string) (*[]shared.Transaction, error) {
	var transactions []shared.Transaction

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := mongo.MongoClient
	collection := client.Database("ledgerdb").Collection("transactions")

	cursor, err := collection.Find(ctx, bson.M{"accountid": accountID})
	if err != nil {
		log.Printf("Failed to find transactions: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var t shared.Transaction
		if err := cursor.Decode(&t); err != nil {
			log.Printf("Failed to decode transaction: %v", err)
			continue
		}
		transactions = append(transactions, t)
		fmt.Println(transactions)
	}

	return &transactions, nil
}
