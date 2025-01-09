package data

import (
	"database/sql"
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
