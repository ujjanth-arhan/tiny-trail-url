package main

import (
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/joho/godotenv"
	"github.com/ujjanth-arhan/tiny-trail-url/repository"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("Logger Initialized!")

	LoadEnvironmentVariables()
	RegisterRoutes()

	// Todo: Code cleanup of repo.setupdatabase
	repository.SetupDatabase()

	slog.Info("Starting server on port " + os.Getenv("PORT") + "!")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func LoadEnvironmentVariables() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading environment variables " + err.Error())
	}
}
