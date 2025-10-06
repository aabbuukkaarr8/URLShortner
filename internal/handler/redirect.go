package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Redirect(c *gin.Context) {
	shortCode := c.Param("short_url")
	ua := c.Request.UserAgent()
	ip := c.ClientIP()

	url, err := h.srv.ResolveAndTrack(c.Request.Context(), shortCode, ua, ip)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "short url not found"})
		return
	}
	c.Redirect(http.StatusFound, url.OriginalURL)
}
