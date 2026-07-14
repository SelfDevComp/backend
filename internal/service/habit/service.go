package habit

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"

	model "github.com/sklyar-vlad/selfDev/internal/model/habit"
)

type HabitRepository interface {
	GetHabits(ctx context.Context, userId uuid.UUID) ([]model.Habit, error)
	CreateHabit(ctx context.Context, habit model.Habit) error
	DeleteHabit(ctx context.Context, habitId uuid.UUID) error
	ConfirmHabit(ctx context.Context, habitId uuid.UUID) error
	CancelHabit(ctx context.Context, habitId uuid.UUID) error
	GetHabitConfirmDates(ctx context.Context, habitId uuid.UUID) ([]model.Date, error)
}

type UserService interface {
	// GetById(ctx context.Context, id uuid.UUID) (userModel.User, error)
}

type service struct {
	repo        HabitRepository
	userService UserService
	logger      *zap.Logger
}

func NewService(repo HabitRepository, userService UserService, logger *zap.Logger) *service {
	return &service{
		repo:        repo,
		userService: userService,
		logger:      logger,
	}
}

func (s *service) GetHabits(ctx context.Context, userId uuid.UUID) ([]model.Habit, error) {
	habits, err := s.repo.GetHabits(ctx, userId)
	if err != nil {
		return []model.Habit{}, fmt.Errorf("failed get habits: %w", err)
	}

	s.logger.Info("success get habits", zap.String("user_id", userId.String()))
	return habits, nil
}

func (s *service) CreateHabit(
	ctx context.Context,
	userId uuid.UUID,
	name, description string,
	isGood bool,
) (model.Habit, error) {
	habit, err := model.NewHabit(userId, name, description, isGood)
	if err != nil {
		return model.Habit{}, fmt.Errorf("failed create habit: %w", err)
	}

	if err = s.repo.CreateHabit(ctx, habit); err != nil {
		return model.Habit{}, fmt.Errorf("failed insert habit: %w", err)
	}

	s.logger.Info("success create habit", zap.String("habit_id", habit.HabitId.String()))
	return habit, nil
}

func (s *service) DeleteHabit(ctx context.Context, habitId uuid.UUID) error {
	if err := s.repo.DeleteHabit(ctx, habitId); err != nil {
		return fmt.Errorf("failed delete habit: %w", err)
	}

	s.logger.Info("success delete habit", zap.String("habit_id", habitId.String()))
	return nil
}

func (s *service) ConfirmHabit(ctx context.Context, habitId uuid.UUID) error {
	if err := s.repo.ConfirmHabit(ctx, habitId); err != nil {
		return fmt.Errorf("failed confirm date: %w", err)
	}

	s.logger.Info("success confirm date", zap.String("habit_id", habitId.String()))
	return nil
}

func (s *service) CancelHabit(ctx context.Context, habitId uuid.UUID) error {
	if err := s.repo.CancelHabit(ctx, habitId); err != nil {
		return fmt.Errorf("failed cancel date: %w", err)
	}

	s.logger.Info("success cancel date", zap.String("habit_id", habitId.String()))
	return nil
}

func (s *service) GetHabitConfirmDates(ctx context.Context, habitId uuid.UUID) ([]model.Date, error) {
	dates, err := s.repo.GetHabitConfirmDates(ctx, habitId)
	if err != nil {
		return []model.Date{}, fmt.Errorf("failed get dates: %w", err)
	}

	s.logger.Info("success get dates", zap.String("habit_id", habitId.String()))
	return dates, nil
}
