package main

import (
	"database/sql"
	"encoding/json"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/loxt/imersao-fullstack-fullcycle-5/infrastructure/adapter/broker/kafka"
	"github.com/loxt/imersao-fullstack-fullcycle-5/infrastructure/adapter/factory"
	"github.com/loxt/imersao-fullstack-fullcycle-5/infrastructure/adapter/presenter/transaction"
	"github.com/loxt/imersao-fullstack-fullcycle-5/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}
	topics := []string{"transactions"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go func() {
		_ = consumer.Consume(msgChan)
	}()

	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		_ = json.Unmarshal(msg.Value, &input)
		_, _ = usecase.Execute(input)
	}
}
