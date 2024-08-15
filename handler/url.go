package handler

import (
	_ "crypto/sha1"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ujjanth-arhan/tiny-trail-url/model/dto"
	"github.com/ujjanth-arhan/tiny-trail-url/model/entity"
	"github.com/ujjanth-arhan/tiny-trail-url/model/request"
	"github.com/ujjanth-arhan/tiny-trail-url/model/response"
	"github.com/ujjanth-arhan/tiny-trail-url/repository"
)

func HandleGetByShortUrl(w http.ResponseWriter, r *http.Request) {
	funcDetails := "Handler: HandleGetByShortUrl - "
	slog.Debug(funcDetails)

	shortUrl := r.PathValue("short_url")
	url, err := repository.GetByShortUrl(shortUrl)
	if err != nil {
		slog.Error(funcDetails + "Error getting short URL from repository: " + err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching short URL!"))
		return
	}

	if url.Id <= 0 {
		slog.Debug(funcDetails + "No data found for the given short URL: " + shortUrl)

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No data found for the given short URL!"))
		return
	}

	resUrl := response.Url{
		OriginalUrl:  url.OriginalUrl,
		ShortenedUrl: url.ShortenedUrl,
		Description:  url.Description,
		CreatedAt:    url.CreatedAt,
	}

	res, err := json.Marshal(resUrl)
	if err != nil {
		slog.Error(funcDetails + "Error marshalling JSON: " + err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching short URL!"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func HandleShortenUrl(w http.ResponseWriter, r *http.Request) {
	funcDetails := "Handler: Shorten URL - "
	slog.Debug(funcDetails)

	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error(funcDetails + "Error reading request body: " + err.Error())

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading request data!"))
		return
	}

	defer r.Body.Close()

	var reqUrl request.Url
	if marshalErr := json.Unmarshal(rBody, &reqUrl); marshalErr != nil {
		slog.Error(funcDetails + "Error unmarshalling request body: " + marshalErr.Error())

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error converting request body to JSON!"))
		return
	}

	reqUrl.OriginalUrl = strings.Trim(reqUrl.OriginalUrl, " ")

	if len(reqUrl.OriginalUrl) == 0 {
		slog.Debug(funcDetails + "Original URL is empty")

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Original URL cannot be empty!"))
		return
	}

	url, err := repository.GetByOriginalUrl(reqUrl.OriginalUrl)
	if err != nil {
		slog.Error(funcDetails + "Error getting original URL: " + err.Error())

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error converting request body to JSON!"))
		return
	}

	if url.Id > 0 {
		slog.Debug(funcDetails + "Shortened URL already exists: " + url.ShortenedUrl)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Shortened URL already exists: " + url.ShortenedUrl))
		return
	}

	uId := ""
	for attempt := 0; attempt < 10; attempt++ {
		uId = uuid.NewString()
		var checkUrl entity.Url
		if checkUrl, err = repository.GetByShortUrl(uId); err != nil {
			slog.Error(funcDetails + "Error getting short URL from repository: " + err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error checking for URL's existence!"))
			return
		}

		if checkUrl.Id <= 0 {
			break
		}
	}

	dtoUrl := dto.Url{
		OriginalUrl:  reqUrl.OriginalUrl,
		ShortenedUrl: uId,
		Description:  reqUrl.Description,
		CreatedAt:    time.Now().UTC(),
	}

	id, err := repository.InsertUrl(dtoUrl)
	if err != nil {
		slog.Error(funcDetails + "Error inserting URL data: " + err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error adding URL data!"))
		return
	}

	repoUrl, err := repository.GetById(id)
	if err != nil {
		slog.Error(funcDetails + "Error getting URL data by id: " + err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting URL data!"))
		return
	}

	resUrl := response.Url{
		OriginalUrl:  repoUrl.OriginalUrl,
		ShortenedUrl: repoUrl.ShortenedUrl,
		Description:  repoUrl.Description,
		CreatedAt:    repoUrl.CreatedAt,
	}

	res, err := json.Marshal(resUrl)
	if err != nil {
		slog.Error(funcDetails + "Error marshalling JSON: " + err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching short URL!"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
