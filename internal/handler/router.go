package handler

import (
	"strconv"

	"redis-demo/internal/service"

	"github.com/gin-gonic/gin"
)

// Router настраивает маршруты для HTTP сервера (Gin)
type Router struct {
	userService service.UserService
}

// NewRouter создаёт новый экземпляр Router
func NewRouter(userService service.UserService) *Router {
	return &Router{
		userService: userService,
	}
}

// Setup настраивает все маршруты и возвращает *gin.Engine
func (r *Router) Setup() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", HealthCheck)

	api := router.Group("/api/v1")
	{
		api.POST("/users", r.CreateUser)
		api.GET("/users", r.GetAllUsers)
		api.GET("/users/:id", r.GetUserByID)
		api.GET("/users/:id/cached", r.GetUserByIDWithCache)
		api.PUT("/users/:id", r.UpdateUser)
		api.DELETE("/users/:id", r.DeleteUser)
		api.GET("/users/:id/stats", r.GetUserStats)
		api.GET("/users/:id/stats/cached", r.GetUserStatsWithCache)
	}

	return router
}

func parseUserID(c *gin.Context) (uint, bool) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный ID пользователя"})
		return 0, false
	}
	return uint(id), true
}
