package services

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/likhithkp/banking-ledger-service/db/mongo"
	kafkaclient "github.com/likhithkp/banking-ledger-service/kafka"
	"github.com/likhithkp/banking-ledger-service/shared"
)

func ConsumeTransaction(host string, groupId string, topic string) {
	c := kafkaclient.GetConsumer(host, groupId)
	defer c.Close()

	err := c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Printf("Error while subscribing :%v", err.Error())
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigChan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(1 * time.Second)
			if err != nil {
				if err.(kafka.Error).Code() == kafka.ErrTimedOut {
					continue
				}
				log.Printf("Consumer error: %v\n", err)
				continue
			}

			log.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

			transaction := new(shared.Transaction)
			err = json.Unmarshal(ev.Value, &transaction)

			if err != nil {
				log.Printf("Failed to unmarshal the transaction from consumer: %v", err)
				return
			}

			if transaction.Type != "CREDIT" && transaction.Type != "DEBIT" {
				log.Println("Invalid transaction type")
				return
			}

			if transaction.Amount <= 0 {
				log.Println("Invalid transaction amount")
				return
			}

			mongoColl := mongo.GetTransactionCollection()
			_, err = mongoColl.InsertOne(context.TODO(), transaction)
			if err != nil {
				log.Printf("MongoDB insert failed: %v", err)
				return
			}

			if err := UpdateAccount(transaction); err != nil {
				log.Printf("Failed to update balance: %v", err)
				return
			}
		}
	}
}
