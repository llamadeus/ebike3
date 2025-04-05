package main

import (
	"context"
	"github.com/llamadeus/ebike3/packages/customers/adapter/in"
	"github.com/llamadeus/ebike3/packages/customers/adapter/out/persistence"
	"github.com/llamadeus/ebike3/packages/customers/domain/events"
	"github.com/llamadeus/ebike3/packages/customers/domain/service"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/config"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/database"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/micro"
	"log/slog"
	"net/http"
	"os"
)

const (
	serverAddr = ":5001"
)

func init() {
	err := config.Load()
	if err != nil {
		slog.Error("cannot load config", "error", err)
		os.Exit(1)
	}
}

func main() {
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

	// Configure services
	customerViewRepository := persistence.NewCustomerViewRepository(mongo.Database(config.Get().MongoDatabase).Collection(config.Get().MongoCollection))
	customerService := service.NewCustomerService(kafka, customerViewRepository)
	authEventsProcessor := in.MakeAuthEventsProcessor(customerService)
	customerEventsProcessor := in.MakeCustomerEventsProcessor(customerService)
	accountingEventsProcessor := in.MakeAccountingEventsProcessor(customerService)
	rentalsEventsProcessor := in.MakeRentalEventsProcessor(customerService)

	// Configure service
	mux := http.NewServeMux()
	mux.HandleFunc("GET /customers", in.MakeCustomersHandler(customerService))
	mux.HandleFunc("GET /customers/{id}", in.MakeCustomerHandler(customerService))
	mux.HandleFunc("PATCH /customers/{id}/position", in.MakeUpdateCustomerPositionHandler(customerService))

	// Start event processor
	authConsumer, err := kafka.StartProcessor(events.AuthTopic, config.Get().KafkaGroupID, authEventsProcessor)
	if err != nil {
		slog.Error("failed to start event processor", "error", err)
		os.Exit(1)
	}
	defer authConsumer.Stop()

	customerConsumer, err := kafka.StartProcessor(events.CustomersTopic, config.Get().KafkaGroupID, customerEventsProcessor)
	if err != nil {
		slog.Error("failed to start event processor", "error", err)
		os.Exit(1)
	}
	defer customerConsumer.Stop()

	accountingConsumer, err := kafka.StartProcessor(events.AccountingTopic, config.Get().KafkaGroupID, accountingEventsProcessor)
	if err != nil {
		slog.Error("failed to start event processor", "error", err)
		os.Exit(1)
	}
	defer accountingConsumer.Stop()

	rentalConsumer, err := kafka.StartProcessor(events.RentalsTopic, config.Get().KafkaGroupID, rentalsEventsProcessor)
	if err != nil {
		slog.Error("failed to start event processor", "error", err)
		os.Exit(1)
	}
	defer rentalConsumer.Stop()

	if err := micro.Run(mux, serverAddr); err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}
}
