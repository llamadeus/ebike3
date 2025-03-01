package main

import (
	"github.com/llamadeus/ebike3/packages/auth/adapter/in"
	"github.com/llamadeus/ebike3/packages/auth/adapter/out/persistence"
	"github.com/llamadeus/ebike3/packages/auth/domain/events"
	"github.com/llamadeus/ebike3/packages/auth/domain/service"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/config"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/database"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/utils"
	"log/slog"
	"net/http"
	"os"
)

const (
	serverAddr = ":5001"

	schema = `
CREATE TABLE IF NOT EXISTS users (
	id BIGINT PRIMARY KEY,
	username VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	role VARCHAR(50) NOT NULL,
	last_login TIMESTAMP NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions (
	id BIGINT PRIMARY KEY,
	user_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    
    CONSTRAINT fk_sessions_users
        FOREIGN KEY (user_id)
        REFERENCES users (id)
);`
)

func init() {
	err := config.Load()
	if err != nil {
		slog.Error("cannot load config", "error", err)
		os.Exit(1)
	}
}

func main() {
	snowflake, err := utils.NewSnowflakeGenerator(config.Get().ServiceID, 1)
	if err != nil {
		slog.Error("cannot create snowflake generator", "error", err)
		os.Exit(1)
	}

	db, err := database.Open(database.Options{
		Host:     config.Get().DatabaseHost,
		Port:     config.Get().DatabasePort,
		User:     config.Get().DatabaseUser,
		Password: config.Get().DatabasePassword,
		Database: config.Get().DatabaseName,
	})
	if err != nil {
		slog.Error("failed to open database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	err = database.Migrate(db, schema)
	if err != nil {
		slog.Error("failed to migrate database", "error", err)
		os.Exit(1)
	}

	// Configure kafka
	kafka, err := micro.NewKafka(config.Get().KafkaBroker)
	if err != nil {
		slog.Error("failed to create kafka client", "error", err)
		os.Exit(1)
	}
	defer kafka.Close()

	// Configure services
	authRepository := persistence.NewAuthRepository(db, snowflake)
	authService := service.NewAuthService(kafka, authRepository)
	authEventsProcessor := in.MakeAuthEventsProcessor()

	// Configure service
	mux := http.NewServeMux()
	mux.HandleFunc("GET /auth", in.MakeAuthHandler(authService))
	mux.HandleFunc("POST /login", in.MakeLoginHandler(authService))
	mux.HandleFunc("POST /register", in.MakeRegisterHandler(authService))
	mux.HandleFunc("POST /logout", in.MakeLogoutHandler(authService))

	// Start event processor
	consumer, err := kafka.StartProcessor(events.AuthTopic, config.Get().KafkaGroupID, authEventsProcessor)
	if err != nil {
		slog.Error("failed to start event processor", "error", err)
		os.Exit(1)
	}
	defer consumer.Stop()

	if err := micro.Run(mux, serverAddr); err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}
}
