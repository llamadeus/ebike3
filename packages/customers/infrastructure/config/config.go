package config

import "github.com/caarlos0/env/v6"

type Config struct {
	ServiceID int `env:"SERVICE_ID,required"`

	MongoURI        string `env:"MONGO_URI,required"`
	MongoDatabase   string `env:"MONGO_DATABASE,required"`
	MongoCollection string `env:"MONGO_COLLECTION,required"`

	KafkaBroker  string `env:"KAFKA_BROKER,required"`
	KafkaGroupID string `env:"KAFKA_GROUP_ID,required"`
}

var config Config

func Load() error {
	return env.Parse(&config)
}

func Get() Config {
	return config
}
