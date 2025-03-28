package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	// "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"movieexample.com/rating/pkg/model"
)

// ingester defines a kafka Ingester

type Ingester struct {
	consumer *kafka.Consumer
	topic    string
}

// New ingester creates a new ingester
func New(addr string, groupID string, topic string) (*Ingester, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": addr,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	return &Ingester{
		consumer: consumer,
		topic:    topic,
	}, nil
}

// Ingest starts ingestion from kafka and returns a channel containng rating events
func (i *Ingester) Ingest(ctx context.Context) (chan model.RatingEvent, error) {
	if err := i.consumer.SubscribeTopics([]string{i.topic}, nil); err != nil {
		return nil, err
	}
	ch := make(chan model.RatingEvent, 1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				i.consumer.Close()
			default:

			}
			msg, err := i.consumer.ReadMessage(-1)
			if err != nil {
				fmt.Println("consumer error " + err.Error())
				continue
			}
			var event model.RatingEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				fmt.Println("unmarshal error " + err.Error())
				continue
			}
			ch <- event
		}
	}()
	return ch, nil
}
