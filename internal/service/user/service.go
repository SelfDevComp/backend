package user

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	model "github.com/sklyar-vlad/selfDev/internal/model/user"
)

// TODO: Update(ctx context.Context, user model.User) (model.User, error)
// TODO: Delete(ctx context.Context, user model.User) error
type Repository interface {
	Create(ctx context.Context, user *model.User) error
	GetBySub(ctx context.Context, userSub string) (model.User, error)
	// Update(ctx context.Context, user *model.User) error
	// GetByLogin(ctx context.Context, login string) (model.User, error)
}

type Service struct {
	repo   Repository
	logger *zap.Logger
}

func NewService(repo Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	if err := s.repo.Create(ctx, &user); err != nil {
		return model.User{}, fmt.Errorf("failed insert user: %w", err)
	}
	s.logger.Info("success create user", zap.String("email", user.Email))
	return user, nil
}

func (s *Service) GetUserBySub(ctx context.Context, userSub string) (model.User, error) {
	user, err := s.repo.GetBySub(ctx, userSub)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
