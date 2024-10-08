package handler

import (
	"log/slog"
	"net/http"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handler: Health check")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running!"))
}
