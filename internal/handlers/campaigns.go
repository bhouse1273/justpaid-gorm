package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

// POST /api/campaigns/:id/materialize
func (h *Handler) MaterializeCampaign(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}
	if err := h.DB.Exec("CALL sp_campaign_materialize_actions(?)", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, map[string]any{"status": "queued", "campaignId": id})
}
