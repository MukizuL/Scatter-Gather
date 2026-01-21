package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

type application struct {
	logger *zap.Logger
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("error initializing logger: %s\n", err.Error())
		os.Exit(1)
	}

	defer func() { _ = logger.Sync() }()

	app := &application{logger: logger}

	stdErrLogger, err := zap.NewStdLogAt(logger, zap.ErrorLevel)
	if err != nil {
		fmt.Printf("error converting zap.Logger to log.Logger: %s\n", err.Error())
		os.Exit(1)
	}

	addr := getEnv("ADDR", ":8080")

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		ErrorLog:     stdErrLogger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting http server on", zap.String("addr", addr))

	err = srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		logger.Error("error running http server", zap.Error(err))
		os.Exit(1)
	}
}
