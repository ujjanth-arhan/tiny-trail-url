package dto

import "time"

type Url struct {
	OriginalUrl  string
	ShortenedUrl string
	Description  string
	CreatedAt    time.Time
}
