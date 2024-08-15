package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/ujjanth-arhan/tiny-trail-url/model/dto"
	"github.com/ujjanth-arhan/tiny-trail-url/model/entity"
)

func GetById(id int) (entity.Url, error) {
	funcDetails := "Repository: GetById - "
	slog.Debug(funcDetails)

	query := `
		SELECT
			Id,
			OriginalUrl,
			ShortenedUrl,
			Description,
			CreatedAt
		FROM Urls WITH (NOLOCK)
		WHERE Id = @id
	`

	// Check for SQL injection and Change to prepared context
	rows, err := DB.QueryContext(context.Background(), query, sql.Named("id", id))
	if err != nil {
		slog.Error(funcDetails + "Error trying to query by id of url: " + err.Error())

		return entity.Url{}, err
	}

	defer rows.Close()

	var url entity.Url
	for rows.Next() {
		var (
			desc    sql.NullString
			crtedAt sql.NullTime
		)
		if err := rows.Scan(&url.Id, &url.OriginalUrl, &url.ShortenedUrl, &desc, &crtedAt); err != nil {
			slog.Error(funcDetails + "Error scaning rows: " + err.Error())

			return url, err
		}

		if desc.Valid {
			url.Description = desc.String
		}

		if crtedAt.Valid {
			url.CreatedAt = crtedAt.Time
		}
	}

	return url, nil
}

func GetByOriginalUrl(originalUrl string) (entity.Url, error) {
	funcDetails := "Repository: GetByOriginalUrl - "
	slog.Debug(funcDetails)

	query := `
		SELECT
			Id,
			OriginalUrl,
			ShortenedUrl,
			Description,
			CreatedAt
		FROM Urls WITH (NOLOCK)
		WHERE OriginalUrl = @originalUrl
	`

	// Check for SQL injection and Change to prepared context
	rows, err := DB.QueryContext(context.Background(), query, sql.Named("originalUrl", originalUrl))
	if err != nil {
		slog.Error(funcDetails + "Error trying to query by original url: " + err.Error())

		return entity.Url{}, err
	}

	defer rows.Close()

	var url entity.Url
	for rows.Next() {
		var (
			desc    sql.NullString
			crtedAt sql.NullTime
		)
		if err := rows.Scan(&url.Id, &url.OriginalUrl, &url.ShortenedUrl, &desc, &crtedAt); err != nil {
			slog.Error(funcDetails + "Error scaning rows: " + err.Error())

			return url, err
		}

		if desc.Valid {
			url.Description = desc.String
		}

		if crtedAt.Valid {
			url.CreatedAt = crtedAt.Time
		}
	}

	return url, nil
}

func GetByShortUrl(shortUrl string) (entity.Url, error) {
	funcDetails := "Repository: GetByShortUrl - "
	slog.Debug(funcDetails)

	query := `
		SELECT
			Id,
			OriginalUrl,
			ShortenedUrl,
			Description,
			CreatedAt
		FROM Urls WITH (NOLOCK)
		WHERE ShortenedUrl = @shortUrl
	`

	// Change to prepared context
	rows, err := DB.QueryContext(context.Background(), query, sql.Named("shortUrl", shortUrl))
	if err != nil {
		slog.Error(funcDetails + "Error trying to query by short url: " + err.Error())

		return entity.Url{}, err
	}

	defer rows.Close()

	var url entity.Url
	if rows.Next() {
		var (
			desc    sql.NullString
			crtedAt sql.NullTime
		)
		if err := rows.Scan(&url.Id, &url.OriginalUrl, &url.ShortenedUrl, &desc, &crtedAt); err != nil {
			slog.Error(funcDetails + "Error scaning rows: " + err.Error())

			return url, err
		}

		if desc.Valid {
			url.Description = desc.String
		}

		if crtedAt.Valid {
			url.CreatedAt = crtedAt.Time
		}
	}

	return url, nil
}

func InsertUrl(url dto.Url) (int, error) {
	funcDetails := "Repository: InsertUrl - "
	slog.Debug(funcDetails)

	query := `
		INSERT INTO Urls (
			OriginalUrl,
			ShortenedUrl,
			Description,
			CreatedAt
		)
		OUTPUT INSERTED.ID
		VALUES (
			@originalUrl,
			@shortenedUrl,
			@description,
			@createdAt
		)
	`

	var (
		err   error
		rows  *sql.Rows
		rowId int
	)

	if rows, err = DB.QueryContext(
		context.Background(),
		query,
		sql.Named("originalUrl", url.OriginalUrl),
		sql.Named("shortenedUrl", url.ShortenedUrl),
		sql.Named("description", url.Description),
		sql.Named("createdAt", url.CreatedAt)); err != nil {

		slog.Error(funcDetails + "Failed to INSERT URL: " + err.Error())

		return rowId, err
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&rowId)
	}

	return rowId, nil
}
