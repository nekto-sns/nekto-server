package main

import (
	"fmt"
	"log"
	"time"
	"context"

	"net/http"
	"github.com/go-chi/chi/v5"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nekto-sns/nekto-server/app/handler"
	"github.com/nekto-sns/nekto-server/app/config"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, cfg.DBUrl)
	if err != nil {
		log.Fatalf("Failed to create DB pool: %v\n", err)
	}
	defer dbPool.Close()

	err = dbPool.Ping(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	r := chi.NewRouter()
	r.Get("/hello", handler.Hello)

	fmt.Println("Server is running on" + cfg.Port)
	http.ListenAndServe(cfg.Port, r)
}
