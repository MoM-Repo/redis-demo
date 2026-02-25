package repository

import (
	"context"

	"redis-demo/internal/dto"
	"redis-demo/internal/entity"
)

// UserCacheRepository - кеш профилей пользователей в Redis (ключи user:{id}, TTL настраивается при создании)
type UserCacheRepository interface {
	Get(ctx context.Context, id uint) (*entity.User, error)
	Set(ctx context.Context, id uint, user *entity.User) error
	Delete(ctx context.Context, id uint) error
}

// UserStatsCacheRepository - кеш статистики по заказам в Redis (ключи user_stats:{id}, TTL настраивается при создании)
type UserStatsCacheRepository interface {
	Get(ctx context.Context, id uint) (*dto.UserStatsResponse, error)
	Set(ctx context.Context, id uint, stats *dto.UserStatsResponse) error
}

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id uint) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
}

// OrderRepository интерфейс для работы с заказами
type OrderRepository interface {
	GetUserStats(userID uint) (*dto.UserStatsResponse, error)
}
