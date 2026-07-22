package main

import (
	"os"
	"log/slog"
	"time"
	"context"

	"net/http"
	"github.com/go-chi/chi/v5"

	"github.com/nekto-sns/nekto-server/app/handler"
	"github.com/nekto-sns/nekto-server/app/repository"
	"github.com/nekto-sns/nekto-server/app/service"

	"github.com/nekto-sns/nekto-server/app/config"
	"github.com/nekto-sns/nekto-server/app/shared/database"
	"github.com/nekto-sns/nekto-server/app/shared/logger"
)

func main() {
	cfg := config.Load()

	logger.Setup(cfg.IsProd)

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	dbPool, err := database.NewPool(ctx, cfg.DBUrl)
	defer dbPool.Close()
	if err != nil {
		slog.Error("Server startup failed", "error", err)
		os.Exit(1)
	}

	err = database.RunMigrations(ctx, dbPool)
	if err != nil {
		slog.Error("DB migration failed", "error", err)
		os.Exit(1)
	}

	userRepo := repository.NewUserRepository(dbPool)
	userSvc  := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)


	r := chi.NewRouter()
	r.Get("/{name}", userHandler.ByName)

	slog.Info("Server is running", "port", cfg.Port)
	http.ListenAndServe(cfg.Port, r)
}
