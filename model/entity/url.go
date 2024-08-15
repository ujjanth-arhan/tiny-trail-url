package entity

import "time"

type Url struct {
	Id           int
	OriginalUrl  string
	ShortenedUrl string
	Description  string
	CreatedAt    time.Time
}
