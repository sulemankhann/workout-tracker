package main

import (
	"errors"
	"fmt"
	"net/http"
	"sulemankhann/workout-tracker/internal/data"
	"sulemankhann/workout-tracker/internal/validator"
	"time"
)

func (app *application) createWorkoutHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		ScheduledAt time.Time `json:"scheduled_at"`
		Exercises   []struct {
			ExerciseID   int64   `json:"exercise_id"`
			Sets         int     `json:"sets"`
			Repetitions  int     `json:"repetitions"`
			Weight       float64 `json:"weight"`
			RestInterval int     `json:"rest_interval"`
		} `json:"exercises"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	workout := &data.Workout{
		UserID:      user.ID,
		Title:       input.Title,
		Description: input.Description,
		ScheduledAt: input.ScheduledAt,
	}

	v := validator.New()

	if data.ValidateWorkout(v, workout); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	workoutExercises := []data.WorkoutExercise{}

	for _, exerciseInput := range input.Exercises {
		exercise, err := app.models.Exercises.Get(exerciseInput.ExerciseID)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				v.AddError(
					"exercises",
					fmt.Sprintf(
						"exercise %d could not be found",
						exerciseInput.ExerciseID,
					),
				)
				app.failedValidationResponse(w, r, v.Errors)
			default:
				app.serverErrorResponse(w, r, err)
			}

			return
		}

		workoutExercise := data.WorkoutExercise{
			ExerciseID:   exerciseInput.ExerciseID,
			Exercise:     *exercise,
			Sets:         exerciseInput.Sets,
			Repetitions:  exerciseInput.Repetitions,
			Weight:       exerciseInput.Weight,
			RestInterval: exerciseInput.RestInterval,
		}

		workoutExercises = append(workoutExercises, workoutExercise)

		if data.ValidateWorkoutEXercise(v, &workoutExercise); !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}

	}

	workout.Exercises = workoutExercises

	err = app.models.Workouts.CreateWorkoutWithExercises(workout)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(
		w,
		http.StatusCreated,
		envelope{"workout": workout},
		nil,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listWorkoutsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	user := app.contextGetUser(r)

	workouts, err := app.models.Workouts.GetAllForUser(user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(
		w,
		http.StatusOK,
		envelope{"workouts": workouts},
		nil,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
