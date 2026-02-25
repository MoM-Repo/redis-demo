package dto

import (
	"time"

	"redis-demo/internal/entity"
)

// CreateUserRequest - запрос на создание/обновление пользователя
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"min=1,max=120"`
}

// UserResponse - ответ с данными пользователя
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FromUser конвертирует доменную модель User в DTO для ответа
func FromUser(u *entity.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Age:       u.Age,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
