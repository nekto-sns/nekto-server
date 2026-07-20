package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/nekto-sns/nekto-server/app/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Get("/hello", handlers.Hello)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}
