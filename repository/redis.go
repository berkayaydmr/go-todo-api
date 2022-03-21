package repository

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisClientInterface interface {
	GetData(key string) string
	SetData(key string, value string, exp time.Duration) error
	DeleteData(key string) error
}

type RedisClient struct {
	redis *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) RedisClientInterface {
	return &RedisClient{redis: redisClient}
}

func (redisClient *RedisClient) SetData(key string, value string, exp time.Duration) error {
	err := redisClient.redis.Set(key, value, exp).Err()
	return err
}

func (redisClient *RedisClient) GetData(key string) string {
	data := redisClient.redis.Get(key)
	return data.Val()
}

func (redisClient *RedisClient) DeleteData(key string) error {
	err := redisClient.redis.Del(key).Err()
	return err
}
