package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
	ErrDuplicateEmail = errors.New("duplicate email")
)

type Models struct {
	Users     UserModel
	Tokens    TokenModel
	Exercises ExerciseModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:     UserModel{DB: db},
		Tokens:    TokenModel{DB: db},
		Exercises: ExerciseModel{DB: db},
	}
}
