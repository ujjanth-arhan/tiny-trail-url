package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("Logger Initialized!")

	RegisterRoutes()

	port := "8080"
	slog.Info("Starting server on port " + port + "!")
	http.ListenAndServe(":"+port, nil)
}
