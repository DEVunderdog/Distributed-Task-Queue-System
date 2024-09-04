package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/DEVunderdog/auth-service/api"
	database "github.com/DEVunderdog/auth-service/database/sqlc"
	"github.com/DEVunderdog/auth-service/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	
	config, err := utils.LoadConfig("auth-service.env")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := database.NewStore(connPool)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServer)
	
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

