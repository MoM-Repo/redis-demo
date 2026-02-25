package user

import (
	"context"
	"encoding/json"
	"fmt"
	"redis-demo/internal/repository"
	"time"

	"github.com/go-redis/redis/v8"
	"redis-demo/internal/dto"
)

const userStatsKeyPrefix = "user_stats:"

var _ repository.UserStatsCacheRepository = (*RedisUserStatsCacheRepository)(nil)

// RedisUserStatsCacheRepository — реализация кеша статистики по заказам в Redis
type RedisUserStatsCacheRepository struct {
	client   *redis.Client
	cacheTTL time.Duration
}

// NewRedisUserStatsCacheRepository создаёт репозиторий кеша статистики пользователей
func NewRedisUserStatsCacheRepository(client *redis.Client, cacheTTL time.Duration) *RedisUserStatsCacheRepository {
	return &RedisUserStatsCacheRepository{
		client:   client,
		cacheTTL: cacheTTL,
	}
}

func (r *RedisUserStatsCacheRepository) buildKey(id uint) string {
	return fmt.Sprintf("%s%d", userStatsKeyPrefix, id)
}

// Get возвращает статистику из кеша
func (r *RedisUserStatsCacheRepository) Get(ctx context.Context, id uint) (*dto.UserStatsResponse, error) {
	key := r.buildKey(id)
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var stats dto.UserStatsResponse
	if err := json.Unmarshal([]byte(data), &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}

// Set сохраняет статистику в кеш
func (r *RedisUserStatsCacheRepository) Set(ctx context.Context, id uint, stats *dto.UserStatsResponse) error {
	key := r.buildKey(id)
	data, err := json.Marshal(stats)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, r.cacheTTL).Err()
}
