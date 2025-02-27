package micro

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log/slog"
)

type Producer interface {
	Send(topic string, key string, event any) error
	Close() error
}

type producer struct {
	producer sarama.SyncProducer
}

func newProducer(broker string) (*producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	p, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &producer{
		producer: p,
	}, nil
}

// Send sends a message to a kafka topic.
func (p *producer) Send(topic string, key string, event any) error {
	// Marshal the event to JSON
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Send the message
	partition, offset, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(data),
	})
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	slog.Info(
		"sent message to kafka",
		"topic", topic,
		"partition", partition,
		"offset", offset,
	)

	return nil
}

// Close closes the producer and releases all resources.
func (p *producer) Close() error {
	return p.producer.Close()
}
