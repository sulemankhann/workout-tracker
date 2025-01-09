CREATE TABLE IF NOT EXISTS exercises (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    description text,
    category text NOT NULL,
    muscle_group text,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
