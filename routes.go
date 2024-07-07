package main

import (
	"net/http"

	"github.com/ujjanth-arhan/tiny-trail-url/handler"
)

func RegisterRoutes() {
	http.HandleFunc("POST /shortenurl", handler.HandleShortenUrl)

	http.HandleFunc("/{$}", handler.HandleHealthCheck)
	http.HandleFunc("/health", handler.HandleHealthCheck)
}
