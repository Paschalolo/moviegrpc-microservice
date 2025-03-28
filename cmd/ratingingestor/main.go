package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"movieexample.com/rating/pkg/model"
)

func main() {
	fmt.Println("Creating a kafka producer ")

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	const fileName = "ratingsdata.json"
	fmt.Println("reading data from file " + fileName)
	ratingEvents, err := readratingEvents(fileName)
	if err != nil {
		panic(err)
	}
	const topic = "ratings"
	if err := produceRatingEvents(topic, producer, ratingEvents); err != nil {
		panic(err)
	}
	const timeout = 10 * time.Second
	fmt.Println("waiting" + timeout.String() + "until all events get produced ")
	producer.Flush(int(timeout.Milliseconds()))
}

func readratingEvents(file string) ([]model.RatingEvent, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var ratings []model.RatingEvent
	if err := json.NewDecoder(f).Decode(&ratings); err != nil {
		return nil, err
	}
	return ratings, nil
}

func produceRatingEvents(topic string, p *kafka.Producer, events []model.RatingEvent) error {
	for _, ratingEvent := range events {
		encodedEvent, err := json.Marshal(ratingEvent)
		if err != nil {
			return err
		}
		if err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(encodedEvent),
		}, nil); err != nil {
			return err
		}
	}
	return nil
}
