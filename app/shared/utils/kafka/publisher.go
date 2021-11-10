package kafka

import (
	"context"
	"errors"
	"fmt"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"time"
)

type KafkaPublisher struct {
	brokers []string
}

func NewKafkaPublisher(brokers ...string) *KafkaPublisher {
	return &KafkaPublisher{
		brokers: brokers,
	}
}

func (k *KafkaPublisher) Publish(topic string, message interface{}) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          k.brokers,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		CompressionCodec: snappy.NewCompressionCodec(),
		BatchSize:        1,
		BatchTimeout:     10 * time.Millisecond,
	})
	defer w.Close()
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(utils.Guid()),
			Value: []byte(utils.EntityToJson(message)),
		},
	)

	if err != nil {
		return errors.New(fmt.Sprintf("error publishing message into kafka queue: %s", err.Error()))
	}

	return nil
}
