package repository

import (
	"context"
	"github.com/ArtyomYatsenko/gateway/internal/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository struct {
	ctx context.Context
	rdb *redis.Client
}

func NewRedisRepository(config config.DataBaseConfig) *RedisRepository {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address + ":" + config.Port,
		Password: config.Password,
		DB:       config.NumberDB,
	})

	return &RedisRepository{
		ctx: ctx,
		rdb: rdb,
	}
}

func (r *RedisRepository) Set(key string, value string, expiration time.Duration) error {
	return r.rdb.Set(r.ctx, key, value, expiration).Err()
}

func (r *RedisRepository) Get(key string) (string, error) {
	return r.rdb.Get(r.ctx, key).Result()
}

func (r *RedisRepository) Exists(key string) (bool, error) {
	n, err := r.rdb.Exists(r.ctx, key).Result()
	return n > 0, err
}
