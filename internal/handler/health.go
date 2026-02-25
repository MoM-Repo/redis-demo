package handler

import (
	"github.com/gin-gonic/gin"
)

// HealthCheck проверка здоровья сервера
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "Сервер работает",
		"service": "redis-demo",
	})
}
