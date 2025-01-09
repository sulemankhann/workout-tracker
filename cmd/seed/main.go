package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Seeder struct {
	DB *sql.DB
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbDSN, exist := os.LookupEnv("DB_DSN")
	if !exist {
		log.Fatal("DB DSN not set")
	}

	db, err := sql.Open("postgres", dbDSN)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal(err.Error())
	}

	seeder := Seeder{DB: db}

	seeder.SeedExercises()
}
