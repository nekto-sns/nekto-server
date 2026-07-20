package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/nekto-sns/nekto-server/app/handler"
	"github.com/nekto-sns/nekto-server/app/config"
)

func main() {
	cfg := config.Load()
	r := chi.NewRouter()
	r.Get("/hello", handler.Hello)
	fmt.Println("Server is running on" + cfg.Port)
	http.ListenAndServe(cfg.Port, r)
}
