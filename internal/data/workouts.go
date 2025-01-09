package data

import (
	"context"
	"database/sql"
	"sulemankhann/workout-tracker/internal/validator"
	"time"
)

type Workout struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ScheduledAt time.Time `json:"scheduled_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkoutModel struct {
	DB *sql.DB
}

func (m WorkoutModel) CreateWorkoutWithExercises(
	workout *Workout,
	workoutExercises []WorkoutExercise,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
        INSERT INTO workouts (user_id, title, description, scheduled_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at`

	args := []any{
		workout.UserID,
		workout.Title,
		workout.Description,
		workout.ScheduledAt,
	}

	err = tx.QueryRowContext(ctx, query, args...).Scan(
		&workout.ID,
		&workout.CreatedAt,
		&workout.UpdatedAt,
	)
	if err != nil {
		return err
	}

	for _, workoutExercise := range workoutExercises {
		query = `
		INSERT INTO workout_exercises (workout_id, exercise_id, sets, repetitions, weight, rest_interval)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
		args = []any{
			workout.ID,
			workoutExercise.ExerciseID,
			workoutExercise.Sets,
			workoutExercise.Repetitions,
			workoutExercise.Weight,
			workoutExercise.RestInterval,
		}

		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func ValidateWorkout(v *validator.Validator, workout *Workout) {
	v.Check(workout.Title != "", "title", "must be provided")
	v.Check(
		len(workout.Title) <= 500,
		"title",
		"must not be more than 500 bytes long",
	)

	if !workout.ScheduledAt.IsZero() {
		v.Check(
			workout.ScheduledAt.After(
				time.Now(),
			),
			"scheduled_at",
			"must not be in the past",
		)
	}
}
