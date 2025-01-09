package main

import (
	"fmt"
	"log"
	"sulemankhann/workout-tracker/internal/data"
)

func (s Seeder) SeedExercises() {
	em := data.ExerciseModel{DB: s.DB}

	exercises := getAllExercises()

	for _, exercise := range exercises {
		if err := em.Insert(&exercise); err != nil {
			log.Fatalf("Failed to seed exercise '%s': %v", exercise.Name, err)
		}
	}

	fmt.Println("Exercises data successfully seeded.")
}

func getAllExercises() []data.Exercise {
	return []data.Exercise{
		{
			Name:        "Push Up",
			Description: "A basic bodyweight exercise that works the chest, shoulders, and triceps.",
			Category:    "Strength",
			MuscleGroup: "Chest",
		},
		{
			Name:        "Squat",
			Description: "A lower body exercise targeting the quads, hamstrings, and glutes.",
			Category:    "Strength",
			MuscleGroup: "Legs",
		},
		{
			Name:        "Plank",
			Description: "An isometric core strength exercise that targets the abs and back.",
			Category:    "Flexibility",
			MuscleGroup: "Core",
		},
		{
			Name:        "Running",
			Description: "A cardiovascular exercise that improves endurance and burns calories.",
			Category:    "Cardio",
			MuscleGroup: "",
		},
		{
			Name:        "Deadlift",
			Description: "A strength exercise that targets the back, glutes, and hamstrings.",
			Category:    "Strength",
			MuscleGroup: "Back",
		},
		{
			Name:        "Bench Press",
			Description: "A strength exercise that works the chest, shoulders, and triceps.",
			Category:    "Strength",
			MuscleGroup: "Chest",
		},
		{
			Name:        "Pull Up",
			Description: "A bodyweight exercise targeting the back and biceps.",
			Category:    "Strength",
			MuscleGroup: "Back",
		},
		{
			Name:        "Lunges",
			Description: "A lower body exercise targeting the quads, hamstrings, and glutes.",
			Category:    "Strength",
			MuscleGroup: "Legs",
		},
		{
			Name:        "Bicep Curl",
			Description: "A strength exercise that isolates the biceps.",
			Category:    "Strength",
			MuscleGroup: "Arms",
		},
		{
			Name:        "Tricep Dips",
			Description: "A bodyweight exercise focusing on the triceps.",
			Category:    "Strength",
			MuscleGroup: "Arms",
		},
		{
			Name:        "Burpees",
			Description: "A full-body exercise combining strength and cardio.",
			Category:    "Cardio",
			MuscleGroup: "",
		},
		{
			Name:        "Mountain Climbers",
			Description: "A cardio exercise that also engages the core and legs.",
			Category:    "Cardio",
			MuscleGroup: "Core",
		},
		{
			Name:        "Shoulder Press",
			Description: "A strength exercise for the shoulders and arms.",
			Category:    "Strength",
			MuscleGroup: "Shoulders",
		},
		{
			Name:        "Bicycle Crunches",
			Description: "A core exercise targeting the abs and obliques.",
			Category:    "Flexibility",
			MuscleGroup: "Core",
		},
		{
			Name:        "Jumping Jacks",
			Description: "A full-body cardio exercise.",
			Category:    "Cardio",
			MuscleGroup: "",
		},
		{
			Name:        "Calf Raises",
			Description: "A lower-body exercise targeting the calves.",
			Category:    "Strength",
			MuscleGroup: "Legs",
		},
		{
			Name:        "Chest Fly",
			Description: "An exercise focusing on the chest muscles.",
			Category:    "Strength",
			MuscleGroup: "Chest",
		},
		{
			Name:        "Leg Press",
			Description: "A lower-body exercise targeting the quads and glutes.",
			Category:    "Strength",
			MuscleGroup: "Legs",
		},
		{
			Name:        "Russian Twists",
			Description: "A core exercise targeting the obliques.",
			Category:    "Flexibility",
			MuscleGroup: "Core",
		},
		{
			Name:        "High Knees",
			Description: "A cardio exercise to elevate heart rate and engage the legs.",
			Category:    "Cardio",
			MuscleGroup: "Legs",
		},
		{
			Name:        "Rowing Machine",
			Description: "A cardio exercise that engages the back, legs, and arms.",
			Category:    "Cardio",
			MuscleGroup: "Back",
		},
		{
			Name:        "Side Plank",
			Description: "A core stability exercise targeting the obliques.",
			Category:    "Flexibility",
			MuscleGroup: "Core",
		},
		{
			Name:        "Hip Thrust",
			Description: "A lower-body exercise focusing on the glutes.",
			Category:    "Strength",
			MuscleGroup: "Glutes",
		},
		{
			Name:        "Farmer's Walk",
			Description: "A strength exercise for the grip, shoulders, and core.",
			Category:    "Strength",
			MuscleGroup: "Full Body",
		},
		{
			Name:        "Jump Squats",
			Description: "A plyometric leg exercise combining strength and cardio.",
			Category:    "Cardio",
			MuscleGroup: "Legs",
		},
	}
}
