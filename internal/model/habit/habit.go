package habit

import (
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	HabitId     uuid.UUID
	UserId      uuid.UUID
	Name        string
	Description string
	Category	string
	Color		string
	IsGood      bool
}

type Date struct {
	HabitId uuid.UUID
	Date    time.Time
}

func NewHabit(userId uuid.UUID, name, description, category, color string, isGood bool) (Habit, error) {
	return Habit{uuid.New(), userId, name, description, category, color, isGood}, nil
}
