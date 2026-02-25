package service

import (
	"context"
	"errors"
	"fmt"

	"redis-demo/internal/dto"
	"redis-demo/internal/entity"
	"redis-demo/internal/repository"

	"gorm.io/gorm"
)

// userService реализация UserService
type userService struct {
	userRepo           repository.UserRepository
	orderRepo          repository.OrderRepository
	userCacheRepo      repository.UserCacheRepository
	userStatsCacheRepo repository.UserStatsCacheRepository
}

// NewUserService создает новый экземпляр UserService
func NewUserService(
	userRepo repository.UserRepository,
	orderRepo repository.OrderRepository,
	userCacheRepo repository.UserCacheRepository,
	userStatsCacheRepo repository.UserStatsCacheRepository,
) UserService {
	return &userService{
		userRepo:           userRepo,
		orderRepo:          orderRepo,
		userCacheRepo:      userCacheRepo,
		userStatsCacheRepo: userStatsCacheRepo,
	}
}

// CreateUser создает нового пользователя
func (s *userService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("ошибка при проверке существования пользователя: %w", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("пользователь с email %s уже существует", req.Email)
	}

	user := &entity.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("ошибка при создании пользователя: %w", err)
	}

	resp := dto.FromUser(user)
	return &resp, nil
}

// GetUserByID получает пользователя по ID без кеширования
func (s *userService) GetUserByID(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("пользователь с ID %d не найден", id)
		}
		return nil, fmt.Errorf("ошибка при получении пользователя: %w", err)
	}

	resp := dto.FromUser(user)
	return &resp, nil
}

// GetUserByIDWithCache получает пользователя по ID с кешированием
func (s *userService) GetUserByIDWithCache(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user, err := s.userCacheRepo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении кеша: %w", err)
	}
	if user != nil {
		resp := dto.FromUser(user)
		return &resp, nil
	}

	user, err = s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("пользователь с ID %d не найден", id)
		}
		return nil, fmt.Errorf("ошибка при получении пользователя: %w", err)
	}

	if err := s.userCacheRepo.Set(ctx, id, user); err != nil {
		return nil, fmt.Errorf("ошибка при записи в кеш: %w", err)
	}

	resp := dto.FromUser(user)
	return &resp, nil
}

// GetAllUsers получает всех пользователей
func (s *userService) GetAllUsers(ctx context.Context) ([]*dto.UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении пользователей: %w", err)
	}

	var responses []*dto.UserResponse
	for _, user := range users {
		resp := dto.FromUser(user)
		responses = append(responses, &resp)
	}

	return responses, nil
}

// UpdateUser обновляет пользователя
func (s *userService) UpdateUser(ctx context.Context, id uint, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("пользователь с ID %d не найден", id)
		}
		return nil, fmt.Errorf("ошибка при получении пользователя: %w", err)
	}

	if user.Email != req.Email {
		existingUser, err := s.userRepo.GetByEmail(req.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ошибка при проверке существования пользователя: %w", err)
		}
		if existingUser != nil {
			return nil, fmt.Errorf("пользователь с email %s уже существует", req.Email)
		}
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Age = req.Age

	if err := s.userRepo.Update(user); err != nil {
		return nil, fmt.Errorf("ошибка при обновлении пользователя: %w", err)
	}

	if err := s.userCacheRepo.Delete(ctx, id); err != nil {
		return nil, fmt.Errorf("ошибка при инвалидации кеша: %w", err)
	}

	resp := dto.FromUser(user)
	return &resp, nil
}

// DeleteUser удаляет пользователя
func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("пользователь с ID %d не найден", id)
		}
		return fmt.Errorf("ошибка при получении пользователя: %w", err)
	}

	if err := s.userRepo.Delete(id); err != nil {
		return fmt.Errorf("ошибка при удалении пользователя: %w", err)
	}

	if err := s.userCacheRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("ошибка при инвалидации кеша: %w", err)
	}

	return nil
}

// GetUserStats получает статистику пользователя (сложный запрос без кеша)
func (s *userService) GetUserStats(ctx context.Context, id uint) (*dto.UserStatsResponse, error) {
	stats, err := s.orderRepo.GetUserStats(id)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении статистики пользователя: %w", err)
	}

	return stats, nil
}

// GetUserStatsWithCache получает статистику пользователя с кешированием
func (s *userService) GetUserStatsWithCache(ctx context.Context, id uint) (*dto.UserStatsResponse, error) {
	stats, err := s.userStatsCacheRepo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении кеша статистики: %w", err)
	}
	if stats != nil {
		return stats, nil
	}

	stats, err = s.GetUserStats(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.userStatsCacheRepo.Set(ctx, id, stats); err != nil {
		return nil, fmt.Errorf("ошибка при записи статистики в кеш: %w", err)
	}

	return stats, nil
}
