package order

import (
	"redis-demo/internal/repository"
	"time"

	"redis-demo/internal/dto"

	"gorm.io/gorm"
)

// orderRepository реализация OrderRepository
type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository создает новый экземпляр OrderRepository
func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepository{db: db}
}

// GetUserStats получает статистику пользователя (сложный запрос с JOIN)
func (r *orderRepository) GetUserStats(userID uint) (*dto.UserStatsResponse, error) {
	var result struct {
		UserID          uint       `json:"user_id"`
		UserName        string     `json:"user_name"`
		TotalOrders     int        `json:"total_orders"`
		CompletedOrders int        `json:"completed_orders"`
		TotalAmount     float64    `json:"total_amount"`
		AverageOrder    float64    `json:"average_order"`
		LastOrderDate   *time.Time `json:"last_order_date"`
	}

	// ОЧЕНЬ сложный SQL запрос с множественными подзапросами и вычислениями
	// Добавляем искусственную задержку для демонстрации разницы с кешем
	err := r.db.Raw(`
		WITH user_orders AS (
			SELECT 
				o.user_id,
				o.id,
				o.amount,
				o.status,
				o.created_at,
				ROW_NUMBER() OVER (PARTITION BY o.user_id ORDER BY o.created_at DESC) as rn
			FROM orders o
			WHERE o.user_id = ?
		),
		monthly_stats AS (
			SELECT 
				user_id,
				DATE_TRUNC('month', created_at) as month,
				COUNT(*) as orders_count,
				SUM(amount) as month_total
			FROM orders 
			WHERE user_id = ?
			GROUP BY user_id, DATE_TRUNC('month', created_at)
		),
		complex_calculations AS (
			SELECT 
				u.id as user_id,
				u.name as user_name,
				COUNT(uo.id) as total_orders,
				COUNT(CASE WHEN uo.status = 'completed' THEN 1 END) as completed_orders,
				COALESCE(SUM(uo.amount), 0) as total_amount,
				COALESCE(AVG(uo.amount), 0) as average_order,
				MAX(uo.created_at) as last_order_date,
				-- Дополнительные сложные вычисления
				COUNT(CASE WHEN uo.status = 'pending' THEN 1 END) as pending_orders,
				COUNT(CASE WHEN uo.status = 'cancelled' THEN 1 END) as cancelled_orders,
				COALESCE(SUM(CASE WHEN uo.status = 'completed' THEN uo.amount ELSE 0 END), 0) as completed_amount,
				COALESCE(AVG(CASE WHEN uo.status = 'completed' THEN uo.amount END), 0) as avg_completed_order,
				-- Статистика по периодам
				COUNT(CASE WHEN uo.created_at >= NOW() - INTERVAL '30 days' THEN 1 END) as orders_last_30_days,
				COUNT(CASE WHEN uo.created_at >= NOW() - INTERVAL '7 days' THEN 1 END) as orders_last_7_days,
				-- Сложные агрегации
				COALESCE(STDDEV(uo.amount), 0) as amount_stddev,
				COALESCE(MIN(uo.amount), 0) as min_order,
				COALESCE(MAX(uo.amount), 0) as max_order,
				-- Временные вычисления
				EXTRACT(EPOCH FROM (MAX(uo.created_at) - MIN(uo.created_at))) / 86400 as days_active
			FROM users u
			LEFT JOIN user_orders uo ON u.id = uo.user_id
			WHERE u.id = ?
			GROUP BY u.id, u.name
		)
		SELECT 
			user_id,
			user_name,
			total_orders,
			completed_orders,
			total_amount,
			average_order,
			last_order_date,
			pending_orders,
			cancelled_orders,
			completed_amount,
			avg_completed_order,
			orders_last_30_days,
			orders_last_7_days,
			amount_stddev,
			min_order,
			max_order,
			days_active,
			-- Дополнительные вычисления
			CASE 
				WHEN total_orders > 0 THEN (completed_orders::FLOAT / total_orders::FLOAT) * 100 
				ELSE 0 
			END as completion_rate,
			CASE 
				WHEN orders_last_30_days > 0 THEN total_amount / orders_last_30_days 
				ELSE 0 
			END as avg_daily_spend_30_days,
			-- Искусственная задержка для демонстрации
			pg_sleep(0.1) as delay
		FROM complex_calculations
	`, userID, userID, userID).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	// Проверяем, существует ли пользователь
	if result.UserID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	stats := &dto.UserStatsResponse{
		UserID:          result.UserID,
		UserName:        result.UserName,
		TotalOrders:     result.TotalOrders,
		CompletedOrders: result.CompletedOrders,
		TotalAmount:     result.TotalAmount,
		AverageOrder:    result.AverageOrder,
		LastOrderDate:   result.LastOrderDate,
	}

	return stats, nil
}
