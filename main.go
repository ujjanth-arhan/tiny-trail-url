package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("Logger Initialized!")

	LoadEnvironmentVariables()
	RegisterRoutes()
	// repository.SetupDatabase()

	slog.Info("Starting server on port " + os.Getenv("PORT") + "!")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func LoadEnvironmentVariables() {
	loadErr := godotenv.Load()
	if loadErr != nil {
		slog.Error("Error loading environment variables " + loadErr.Error())
	}
}
