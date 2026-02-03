package cache

import (
    "context"
    "encoding/json"
    "fmt"
    "task-management-system/internal/config"
    "time"
    
    "github.com/redis/go-redis/v9"
)

type RedisClient struct {
    client *redis.Client
}

func NewRedisClient(cfg *config.RedisConfig) *RedisClient {
    return &RedisClient{
        client: redis.NewClient(&redis.Options{
            Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
            Password: cfg.Password,
            DB:       cfg.DB,
        }),
    }
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
    jsonValue, err := json.Marshal(value)
    if err != nil {
        return err
    }
    return r.client.Set(ctx, key, jsonValue, expiration).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string, dest interface{}) error {
    val, err := r.client.Get(ctx, key).Result()
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(val), dest)
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
    return r.client.Del(ctx, key).Err()
}

func (r *RedisClient) Close() error {
    return r.client.Close()
}