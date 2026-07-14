package auth

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type repository struct {
	pool   *pgxpool.Pool
	redis  *redis.Client
	logger *zap.Logger
}

func NewRepository(pool *pgxpool.Pool, rds *redis.Client, logger *zap.Logger) *repository {
	return &repository{
		pool:   pool,
		redis:  rds,
		logger: logger,
	}
}

func (r *repository) CreateSession(ctx context.Context, sessionID string, userID uuid.UUID) error {
	return r.redis.Set(ctx, "session:"+sessionID, userID.String(), 30*24*time.Hour).Err()
}

func (r *repository) GetSession(ctx context.Context, sessionID string) (uuid.UUID, error) {
	userIDStr, err := r.redis.Get(ctx, "session:"+sessionID).Result()
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
