package user

import (
	"context"
	"errors"

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
	INSERT INTO users (user_id, sub, username, email)
	VALUES ($1, $2, $3,	$4)
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		user.UserId,
		user.Sub,
		user.Username,
		user.Email,
	)
	if err != nil {
		return err
	}

	r.logger.Info("success insert user", zap.String("email", user.Email))
	return nil
}

func (r *repository) GetBySub(ctx context.Context, sub string) (model.User, error) {
	query := `
	SELECT user_id, sub, username, email
	FROM users
	WHERE sub = $1
	LIMIT 1
	`

	var user model.User

	err := r.pool.QueryRow(ctx, query).Scan(
		&user.UserId,
		&user.Sub,
		&user.Username,
		&user.Email,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, appErrors.ErrUserNotFound
	}

	return user, nil
}
