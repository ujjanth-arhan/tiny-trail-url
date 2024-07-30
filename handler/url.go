package handler

import (
	_ "crypto/sha1"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/ujjanth-arhan/tiny-trail-url/models/request"
	"github.com/ujjanth-arhan/tiny-trail-url/repository"
)

func HandleFetchByShortUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("short_url")
	longUrl := repository.GetByShortUrl(shortUrl)
	if len(longUrl) > 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(longUrl))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Requested short URL does not exist!"))
}

func HandleShortenUrl(w http.ResponseWriter, r *http.Request) {
	rBody, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		slog.Debug("Error reading request body: " + readErr.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading request body"))
		return
	}

	defer r.Body.Close()

	var reqUrl request.Url
	if marshalErr := json.Unmarshal(rBody, &reqUrl); marshalErr != nil {
		slog.Debug("Error unmarshalling request body: " + marshalErr.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error unmarshalling request body"))
		return
	}

	reqUrl.Url = strings.Trim(reqUrl.Url, " ")
	shortUrl := repository.GetByUrl(reqUrl.Url)
	if len(strings.Trim(shortUrl, " ")) != 0 {
		slog.Debug("Shortened URL already exists: " + shortUrl)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Shortened URL already exists: " + shortUrl))
		return
	}

	shortUrl = shortenUrl()
	repository.InsertUrl(reqUrl.Url, shortUrl)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(shortUrl))
}

func shortenUrl() string {
	uId := uuid.NewString()
	url := repository.GetByShortUrl(uId)
	for len(url) > 0 {
		uId = uuid.NewString()
		url = repository.GetByShortUrl(uId)
	}

	return uId
}
