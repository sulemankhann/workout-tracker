package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	portStr, exist := os.LookupEnv("PORT")
	if !exist {
		portStr = "4000"
	}

	env, exist := os.LookupEnv("ENV")
	if !exist {
		env = "development"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		// invalid port, setting default port
		port = 4000
	}

	dbDSN, exist := os.LookupEnv("DB_DSN")
	if !exist {
		logger.Error("DB DSN not set")
		os.Exit(1)
	}

	db, err := openDB(dbDSN)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	logger.Info("database connection pool established")

	cfg := config{
		port: port,
		env:  env,
	}

	app := application{
		config: cfg,
		logger: logger,
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
