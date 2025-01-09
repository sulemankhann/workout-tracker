-- Create the workouts table
CREATE TABLE IF NOT EXISTS workouts (
    id bigserial PRIMARY KEY,                   
    user_id bigint NOT NULL,                    
    title text NOT NULL,                        
    description text,                           
    scheduled_at timestamp with time zone,     
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW()
);

-- Create the foreign key constraint for user_id
ALTER TABLE workouts 
ADD CONSTRAINT fk_user FOREIGN KEY (user_id) 
REFERENCES users(id) ON DELETE CASCADE;

-- Create an index on user_id for faster lookups
CREATE INDEX IF NOT EXISTS idx_user_id ON workouts(user_id);
