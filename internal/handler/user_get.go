package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserByID получает пользователя по ID без кеширования
func (r *Router) GetUserByID(c *gin.Context) {
	id, ok := parseUserID(c)
	if !ok {
		return
	}

	user, err := r.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Пользователь получен (без кеша)",
		"data":    user,
	})
}

// GetUserByIDWithCache получает пользователя по ID с кешированием
func (r *Router) GetUserByIDWithCache(c *gin.Context) {
	id, ok := parseUserID(c)
	if !ok {
		return
	}

	user, err := r.userService.GetUserByIDWithCache(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Пользователь получен (с кешем)",
		"data":    user,
	})
}
