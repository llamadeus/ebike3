package main

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/in"
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/persistence"
	"github.com/llamadeus/ebike3/packages/accounting/domain/service"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/config"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/database"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/utils"
	"log/slog"
	"net/http"
	"os"
)

const (
	serverAddr = ":5001"

	schema = `
CREATE TABLE IF NOT EXISTS payments (
	id BIGINT PRIMARY KEY,
	customer_id BIGINT NOT NULL,
	amount INT NOT NULL,
	status VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS expenses (
	id BIGINT PRIMARY KEY,
	customer_id BIGINT NOT NULL,
	rental_id BIGINT NOT NULL,
	amount INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS preliminary_expenses (
	id BIGINT PRIMARY KEY,
	inquiry_id BIGINT NOT NULL,
	customer_id BIGINT NOT NULL,
	rental_id BIGINT NOT NULL,
	amount INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	expires_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
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
	paymentRepository := persistence.NewPaymentRepository(db, snowflake)
	expenseRepository := persistence.NewExpenseRepository(db, snowflake)
	preliminaryExpenseRepository := persistence.NewPreliminaryExpenseRepository(db, snowflake)
	accountingService := service.NewAccountingService(kafka, db, paymentRepository, expenseRepository, preliminaryExpenseRepository)

	// Configure service
	mux := http.NewServeMux()
	mux.HandleFunc("GET /payments", in.MakeGetPaymentsHandler(accountingService))
	mux.HandleFunc("PUT /payments", in.MakeCreatePaymentHandler(accountingService))
	mux.HandleFunc("PATCH /payments/{id}", in.MakeUpdatePaymentHandler(accountingService))
	mux.HandleFunc("DELETE /payments/{id}", in.MakeDeletePaymentHandler(accountingService))
	mux.HandleFunc("POST /expenses", in.MakeCreateExpenseHandler(accountingService))
	mux.HandleFunc("PUT /preliminary-expenses", in.MakeCreatePreliminaryExpenseHandler(accountingService))
	mux.HandleFunc("POST /preliminary-expenses/{id}/finalize", in.MakeFinalizePreliminaryExpenseHandler(accountingService))

	mux.HandleFunc("GET /customers/{id}/credit-balance", in.MakeGetCreditBalanceForCustomerHandler(accountingService))
	mux.HandleFunc("GET /customers/{id}/payments", in.MakeGetPaymentsForCustomerHandler(accountingService))
	mux.HandleFunc("GET /customers/{id}/expenses", in.MakeGetExpensesForCustomerHandler(accountingService))

	if err := micro.Run(mux, serverAddr); err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}
}
