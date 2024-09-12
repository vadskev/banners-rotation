package env

import (
	"os"
	"strings"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

const (
	brokersProducerEnvName = "KAFKA_BROKERS"
)

var _ KafkaProducerConfig = (*kafkaProducerConfig)(nil)

type KafkaProducerConfig interface {
	Brokers() []string
	Config() *sarama.Config
}

type kafkaProducerConfig struct {
	brokers []string
}

func NewKafkaProducerConfig() (KafkaProducerConfig, error) {
	brokersStr := os.Getenv(brokersProducerEnvName)
	if len(brokersStr) == 0 {
		return nil, errors.New("kafka brokers address not found")
	}

	brokers := strings.Split(brokersStr, ",")

	return &kafkaProducerConfig{
		brokers: brokers,
	}, nil
}

func (cfg *kafkaProducerConfig) Brokers() []string {
	return cfg.brokers
}

func (cfg *kafkaProducerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	return config
}
