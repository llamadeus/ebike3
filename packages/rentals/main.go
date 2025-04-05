package main

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/llamadeus/ebike3/packages/rentals/adapter/in"
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/persistence"
	"github.com/llamadeus/ebike3/packages/rentals/domain/events"
	"github.com/llamadeus/ebike3/packages/rentals/domain/service"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/config"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/database"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/micro"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/utils"
	"log/slog"
	"net/http"
	"os"
)

const (
	serverAddr = ":5001"

	schema = `
CREATE TABLE IF NOT EXISTS rentals (
	id BIGINT PRIMARY KEY,
	customer_id BIGINT NOT NULL,
	vehicle_id BIGINT NOT NULL,
	"start" TIMESTAMP NOT NULL,
	"end" TIMESTAMP,
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
	defer mongo.Disconnect(context.Background())

	// Configure kafka
	kafka, err := micro.NewKafka(config.Get().KafkaBroker)
	if err != nil {
		slog.Error("failed to create kafka client", "error", err)
		os.Exit(1)
	}
	defer kafka.Close()

	// Configure asynq
	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: config.Get().RedisURI})
	defer asynqClient.Close()

	// Configure services
	vehicleViewRepository := persistence.NewVehicleViewRepository(mongo.Database(config.Get().MongoDatabase).Collection("vehicles"))
	vehicleService := service.NewVehicleService(vehicleViewRepository)

	rentalRepository := persistence.NewRentalRepository(db, snowflake)
	rentalViewRepository := persistence.NewRentalViewRepository(mongo.Database(config.Get().MongoDatabase).Collection(config.Get().MongoCollection))
	rentalService := service.NewRentalService(kafka, asynqClient, rentalRepository, rentalViewRepository, vehicleService)
	rentalEventsProcessor := in.MakeRentalEventsProcessor(rentalService)
	vehicleEventsProcessor := in.MakeVehicleEventsProcessor(vehicleService)
	accountingEventsProcessor := in.MakeAccountingEventsProcessor(rentalService)

	// Configure service
	mux := http.NewServeMux()
	mux.HandleFunc("GET /customers/{id}/rentals/active", in.MakeGetActiveRentalHandler(rentalService))
	mux.HandleFunc("GET /customers/{id}/rentals/past", in.MakeGetPastRentalsHandler(rentalService))
	mux.HandleFunc("POST /rentals/start", in.MakeStartRentalHandler(rentalService))
	mux.HandleFunc("POST /rentals/{id}/stop", in.MakeStopRentalHandler(rentalService))

	// Start event processor
	rentalConsumer, err := kafka.StartProcessor(events.RentalsTopic, config.Get().KafkaGroupID, rentalEventsProcessor)
	if err != nil {
		slog.Error("failed to start event processor", "error", err)
		os.Exit(1)
	}
	defer rentalConsumer.Stop()

	vehicleConsumer, err := kafka.StartProcessor(events.VehiclesTopic, config.Get().KafkaGroupID, vehicleEventsProcessor)
	if err != nil {
		slog.Error("failed to start event processor", "error", err)
		os.Exit(1)
	}
	defer vehicleConsumer.Stop()

	accountingConsumer, err := kafka.StartProcessor(events.AccountingTopic, config.Get().KafkaGroupID, accountingEventsProcessor)
	if err != nil {
		slog.Error("failed to start event processor", "error", err)
		os.Exit(1)
	}
	defer accountingConsumer.Stop()

	// Start task processor
	asyncServer := asynq.NewServer(
		asynq.RedisClientOpt{Addr: config.Get().RedisURI},
		asynq.Config{},
	)
	asyncMux := asynq.NewServeMux()
	in.MakeRentalTasksProcessor(asyncMux, rentalService)

	err = asyncServer.Start(asyncMux)
	if err != nil {
		slog.Error("failed to start task processor", "error", err)
		os.Exit(1)
	}

	//task, err := tasks.NewRentalsChargeActiveRentalTask(1234)
	//if err != nil {
	//	log.Fatalf("could not create task: %v", err)
	//}
	//info, err := client.Enqueue(task, asynq.ProcessIn(10*time.Second))
	//if err != nil {
	//	log.Fatalf("could not schedule task: %v", err)
	//}
	//log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	if err := micro.Run(mux, serverAddr); err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}

	asyncServer.Shutdown()
}
