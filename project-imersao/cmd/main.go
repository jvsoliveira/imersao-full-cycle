package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jvsoliveira/imersao-full-cycle/adapter/broker/kafka"
	"github.com/jvsoliveira/imersao-full-cycle/adapter/factory"
	"github.com/jvsoliveira/imersao-full-cycle/adapter/presenter/transaction"
	"github.com/jvsoliveira/imersao-full-cycle/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	transactionRepository := repositoryFactory.CreateTransactionRepository()

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

	go consumer.Consume(msgChan)

	usecase := process_transaction.NewProcessTransaction(transactionRepository, producer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
