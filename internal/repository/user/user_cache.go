package user

import (
	"context"
	"encoding/json"
	"fmt"
	"redis-demo/internal/repository"
	"time"

	"github.com/go-redis/redis/v8"
	"redis-demo/internal/entity"
)

const userKeyPrefix = "user:"

var _ repository.UserCacheRepository = (*RedisUserCacheRepository)(nil)

// RedisUserCacheRepository — реализация кеша профилей пользователей в Redis
type RedisUserCacheRepository struct {
	client   *redis.Client
	cacheTTL time.Duration
}

// NewRedisUserCacheRepository создаёт репозиторий кеша пользователей
func NewRedisUserCacheRepository(client *redis.Client, cacheTTL time.Duration) *RedisUserCacheRepository {
	return &RedisUserCacheRepository{
		client:   client,
		cacheTTL: cacheTTL,
	}
}

func (r *RedisUserCacheRepository) buildKey(id uint) string {
	return fmt.Sprintf("%s%d", userKeyPrefix, id)
}

// Get возвращает пользователя из кеша
func (r *RedisUserCacheRepository) Get(ctx context.Context, id uint) (*entity.User, error) {
	key := r.buildKey(id)
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var user entity.User
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// Set сохраняет пользователя в кеш
func (r *RedisUserCacheRepository) Set(ctx context.Context, id uint, user *entity.User) error {
	key := r.buildKey(id)
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, r.cacheTTL).Err()
}

// Delete удаляет пользователя из кеша
func (r *RedisUserCacheRepository) Delete(ctx context.Context, id uint) error {
	key := r.buildKey(id)
	return r.client.Del(ctx, key).Err()
}
