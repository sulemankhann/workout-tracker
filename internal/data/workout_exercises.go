package data

import (
	"sulemankhann/workout-tracker/internal/validator"
	"time"
)

type WorkoutExercise struct {
	ID           int64     `json:"-"`
	WorkoutID    int64     `json:"-"`
	ExerciseID   int64     `json:"-"`
	Exercise     Exercise  `json:"exercise"`
	Sets         int       `json:"set"`
	Repetitions  int       `json:"repetitions"`
	Weight       float64   `json:"weight"` // 0 for bodyweight exercises
	RestInterval int       `json:"rest_interval"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

func ValidateWorkoutEXercise(
	v *validator.Validator,
	workoutExercise *WorkoutExercise,
) {
	v.Check(workoutExercise.Sets > 0, "sets", "must be greater than zero")
	v.Check(
		workoutExercise.Repetitions > 0,
		"repetitions",
		"must be greater than zero",
	)
	v.Check(workoutExercise.Weight >= 0, "weight", "must be zero or greater")
	v.Check(
		workoutExercise.RestInterval >= 0,
		"rest_interval",
		"must be zero or greater",
	)
}
