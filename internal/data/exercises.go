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
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
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

func (m ExerciseModel) GetAll() ([]*Exercise, error) {
	query := `SELECT id, name, description, category, muscle_group, created_at, updated_at from exercises`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	exercises := []*Exercise{}

	for rows.Next() {
		var exercise Exercise

		err := rows.Scan(
			&exercise.ID,
			&exercise.Name,
			&exercise.Description,
			&exercise.Category,
			&exercise.MuscleGroup,
			&exercise.CreatedAt,
			&exercise.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		exercises = append(exercises, &exercise)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return exercises, nil
}
