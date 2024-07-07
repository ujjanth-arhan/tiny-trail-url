package repository

import (
	"github.com/ujjanth-arhan/tiny-trail-url/models/entity"
)

func GetByUrl(url string) ([]entity.Url, error) {
	return []entity.Url{}, nil
}

func InsertUrl(url entity.Url) error {
	return nil
}
