package micro

type Kafka interface {
	Producer() Producer
	NewConsumer(topic string, groupID string) (Consumer, error)
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

func (k *kafka) Close() error {
	if k.producer != nil {
		return k.producer.Close()
	}

	return nil
}
