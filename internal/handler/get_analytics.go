package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) GetAnalytics(c *gin.Context) {
	shortCode := c.Param("short_url")
	if strings.TrimSpace(shortCode) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid short code"})
		return
	}

	analytics, err := h.srv.GetAnalytics(c.Request.Context(), shortCode, 10) // top 10 UA
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "short url not found"})
		return
	}

	c.JSON(http.StatusOK, analytics)
}
