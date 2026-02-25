package handler

import (
	"net/http"

	"redis-demo/internal/dto"

	"github.com/gin-gonic/gin"
)

// CreateUser создаёт нового пользователя
func (r *Router) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат JSON"})
		return
	}

	user, err := r.userService.CreateUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно создан",
		"data":    user,
	})
}
