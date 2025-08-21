package main

import (
	"fmt"
	"go-log-saas/internal/adapter/config"
	http "go-log-saas/internal/adapter/http"
	"go-log-saas/internal/core/usecase"
	"go-log-saas/internal/repository"
	"log"
	"os"

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

	ingestRepository := repository.NewRepository(logger)
	ingestUseCase := usecase.NewIngestUseCase(logger, ingestRepository)
	handler := http.NewHandler(ingestUseCase, logger)

	router := http.NewRouter(config, *handler, logger)

	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.Address, config.HTTP.Port)
	err = router.Serve(listenAddr)
	if err != nil {
		logger.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}

	// Dependency injection
	//database := postgres.NewDatabase(ctx, config.DB, logger)
	//aws := aws.NewAWS(ctx, logger)
}
