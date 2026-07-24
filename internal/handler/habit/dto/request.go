package dto

type HabitRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Color       string `json:"color"`
	IsGood      bool   `json:"is_good"`
}
