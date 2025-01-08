package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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
