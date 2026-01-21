package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /summary", app.summary)

	standard := alice.New(app.recoverPanic, app.logRequest)

	return standard.Then(mux)
}
