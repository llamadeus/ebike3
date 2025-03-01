package micro

import (
	"log/slog"
	"os"
)

type Kafka interface {
	Producer() Producer
	NewConsumer(topic string, groupID string) (Consumer, error)
	StartProcessor(topic string, groupID string, processor *EventsProcessor) (Consumer, error)
	Close() error
}

type kafka struct {
	broker   string
	producer Producer
}

func NewKafka(broker string) (Kafka, error) {
	p, err := newProducer(broker)
	if err != nil {
		return nil, err
	}

	return &kafka{
		broker:   broker,
		producer: p,
	}, nil
}

func (k *kafka) Producer() Producer {
	return k.producer
}

func (k *kafka) NewConsumer(topic string, groupID string) (Consumer, error) {
	return newConsumer(k.broker, topic, groupID)
}

func (k *kafka) StartProcessor(topic string, groupID string, processor *EventsProcessor) (Consumer, error) {
	consumer, err := k.NewConsumer(topic, groupID)
	if err != nil {
		return nil, err
	}

	go func() {
		err = consumer.Start(processor)
		if err != nil {
			slog.Error("error starting consumer", "error", err)
			os.Exit(1)
		}
	}()
	<-consumer.Ready()

	return consumer, nil
}

func (k *kafka) Close() error {
	if k.producer != nil {
		return k.producer.Close()
	}

	return nil
}
