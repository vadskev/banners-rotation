package kafka

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/vadskev/banners-rotation/internal/logger"
	"github.com/vadskev/banners-rotation/internal/models"
	"github.com/vadskev/banners-rotation/internal/queue"
	"go.uber.org/zap"
)

const (
	topicName = "banners"
)

type queueKafka struct {
	producer sarama.AsyncProducer
}

var _ queue.Queue = (*queueKafka)(nil)

func New(ctx context.Context, brokerList []string, config *sarama.Config) (queue.Queue, error) {
	pr, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return &queueKafka{
		producer: pr,
	}, nil
}

func (q queueKafka) SendMessage(msg models.Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		logger.Error("failed to marshal data", zap.Error(err))
	}

	mg := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder(data),
	}

	q.producer.Input() <- mg

	select {
	case suc := <-q.producer.Successes():
		logger.Debug(
			"message sent",
			zap.String("partition", strconv.FormatInt(int64(suc.Partition), 10)),
			zap.String("offset", strconv.FormatInt(int64(suc.Offset), 10)),
		)
	case err := <-q.producer.Errors():
		logger.Error("failed to marshal data", zap.Error(err))
	}
}

func (q queueKafka) Close() {
	if q.producer != nil {
		defer func() {
			if err := q.producer.Close(); err != nil {
				logger.Error("failed to close producer", zap.Error(err))
			}
		}()
	}
}
