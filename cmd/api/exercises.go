package main

import "net/http"

func (app *application) listExercisesHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	exercises, err := app.models.Exercises.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(
		w,
		http.StatusOK,
		envelope{"exercises": exercises},
		nil,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
