package dto

import (
	"time"
)

// UserStatsResponse — статистика по заказам пользователя (результат тяжёлого запроса, API output)
type UserStatsResponse struct {
	UserID          uint       `json:"user_id"`
	UserName        string     `json:"user_name"`
	TotalOrders     int        `json:"total_orders"`
	CompletedOrders int        `json:"completed_orders"`
	TotalAmount     float64    `json:"total_amount"`
	AverageOrder    float64    `json:"average_order"`
	LastOrderDate   *time.Time `json:"last_order_date,omitempty"`
}
