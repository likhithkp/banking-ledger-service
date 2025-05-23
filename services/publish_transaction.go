package services

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/likhithkp/banking-ledger-service/kafkaclient"
)

func PublishTransaction(topic, key string, byteData []byte, host string) error {
	p := kafkaclient.GetProducer(host)

	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          byteData,
	}, nil)

	if err != nil {
		log.Printf("Kafka publish failed: %v", err)
		return err
	}

	p.Flush(500)
	return nil
}
