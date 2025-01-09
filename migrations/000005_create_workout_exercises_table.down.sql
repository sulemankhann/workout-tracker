-- Drop the foreign key constraint for workout_id if it exists
ALTER TABLE IF EXISTS workout_exercises DROP CONSTRAINT IF EXISTS fk_workout;

-- Drop the foreign key constraint for exercise_id if it exists
ALTER TABLE IF EXISTS workout_exercises DROP CONSTRAINT IF EXISTS fk_exercise;

-- Drop the indexes on workout_id and exercise_id if they exist
DROP INDEX IF EXISTS idx_workout_id;
DROP INDEX IF EXISTS idx_exercise_id;

-- Drop the workout_exercises table
DROP TABLE IF EXISTS workout_exercises;
