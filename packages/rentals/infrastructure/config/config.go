package config

import "github.com/caarlos0/env/v6"

type Config struct {
	ServiceID int `env:"SERVICE_ID,required"`

	DatabaseHost     string `env:"DB_HOST,required"`
	DatabasePort     int    `env:"DB_PORT"`
	DatabaseUser     string `env:"DB_USER,required"`
	DatabasePassword string `env:"DB_PASSWORD,required"`
	DatabaseName     string `env:"DB_DATABASE,required"`

	MongoURI        string `env:"MONGO_URI,required"`
	MongoDatabase   string `env:"MONGO_DATABASE,required"`
	MongoCollection string `env:"MONGO_COLLECTION,required"`

	KafkaBroker  string `env:"KAFKA_BROKER,required"`
	KafkaGroupID string `env:"KAFKA_GROUP_ID,required"`

	RedisURI string `env:"REDIS_URI,required"`
}

var config Config

func Load() error {
	return env.Parse(&config)
}

func Get() Config {
	return config
}
