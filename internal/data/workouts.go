package data

import (
	"context"
	"database/sql"
	"fmt"
	"sulemankhann/workout-tracker/internal/validator"
	"time"

	"github.com/lib/pq"
)

type Workout struct {
	ID          int64             `json:"id"`
	UserID      int64             `json:"-"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	ScheduledAt time.Time         `json:"scheduled_at"`
	Exercises   []WorkoutExercise `json:"exercises"`
	CreatedAt   time.Time         `json:"-"`
	UpdatedAt   time.Time         `json:"-"`
}

type WorkoutModel struct {
	DB *sql.DB
}

func (m WorkoutModel) CreateWorkoutWithExercises(workout *Workout) error {
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

	for _, workoutExercise := range workout.Exercises {
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

func (m WorkoutModel) GetAllForUser(userID int64) ([]*Workout, error) {
	query := `
	       SELECT id, user_id, title, description, scheduled_at, created_at, updated_at
	       FROM workouts
	       WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to execute query to fetch workouts for user %d: %w",
			userID,
			err,
		)
	}

	defer rows.Close()

	workouts := []*Workout{}
	workoutMap := make(map[int64]*Workout)

	for rows.Next() {
		var workout Workout

		err := rows.Scan(
			&workout.ID,
			&workout.UserID,
			&workout.Title,
			&workout.Description,
			&workout.ScheduledAt,
			&workout.CreatedAt,
			&workout.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan workout row: %w", err)
		}

		workouts = append(workouts, &workout)
		workoutMap[workout.ID] = &workout
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"error occurred while iterating over workout rows: %w",
			err,
		)
	}

	// Fetch exercises for all workouts
	workoutIDs := make([]int64, 0, len(workoutMap))
	for _, workout := range workouts {
		workoutIDs = append(workoutIDs, workout.ID)
	}

	if len(workoutIDs) == 0 {
		return workouts, nil
	}

	query = `
        SELECT 
            we.workout_id, we.sets, we.repetitions, we.weight, we.rest_interval,
            e.id as exercise_id, e.name, e.description, e.category, e.muscle_group
        FROM workout_exercises we
        JOIN exercises e ON we.exercise_id = e.id
        WHERE we.workout_id = ANY($1)
    `
	exerciseRows, err := m.DB.QueryContext(ctx, query, pq.Array(workoutIDs))
	if err != nil {
		return nil, fmt.Errorf(
			"failed to execute query to fetch exercises for workouts %v: %w",
			workoutIDs,
			err,
		)
	}
	defer exerciseRows.Close()

	for exerciseRows.Next() {
		var workoutID int64
		var workoutExercise WorkoutExercise

		err := exerciseRows.Scan(
			&workoutID,
			&workoutExercise.Sets,
			&workoutExercise.Repetitions,
			&workoutExercise.Weight,
			&workoutExercise.RestInterval,
			&workoutExercise.Exercise.ID,
			&workoutExercise.Exercise.Name,
			&workoutExercise.Exercise.Description,
			&workoutExercise.Exercise.Category,
			&workoutExercise.Exercise.MuscleGroup,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to scan exercise row for workout %d: %w",
				workoutID,
				err,
			)
		}

		if workout, exists := workoutMap[workoutID]; exists {
			workout.Exercises = append(workout.Exercises, workoutExercise)
		}
	}

	if err = exerciseRows.Err(); err != nil {
		return nil, fmt.Errorf(
			"error occurred while iterating over exercise rows: %w",
			err,
		)
	}

	return workouts, nil
}

func (m WorkoutModel) DeleteByUser(id, userId int64) error {
	if id < 1 || userId < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM workouts
        WHERE id = $1
        AND user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
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
