package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers получает всех пользователей
func (r *Router) GetAllUsers(c *gin.Context) {
	users, err := r.userService.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Список пользователей получен",
		"data":    users,
	})
}
