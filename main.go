package main

import (
	"context"
	"fmt"
	"go-log-saas/internal/adapter/config"
	http "go-log-saas/internal/adapter/http"
	"go-log-saas/internal/core/usecase"
	"go-log-saas/internal/repository"
	postgres "go-log-saas/internal/repository/client/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("Error loading environment variables: ", err)
	}

	zap, _ := zap.NewProduction()
	defer zap.Sync()
	logger := zap.Sugar()
	logger.Info("Starting app: ", config.App.Name, " env: ", config.App.Env)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	pg := postgres.NewDatabase(ctx, config.DB, logger)
	defer pg.Close()

	ingestRepository := repository.NewRepository(pg, logger)
	ingestUseCase := usecase.NewIngestUseCase(logger, ingestRepository)
	handler := http.NewHandler(ingestUseCase, logger)

	router := http.NewRouter(config, *handler, logger)

	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.Address, config.HTTP.Port)

	go func() {
		if err := router.Serve(listenAddr); err != nil {
			logger.Error("Error starting the HTTP server", "error", err)
			cancel()
		}
	}()

	<-stop
	logger.Info("Shutting down gracefully...")
}
