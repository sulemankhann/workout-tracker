-- Drop the foreign key constraint if it exists
ALTER TABLE IF EXISTS workouts DROP CONSTRAINT IF EXISTS fk_user;

-- Drop the index if it exists
DROP INDEX IF EXISTS idx_user_id;

-- Drop the workouts table
DROP TABLE IF EXISTS workouts;
