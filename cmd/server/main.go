package main

import (
	"fmt"
	"log"
	"time"
	"context"

	"net/http"
	"github.com/go-chi/chi/v5"

	"github.com/nekto-sns/nekto-server/app/handler"
	"github.com/nekto-sns/nekto-server/app/config"
	"github.com/nekto-sns/nekto-server/app/shared/database"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	dbPool, err := database.NewPool(ctx, cfg.DBUrl)
	defer dbPool.Close()
	if err != nil {
		log.Fatalf("%v", err)
	}

	r := chi.NewRouter()
	r.Get("/hello", handler.Hello)

	fmt.Println("Server is running on " + cfg.Port)
	http.ListenAndServe(cfg.Port, r)
}
