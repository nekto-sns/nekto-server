package handler

import (
	"time"
	"errors"
	"context"
	"log/slog"
	"net/http"
	"github.com/labstack/echo/v5"

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

func (h *userHandler) ByName(c *echo.Context) error {
	name := c.Param("name")

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	user, err := h.svc.ByName(ctx, name)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return c.JSON(http.StatusNotFound,
				      map[string]any{ "code": "NotFound", "msg": "User not found" })
		}
		slog.Error("Request processing failed", "err", err)
		return c.JSON(http.StatusInternalServerError,
			      map[string]any{ "code": "InternalServerError", "msg": "Error" })
	}
	return c.JSON(http.StatusOK, user)
}
