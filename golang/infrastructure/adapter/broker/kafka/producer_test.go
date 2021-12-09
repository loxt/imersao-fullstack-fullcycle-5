package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/loxt/imersao-fullstack-fullcycle-5/domain/entity"
	"github.com/loxt/imersao-fullstack-fullcycle-5/infrastructure/adapter/presenter/transaction"
	"github.com/loxt/imersao-fullstack-fullcycle-5/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you don't have limit for this transaction",
	}

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
