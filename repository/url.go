package repository

var urls map[string]string = make(map[string]string)

func GetByUrl(url string) string {
	for k, v := range urls {
		if v == url {
			return k
		}
	}

	return ""
}

func GetByShortUrl(shortUrl string) string {
	return urls[shortUrl]
}

func InsertUrl(longUrl, shortUrl string) {
	urls[shortUrl] = longUrl
}

/*
func GetByUrl(url string) ([]entity.Url, error) {
	rrows, err := DB.QueryContext(context.Background(), `SELECT Id, ShortenedUrl, OriginalUrl FROM Url WHERE OriginalUrl = `+url)
	if err != nil {
		slog.Error("Failed to fetch URLs " + err.Error())
		return nil, err
	}

	rrows.Close()

	rows := make([]entity.Url, 0)
	for rrows.Next() {
		row := entity.Url{}
		if err = rrows.Scan(&row); err != nil {
			slog.Error("Failed to read URL row " + err.Error())
			return nil, err
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func InsertUrl(url entity.Url) (int64, error) {
	res, err := DB.ExecContext(context.Background(), "INSERT INTO dbo.Url VALUES(?, ?, ?, ?, ?)", url)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to insert value to URLs table %v ", url))
		return -1, err
	}

	rowId, _ := res.LastInsertId()
	return rowId, err
}
*/
