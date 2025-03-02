package main

import (
	"github.com/llamadeus/ebike3/packages/vehicles/adapter/in"
	"github.com/llamadeus/ebike3/packages/vehicles/adapter/out/persistence"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/events"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/service"
	"github.com/llamadeus/ebike3/packages/vehicles/infrastructure/config"
	"github.com/llamadeus/ebike3/packages/vehicles/infrastructure/database"
	"github.com/llamadeus/ebike3/packages/vehicles/infrastructure/micro"
	"github.com/llamadeus/ebike3/packages/vehicles/infrastructure/utils"
	"log/slog"
	"net/http"
	"os"
)

const (
	serverAddr = ":5001"

	schema = `
CREATE TABLE IF NOT EXISTS vehicles (
	id BIGINT PRIMARY KEY,
	type VARCHAR(255) NOT NULL,
	position_x FLOAT NOT NULL,
	position_y FLOAT NOT NULL,
    battery FLOAT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
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

	mongo, err := database.OpenMongo(config.Get().MongoURI)
	if err != nil {
		slog.Error("failed to open mongo database", "error", err)
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
	vehicleRepository := persistence.NewVehicleRepository(db, snowflake)
	vehicleViewRepository := persistence.NewVehicleViewRepository(mongo.Database(config.Get().MongoDatabase).Collection(config.Get().MongoCollection))
	vehicleService := service.NewVehicleService(kafka, vehicleRepository, vehicleViewRepository)
	vehicleEventsProcessor := in.MakeVehiclesEventsProcessor(vehicleService)

	// Configure service
	mux := http.NewServeMux()
	mux.HandleFunc("GET /vehicles", in.MakeVehiclesHandler(vehicleService))
	mux.HandleFunc("GET /vehicles/available", in.MakeAvailableVehiclesHandler(vehicleService))
	mux.HandleFunc("PUT /vehicles", in.MakeCreateVehicleHandler(vehicleService))
	mux.HandleFunc("DELETE /vehicles/{id}", in.MakeDeleteVehicleHandler(vehicleService))

	// Start event processor
	consumer, err := kafka.StartProcessor(events.VehiclesTopic, config.Get().KafkaGroupID, vehicleEventsProcessor)
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
