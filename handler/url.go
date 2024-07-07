package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/ujjanth-arhan/tiny-trail-url/models/entity"
	"github.com/ujjanth-arhan/tiny-trail-url/models/request"
	"github.com/ujjanth-arhan/tiny-trail-url/repository"
)

func HandleShortenUrl(w http.ResponseWriter, r *http.Request) {
	rBody, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		slog.Debug("Error reading request body: " + readErr.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading request body"))
		return
	}

	defer r.Body.Close()

	reqUrl := request.RequestShortenUrl{}
	marshalErr := json.Unmarshal(rBody, &reqUrl)
	if marshalErr != nil {
		slog.Debug("Error marshalling request body: " + marshalErr.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error marshalling request body"))
		return
	}

	urls, fetchErr := repository.GetByUrl(reqUrl.Url)
	if fetchErr != nil {
		slog.Debug("Error fetching data from DB: " + fetchErr.Error())
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte("Error marshalling request body"))
		return
	}

	if len(urls) != 0 {
		slog.Debug("Shortened URL already exists: " + urls[0].ShortUrl)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Shortened URL already exists: " + urls[0].ShortUrl))
		return
	}

	shortUrl := shortenUrl(reqUrl.Url)
	insertErr := repository.InsertUrl(entity.Url{ShortUrl: shortUrl, LongUrl: reqUrl.Url})
	if insertErr != nil {
		slog.Debug("Error inserting data into DB: "+insertErr.Error(), "Object", insertErr)
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte("Error inserting data into DB"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", shortUrl)))
}

func shortenUrl(longUrl string) string {
	return ""
}
