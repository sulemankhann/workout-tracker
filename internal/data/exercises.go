package data

import (
	"context"
	"database/sql"
	"time"
)

type Exercise struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	MuscleGroup string    `json:"muscle_group"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExerciseModel struct {
	DB *sql.DB
}

func (m ExerciseModel) Insert(exercise *Exercise) error {
	query := `
        INSERT INTO exercises (name, description, category, muscle_group)
        VALUES ($1,$2,$3,$4)`

	args := []any{
		exercise.Name,
		exercise.Description,
		exercise.Category,
		exercise.MuscleGroup,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)

	return err
}
