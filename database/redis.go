package database

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/sklyar-vlad/selfDev/internal/config"
)

func NewRedis(ctx context.Context, cfg config.ConfigDatabase) (*redis.Client, error) {
	opt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	return client, nil
}
