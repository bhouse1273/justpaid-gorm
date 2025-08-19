package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) Healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}
