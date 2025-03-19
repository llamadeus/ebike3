package config

import "github.com/caarlos0/env/v6"

type Config struct {
	ServiceID int `env:"SERVICE_ID,required"`

	DatabaseHost     string `env:"DB_HOST,required"`
	DatabasePort     int    `env:"DB_PORT"`
	DatabaseUser     string `env:"DB_USER,required"`
	DatabasePassword string `env:"DB_PASSWORD,required"`
	DatabaseName     string `env:"DB_DATABASE,required"`

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
