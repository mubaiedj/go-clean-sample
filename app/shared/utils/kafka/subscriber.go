package kafka

import (
	"context"
	"encoding/json"
	"github.com/mubaiedj/go-clean-sample/app/interfaces/kafka/constant_kafka"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
	"github.com/segmentio/kafka-go"
	"time"
)

type kafkaSubscriber struct {
	groupID string
	brokers []string
}

func (k *kafkaSubscriber) getKafkaReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        k.brokers,
		GroupID:        k.groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
	})
}

func (k *kafkaSubscriber) listenTopic(topic string, payload interface{}, executeUseCase func(message kafka.Message) error) {
	reader := k.getKafkaReader(topic)
	for {
		ctx := context.Background()
		msg, err := reader.FetchMessage(ctx)

		if err != nil {
			log.Error("%s error trying to read message: %v", topic, err.Error())
			reader.CommitMessages(ctx, msg)
			continue
		}
		log.Info("%s consuming message values: [ topic: %s, partition: %d, offset: %d, key: %s ]", topic, msg.Topic, msg.Partition, msg.Offset, msg.Key)

		err = json.Unmarshal(msg.Value, &payload)
		if err != nil {
			log.Error("%s error trying to unmarshal, event from kafka queue with error: [ %s ] value: [ %s ]", topic, err.Error(), string(msg.Value))
			if err != nil {
				log.Error("%s error trying to record error message [ %s ]", topic, err.Error())
			}
			reader.CommitMessages(ctx, msg)
			continue
		}

		log.Info("%s transformed message values: [ topic: %s, partition: %d, offset: %d, key: %s ]", topic, msg.Topic, msg.Partition, msg.Offset, msg.Key)
		for {
			err = executeUseCase(msg)
			if err != nil {
				log.Error("%s error trying to save message with error : [ %s ] value: [ %s ]", topic, err.Error(), utils.EntityToJson(payload))
				time.Sleep(constant_kafka.SLEEP_TIME)
				continue
			}
			reader.CommitMessages(ctx, msg)
			break
		}

		log.Info("%s consumed message values: [ topic: %s, partition: %d, offset: %d, key: %s ]", topic, msg.Topic, msg.Partition, msg.Offset, msg.Key)
	}
}

func NewKafkaSubscriber(groupID string, brokers ...string) *kafkaSubscriber {
	return &kafkaSubscriber{
		groupID: groupID,
		brokers: brokers,
	}
}
