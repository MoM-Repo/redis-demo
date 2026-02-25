package entity

import (
	"time"

	"gorm.io/gorm"
)

// Order — доменная модель заказа
type Order struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	ProductName string         `json:"product_name" gorm:"not null"`
	Amount      float64        `json:"amount" gorm:"type:decimal(10,2);not null"`
	Status      string         `json:"status" gorm:"not null;default:'pending'"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Связи
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
