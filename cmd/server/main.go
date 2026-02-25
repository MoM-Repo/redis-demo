package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"redis-demo/internal/repository/order"
	"redis-demo/internal/repository/user"
	"syscall"
	"time"

	"redis-demo/internal/config"
	"redis-demo/internal/handler"
	"redis-demo/internal/repository/conn/postgres"
	"redis-demo/internal/repository/conn/redis"
	"redis-demo/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	cfg := config.LoadConfig()
	log.Println("Configuration loaded successfully")

	ctx := context.Background()

	postClient, err := postgres.NewConn(ctx, cfg.Repository.Postgres)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Printf("Connected to database: %s", cfg.Repository.Postgres.Name)

	redisClient, err := redis.NewConn(ctx, cfg.Repository.Redis)
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	defer redisClient.Close()

	userRepo := user.NewUserRepository(postClient.DB())
	orderRepo := order.NewOrderRepository(postClient.DB())

	userCacheRepo := user.NewRedisUserCacheRepository(redisClient.RedisClient(), 5*time.Minute)
	userStatsCacheRepo := user.NewRedisUserStatsCacheRepository(redisClient.RedisClient(), 10*time.Minute)

	userService := service.NewUserService(userRepo, orderRepo, userCacheRepo, userStatsCacheRepo)

	router := handler.NewRouter(userService)
	engine := router.Setup()

	serverAddr := ":" + cfg.App.ServerPort
	go func() {
		log.Printf("Server starting on port %s", cfg.App.ServerPort)
		if err := engine.Run(serverAddr); err != nil {
			log.Fatal("Server failed to start:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}
