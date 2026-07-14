package dto

import (
	"time"

	"github.com/google/uuid"

	model "github.com/sklyar-vlad/selfDev/internal/model/habit"
)

type HabitResponse struct {
	HabitId     uuid.UUID `json:"habit_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsGood      bool      `json:"is_good"`
}
type HabitsResponse struct {
	Habits []HabitResponse
}

func ToHabitResponse(h model.Habit) HabitResponse {
	return HabitResponse{
		HabitId:     h.HabitId,
		Name:        h.Name,
		Description: h.Description,
		IsGood:      h.IsGood,
	}
}

func ToHabitsResponse(habits []model.Habit) HabitsResponse {
	resp := make([]HabitResponse, 0, len(habits))

	for _, h := range habits {
		resp = append(resp, ToHabitResponse(h))
	}

	return HabitsResponse{
		Habits: resp,
	}
}

type HabitDateResponse struct {
	HabitId uuid.UUID
	Date    time.Time
}

type HabitDatesResponse struct {
	Dates []HabitDateResponse
}

func ToHabitDateResponse(h model.Date) HabitDateResponse {
	return HabitDateResponse{
		HabitId: h.HabitId,
		Date:    h.Date,
	}
}

func ToHabitDatesResponse(habits []model.Date) HabitDatesResponse {
	resp := make([]HabitDateResponse, 0, len(habits))

	for _, h := range habits {
		resp = append(resp, ToHabitDateResponse(h))
	}

	return HabitDatesResponse{
		Dates: resp,
	}
}
