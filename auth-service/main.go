package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "time"

	"github.com/DEVunderdog/auth-service/api"
	database "github.com/DEVunderdog/auth-service/database/sqlc"
	"github.com/DEVunderdog/auth-service/token"
	"github.com/DEVunderdog/auth-service/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	config, err := utils.LoadConfig("auth-service.env")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := database.NewStore(connPool)

	err = token.InitializeJWTKeys(config.Passphrase, store, ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot generate keys")
	}

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	srv := server.Start(config.HTTPServer)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting")
}
