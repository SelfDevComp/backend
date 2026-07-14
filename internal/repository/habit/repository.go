package habit

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	model "github.com/sklyar-vlad/selfDev/internal/model/habit"
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

func (r *repository) GetHabits(ctx context.Context, userId uuid.UUID) ([]model.Habit, error) {
	query := `
	SELECT habit_id, name, description, is_good
	FROM habits
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.pool.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get habits: %w", err)
	}
	defer rows.Close()

	var habits []model.Habit

	for rows.Next() {
		var h model.Habit

		err := rows.Scan(
			&h.HabitId,
			&h.Name,
			&h.Description,
			&h.IsGood,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan habit: %w", err)
		}

		habits = append(habits, h)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	r.logger.Info("success get habits", zap.String("user_id", userId.String()))
	return habits, nil
}

func (r *repository) CreateHabit(ctx context.Context, habit model.Habit) error {
	query := `
	INSERT INTO habits (user_id, habit_id, name, description, is_good)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.pool.Exec(ctx, query, habit.UserId, habit.HabitId, habit.Name, habit.Description, habit.IsGood)
	if err != nil {
		return fmt.Errorf("failed insert habit: %w", err)
	}

	r.logger.Info("success insert habit", zap.String("user_id", habit.UserId.String()))
	return nil
}

func (r *repository) DeleteHabit(ctx context.Context, habitId uuid.UUID) error {
	query := `
	DELETE FROM habits
	WHERE habit_id = $1
	`

	_, err := r.pool.Exec(ctx, query, habitId)
	if err != nil {
		return fmt.Errorf("failed delete habit: %w", err)
	}

	r.logger.Info("success delete habit", zap.String("habit_id", habitId.String()))
	return nil
}

func (r *repository) ConfirmHabit(ctx context.Context, habitId uuid.UUID) error {
	query := `
	INSERT INTO habits_completed (habit_id)
	VALUES ($1)
	`

	_, err := r.pool.Exec(ctx, query, habitId)
	if err != nil {
		return fmt.Errorf("failed insert date: %w", err)
	}

	r.logger.Info("success insert date", zap.String("habit_id", habitId.String()))
	return nil
}

func (r *repository) CancelHabit(ctx context.Context, habitId uuid.UUID) error {
	query := `
	DELETE FROM habits_completed
	WHERE habit_id = $1
	`

	_, err := r.pool.Exec(ctx, query, habitId)
	if err != nil {
		return fmt.Errorf("failed delete date: %w", err)
	}

	r.logger.Info("success delete date", zap.String("habit_id", habitId.String()))
	return nil
}

func (r *repository) GetHabitConfirmDates(ctx context.Context, habitId uuid.UUID) ([]model.Date, error) {
	query := `
	SELECT habit_id, completed_at
	FROM habits_completed
	WHERE habit_id = $1
	`

	rows, err := r.pool.Query(ctx, query, habitId)
	if err != nil {
		return nil, fmt.Errorf("failed get dates: %w", err)
	}
	defer rows.Close()

	var dates []model.Date

	for rows.Next() {
		var h model.Date

		err := rows.Scan(
			&h.HabitId,
			&h.Date,
		)
		if err != nil {
			return nil, fmt.Errorf("failed scan dates: %w", err)
		}

		dates = append(dates, h)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	r.logger.Info("success get dates", zap.String("habit_id", habitId.String()))
	return dates, nil
}
