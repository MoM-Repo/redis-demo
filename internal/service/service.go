package service

import (
	"context"

	"redis-demo/internal/dto"
)

// UserService интерфейс для работы с пользователями
type UserService interface {
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUserByID(ctx context.Context, id uint) (*dto.UserResponse, error)
	GetUserByIDWithCache(ctx context.Context, id uint) (*dto.UserResponse, error)
	GetAllUsers(ctx context.Context) ([]*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id uint, req *dto.CreateUserRequest) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, id uint) error
	GetUserStats(ctx context.Context, id uint) (*dto.UserStatsResponse, error)
	GetUserStatsWithCache(ctx context.Context, id uint) (*dto.UserStatsResponse, error)
}
