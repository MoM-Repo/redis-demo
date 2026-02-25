package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteUser удаляет пользователя
func (r *Router) DeleteUser(c *gin.Context) {
	id, ok := parseUserID(c)
	if !ok {
		return
	}

	err := r.userService.DeleteUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Пользователь успешно удален",
	})
}
