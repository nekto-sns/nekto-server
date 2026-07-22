package handler

import (
	"time"
	"errors"
	"context"
	"log/slog"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/nekto-sns/nekto-server/app/model"
)

type userService interface{
	ByName(context.Context, string) (*model.User, error)
}

type userHandler struct{
	svc userService
}

func NewUserHandler(svc userService) (*userHandler) {
	return &userHandler{
		svc: svc,
	}
}

func (h *userHandler) ByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	user, err := h.svc.ByName(ctx, name)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		slog.Error("Request processing failed", "err", err)
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}
	render.JSON(w, r, user)
}
