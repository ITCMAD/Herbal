package pkg

import "github.com/go-redis/redis/v8"

type Manager struct {
	RedisClient *redis.Client
}

func NewManager(client *redis.Client) *Manager {
	return &Manager{RedisClient: client}
}
