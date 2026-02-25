package user

import (
	"redis-demo/internal/entity"
	"redis-demo/internal/repository"

	"gorm.io/gorm"
)

// userRepository реализация UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository создает новый экземпляр UserRepository
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

// Create создает нового пользователя
func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

// GetByID получает пользователя по ID
func (r *userRepository) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail получает пользователя по email
func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll получает всех пользователей
func (r *userRepository) GetAll() ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Update обновляет пользователя
func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

// Delete удаляет пользователя (мягкое удаление)
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
