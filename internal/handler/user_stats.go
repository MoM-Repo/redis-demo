package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserStats получает статистику пользователя (сложный запрос)
func (r *Router) GetUserStats(c *gin.Context) {
	id, ok := parseUserID(c)
	if !ok {
		return
	}

	stats, err := r.userService.GetUserStats(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Статистика пользователя получена (сложный запрос)",
		"data":    stats,
	})
}

// GetUserStatsWithCache получает статистику пользователя с кешированием
func (r *Router) GetUserStatsWithCache(c *gin.Context) {
	id, ok := parseUserID(c)
	if !ok {
		return
	}

	stats, err := r.userService.GetUserStatsWithCache(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Статистика пользователя получена (с кешем)",
		"data":    stats,
	})
}
