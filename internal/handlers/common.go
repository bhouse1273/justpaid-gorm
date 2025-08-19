package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
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
