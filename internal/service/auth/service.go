package auth

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/sklyar-vlad/selfDev/internal/config"
	appErrors "github.com/sklyar-vlad/selfDev/internal/errors"
	auth "github.com/sklyar-vlad/selfDev/internal/integrations/casdoor"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
)

type UserService interface {
	GetUserBySub(ctx context.Context, userSub string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

type AuthAdapter interface {
	GetAccess(code, state string) (string, error)
	GetUserInfo(token string) (auth.AuthUser, error)
}

type AuthRepository interface {
	CreateSession(ctx context.Context, sessionID string, userID uuid.UUID) error
}

type Service struct {
	userService UserService
	authAdapter AuthAdapter
	repo        AuthRepository
	cfg         config.ConfigJWT
	logger      *zap.Logger
}

func NewService(
	userService UserService,
	authAdapter AuthAdapter,
	repo AuthRepository,
	configJwt config.ConfigJWT,
	logger *zap.Logger,
) *Service {
	return &Service{userService: userService, authAdapter: authAdapter, repo: repo, cfg: configJwt, logger: logger}
}

func (s *Service) Login(ctx context.Context, code string) (string, error) {
	access, err := s.authAdapter.GetAccess(code, "")
	if err != nil {
		return "", err
	}

	authUser, err := s.authAdapter.GetUserInfo(access)
	if err != nil {
		return "", err
	}

	user, err := s.userService.GetUserBySub(ctx, authUser.Sub)
	if err != nil {
		if errors.Is(err, appErrors.ErrUserNotFound) {
			s.logger.Error("not found user", zap.String("email", authUser.Email), zap.Error(err))
			user, err = s.userService.CreateUser(ctx, model.NewUser(authUser.Sub, authUser.Name, authUser.Name))
			if err != nil {
				s.logger.Error("failed create user", zap.String("email", authUser.Name), zap.Error(err))
				return "", err
			}
		} else {
			return "", err
		}
		return "", err
	}

	s.logger.Info("user", zap.String("user_id", user.UserId.String()))

	sessionID := uuid.NewString()
	err = s.repo.CreateSession(ctx, sessionID, user.UserId)
	if err != nil {
		return "", err
	}

	s.logger.Info("session", zap.String("sessionid", sessionID))

	return sessionID, nil
}
