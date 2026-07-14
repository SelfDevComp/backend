package dto

import "github.com/google/uuid"

type HabitRequest struct {
	UserId      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsGood      bool      `json:"is_good"`
}
