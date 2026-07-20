package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/nekto-sns/nekto-server/app/handler"
)

func main() {
	r := chi.NewRouter()
	r.Get("/hello", handler.Hello)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}
