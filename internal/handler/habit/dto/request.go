package dto

type HabitRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsGood      bool   `json:"is_good"`
}
