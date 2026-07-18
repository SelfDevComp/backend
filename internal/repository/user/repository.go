package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	appErrors "github.com/sklyar-vlad/selfDev/internal/errors"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
)

type repository struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewRepository(pool *pgxpool.Pool, logger *zap.Logger) *repository {
	return &repository{
		pool:   pool,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, user *model.User) error {
	query := `
	INSERT INTO users (user_id, sub, username)
	VALUES ($1, $2, $3)
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		user.UserId,
		user.Sub,
		user.Username,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetBySub(ctx context.Context, sub string) (model.User, error) {
	query := `
	SELECT user_id, sub, username
	FROM users
	WHERE sub = $1
	LIMIT 1
	`
	r.logger.Info("get sub", zap.String("sub", sub))
	var user model.User

	err := r.pool.QueryRow(ctx, query).Scan(
		&user.UserId,
		&user.Sub,
		&user.Username,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, appErrors.ErrUserNotFound
	}

	if err != nil { 
		return model.User{}, fmt.Errorf("failed get user: %v", err)
	}

	return user, nil
}
