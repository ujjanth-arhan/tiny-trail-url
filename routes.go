package main

import (
	"log/slog"
	"net/http"

	"github.com/ujjanth-arhan/tiny-trail-url/handler"
)

func RegisterRoutes() {
	slog.Info("Registering routes...")

	http.HandleFunc("POST /shortenurl", handler.HandleShortenUrl)
	http.HandleFunc("GET /shorturl/{short_url}", handler.HandleGetByShortUrl)

	http.HandleFunc("/{$}", handler.HandleHealthCheck)
	http.HandleFunc("/health", handler.HandleHealthCheck)
}
