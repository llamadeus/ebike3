package main

import (
	"github.com/llamadeus/ebike3/packages/auth/adapter/in"
	"github.com/llamadeus/ebike3/packages/auth/adapter/out/persistence"
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
		os.Exit(1)
	}

	err = database.Migrate(db)
	if err != nil {
		os.Exit(1)
	}

	db.MustExec("TRUNCATE TABLE users CASCADE")
	db.MustExec("TRUNCATE TABLE sessions CASCADE")

	// Configure services
	authRepository := persistence.NewAuthRepository(db, snowflake)
	authService := service.NewAuthService(authRepository)

	// Configure service
	mux := http.NewServeMux()
	mux.HandleFunc("GET /auth", in.NewAuthHandler(authService))
	mux.HandleFunc("POST /login", in.MakeLoginHandler(authService))
	mux.HandleFunc("POST /register", in.MakeRegisterHandler(authService))
	mux.HandleFunc("POST /logout", in.MakeLogoutHandler(authService))

	if err := micro.Run(mux, serverAddr); err != nil {
		os.Exit(1)
	}
}
