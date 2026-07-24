-- +goose Up
ALTER TABLE habits_completed
DROP CONSTRAINT habits_completed_habit_id_fkey;

ALTER TABLE habits_completed
ADD CONSTRAINT habits_completed_habit_id_fkey
FOREIGN KEY (habit_id)
REFERENCES habits(habit_id)
ON DELETE CASCADE;


-- +goose Down
ALTER TABLE habits_completed
DROP CONSTRAINT habits_completed_habit_id_fkey;

ALTER TABLE habits_completed
ADD CONSTRAINT habits_completed_habit_id_fkey
FOREIGN KEY (habit_id)
REFERENCES habits(habit_id);