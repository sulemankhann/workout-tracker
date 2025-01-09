-- Create the workout_exercises table
CREATE TABLE IF NOT EXISTS workout_exercises (
    id bigserial PRIMARY KEY,                    
    workout_id bigint NOT NULL,                  
    exercise_id bigint NOT NULL,                   
    sets int NOT NULL,                             
    repetitions int NOT NULL,                      
    weight float8 NOT NULL DEFAULT 0,              -- Weight used (0 for bodyweight)
    rest_interval int NOT NULL,                    
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),  
    updated_at timestamp with time zone NOT NULL DEFAULT NOW()   
);

-- Create the foreign key constraint for workout_id
ALTER TABLE workout_exercises
    ADD CONSTRAINT fk_workout FOREIGN KEY (workout_id)
    REFERENCES workouts(id) ON DELETE CASCADE;

-- Create the foreign key constraint for exercise_id
ALTER TABLE workout_exercises
    ADD CONSTRAINT fk_exercise FOREIGN KEY (exercise_id)
    REFERENCES exercises(id) ON DELETE CASCADE;

-- Create an index on workout_id for faster lookups
CREATE INDEX IF NOT EXISTS idx_workout_id ON workout_exercises(workout_id);

-- Create an index on exercise_id for faster lookups
CREATE INDEX IF NOT EXISTS idx_exercise_id ON workout_exercises(exercise_id);
