package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:     fmt.Sprintf(":%d", app.config.port),
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

	return srv.ListenAndServe()
}
