package main

import (
	"github.com/n0l3r/fiberplate/internal/config"
	"github.com/n0l3r/fiberplate/internal/database"
	"github.com/n0l3r/fiberplate/internal/repository"
	"github.com/n0l3r/fiberplate/internal/route"
	"github.com/n0l3r/fiberplate/internal/service"
	"github.com/n0l3r/fiberplate/pkg/logger"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	stdLog := logger.NewLogger()

	cfg := config.New(
		config.WithAppConfig(),
		config.WithDBConfig(),
	)

	dbConn, err := database.NewDBConnection(&cfg.DB)
	if err != nil {
		stdLog.Fatal().Err(err).Msg("failed to connect to the database")
	}
	defer dbConn.Close()

	repo := repository.NewRepository(dbConn)
	services := service.NewService(repo)

	rtr := route.NewRouter(
		route.WithAppConfig(&cfg.App),
		route.WithLogger(stdLog),
		route.WithService(services),
	)

	app := rtr.Setup()

	// Graceful shutdown setup
	go func() {
		if err := app.Listen(cfg.App.Host + ":" + cfg.App.Port); err != nil {
			stdLog.Fatal().Err(err).Msg("failed to start the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	stdLog.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		stdLog.Error().Err(err).Msg("Server forced to shutdown")
	} else {
		stdLog.Info().Msg("Server exiting gracefully")
	}
}
