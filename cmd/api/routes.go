package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(
		http.MethodGet,
		"/v1/healthcheck",
		app.healthcheckHandler,
	)

	router.HandlerFunc(
		http.MethodPost,
		"/v1/users",
		app.registerUserHandler,
	)

	router.HandlerFunc(
		http.MethodPost,
		"/v1/tokens/authentication",
		app.createAuthenticationTokenHandler,
	)

	router.HandlerFunc(http.MethodGet,
		"/v1/exercises",
		app.requireAuthenticatedUser(app.listExercisesHandler),
	)

	router.HandlerFunc(http.MethodPost,
		"/v1/workouts",
		app.requireAuthenticatedUser(app.createWorkoutHandler),
	)
	router.HandlerFunc(http.MethodGet,
		"/v1/workouts",
		app.requireAuthenticatedUser(app.listWorkoutsHandler),
	)
	router.HandlerFunc(
		http.MethodPut,
		"/v1/workouts/:id",
		app.requireAuthenticatedUser(app.updateWorkoutHandler),
	)
	router.HandlerFunc(
		http.MethodGet,
		"/v1/workouts/:id",
		app.requireAuthenticatedUser(app.showWorkoutHandler),
	)
	router.HandlerFunc(
		http.MethodDelete,
		"/v1/workouts/:id",
		app.requireAuthenticatedUser(app.deleteWorkoutHandler),
	)
	router.HandlerFunc(
		http.MethodPost,
		"/v1/workouts/:id/schedule",
		app.requireAuthenticatedUser(app.scheduleWorkoutHandler),
	)

	return app.recoverPanic(app.authenticate(router))
}
