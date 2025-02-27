package micro

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/IBM/sarama"
	"log/slog"
)

type Consumer interface {
	Start(processor *EventsProcessor) error
	Ready() <-chan struct{}
	Stop() error
}

type MessageHandler func(event Event[json.RawMessage]) error

type consumer struct {
	client    sarama.ConsumerGroup
	topic     string
	ready     chan struct{}
	ctx       context.Context
	cancel    context.CancelFunc
	processor *EventsProcessor
}

func newConsumer(broker string, topic string, groupID string) (*consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_9_0_0
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	client, err := sarama.NewConsumerGroup([]string{broker}, groupID, config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &consumer{
		client: client,
		topic:  topic,
		ready:  make(chan struct{}),
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// Start starts the consumer and consumes messages from the kafka topic.
func (c *consumer) Start(processor *EventsProcessor) error {
	if c.processor != nil {
		return errors.New("consumer already started")
	}

	c.processor = processor

	for {
		err := c.client.Consume(c.ctx, []string{c.topic}, c)
		if err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				break
			}

			return err
		}

		// If the session has ended, reinitialize the ready channel
		if c.ctx.Err() != nil {
			return c.ctx.Err()
		}
		c.ready = make(chan struct{})
	}

	return nil
}

// Ready returns a channel that is closed when the consumer is ready to consume.
func (c *consumer) Ready() <-chan struct{} {
	return c.ready
}

// Stop stops the consumer and releases all resources.
func (c *consumer) Stop() error {
	c.cancel()

	return nil
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (c *consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

// Cleanup is run at the end of a session.
func (c *consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim processes messages from a kafka topic.
func (c *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	processor := c.processor
	if processor == nil {
		return errors.New("missing event processor")
	}

	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				slog.Info("message channel was closed")
				return nil
			}

			slog.Info(
				"consuming message",
				"topic", message.Topic,
				"key", message.Key,
			)

			var event Event[json.RawMessage]
			err := json.Unmarshal(message.Value, &event)
			if err != nil {
				slog.Error("error unmarshalling event", "error", err)
				return err
			}

			err = processor.Handle(event)
			if err != nil {
				slog.Error(
					"error handling message",
					"error", err,
					"topic", message.Topic,
					"key", message.Key,
				)
			} else {
				session.MarkMessage(message, "")
			}

		case <-session.Context().Done():
			return nil
		}
	}
}
