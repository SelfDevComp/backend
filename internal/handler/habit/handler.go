package habit

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/sklyar-vlad/selfDev/internal/handler/habit/dto"
	model "github.com/sklyar-vlad/selfDev/internal/model/habit"
)

// TODO: UpdateHabit(ctx context.Context, habitId uuid.UUID) error
type HabitService interface {
	GetHabits(ctx context.Context, userId uuid.UUID) ([]model.Habit, error)
	CreateHabit(ctx context.Context, userId uuid.UUID, name, description string, isGood bool) (model.Habit, error)
	DeleteHabit(ctx context.Context, habitId uuid.UUID) error
	ConfirmHabit(ctx context.Context, habitId uuid.UUID) error
	CancelHabit(ctx context.Context, habitId uuid.UUID) error
	GetHabitConfirmDates(ctx context.Context, habitId uuid.UUID) ([]model.Date, error)
}

type handler struct {
	service HabitService
	logger  *zap.Logger
}

func NewHandler(service HabitService, logger *zap.Logger) *handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func GetID(r *http.Request, path string) (uuid.UUID, error) {
	id := r.PathValue(path)
	return uuid.Parse(id)
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func (h *handler) GetHabits(w http.ResponseWriter, r *http.Request) {
	userId, err := GetID(r, "user_id")
	if err != nil {
		h.logger.Error("invalid id", zap.Error(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	habits, err := h.service.GetHabits(r.Context(), userId)
	if err != nil {
		h.logger.Error("failed get habits", zap.Error(err))
		http.Error(w, "failed get habits", http.StatusInternalServerError)
		return
	}

	if err = WriteJSON(w, http.StatusOK, dto.ToHabitsResponse(habits)); err != nil {
		h.logger.Error("failed create response", zap.String("user_id", userId.String()))
	}
}

func (h *handler) CreateHabit(w http.ResponseWriter, r *http.Request) {
	var input dto.HabitRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Error("failed decode request", zap.Error(err))
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	habit, err := h.service.CreateHabit(r.Context(), input.UserId, input.Name, input.Description, input.IsGood)
	if err != nil {
		h.logger.Error("failed create habit", zap.Error(err))
		http.Error(w, "failed create habit", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(dto.ToHabitResponse(habit)); err != nil {
		h.logger.Error(
			"failed create response",
			zap.String("habit_id", habit.HabitId.String()),
			zap.Error(err),
		)
	}
}

func (h *handler) DeleteHabit(w http.ResponseWriter, r *http.Request) {
	habitId, err := GetID(r, "id")
	if err != nil {
		h.logger.Error("invalid id", zap.Error(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteHabit(r.Context(), habitId)
	if err != nil {
		h.logger.Error("failed delete habit", zap.Error(err))
		http.Error(w, "failed delete habit", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) ConfirmHabit(w http.ResponseWriter, r *http.Request) {
	habitId, err := GetID(r, "id")
	if err != nil {
		h.logger.Error("invalid id", zap.Error(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.service.ConfirmHabit(r.Context(), habitId)
	if err != nil {
		h.logger.Error("failed confirm habit", zap.Error(err))
		http.Error(w, "failed confirm habit", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) CancelHabit(w http.ResponseWriter, r *http.Request) {
	habitId, err := GetID(r, "id")
	if err != nil {
		h.logger.Error("invalid id", zap.Error(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.service.CancelHabit(r.Context(), habitId)
	if err != nil {
		h.logger.Error("failed cancel habit", zap.Error(err))
		http.Error(w, "failed cancel habit", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) GetHabitConfirmDates(w http.ResponseWriter, r *http.Request) {
	habitId, err := GetID(r, "id")
	if err != nil {
		h.logger.Error("invalid id", zap.Error(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	dates, err := h.service.GetHabitConfirmDates(r.Context(), habitId)
	if err != nil {
		h.logger.Error("failed get dates", zap.Error(err))
		http.Error(w, "failed get dates", http.StatusInternalServerError)
		return
	}

	if err = WriteJSON(w, http.StatusOK, dto.ToHabitDatesResponse(dates)); err != nil {
		h.logger.Error("failed create response", zap.String("habit_id", habitId.String()))
	}
}
