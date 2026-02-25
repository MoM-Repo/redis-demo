package handler

import (
	"net/http"

	"redis-demo/internal/dto"

	"github.com/gin-gonic/gin"
)

// UpdateUser обновляет пользователя
func (r *Router) UpdateUser(c *gin.Context) {
	id, ok := parseUserID(c)
	if !ok {
		return
	}

	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат JSON"})
		return
	}

	user, err := r.userService.UpdateUser(c.Request.Context(), id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Пользователь успешно обновлен",
		"data":    user,
	})
}
