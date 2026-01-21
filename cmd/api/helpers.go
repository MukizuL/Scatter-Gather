package main

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type envelope map[string]any

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), zap.String("method", method), zap.String("uri", uri))

	_ = app.writeJSON(w, http.StatusInternalServerError, envelope{"message": http.StatusText(http.StatusInternalServerError)})
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	w.Write(js)

	return nil
}
