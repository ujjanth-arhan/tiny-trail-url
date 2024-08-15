package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

var DB *sql.DB

func SetupDatabase() {
	funcDetails := "Setup Database - "
	slog.Info("Setting up database...")

	createDatabaseAndUser()
	createTable()

	if err := DB.PingContext(context.Background()); err != nil {
		slog.Error(funcDetails + "Failed to connect to DB: " + err.Error())
	}

}

func createTable() {
	funcDetails := "Setup Database - createTable "

	conProp := &url.URL{
		Scheme:   os.Getenv("DB_SCHEME"),
		User:     url.UserPassword(os.Getenv("DB_SA_NAME"), os.Getenv("DB_SA_PASSWORD")),
		Host:     fmt.Sprintf("%s:%s", os.Getenv("HOST_NAME"), os.Getenv("DB_PORT")),
		RawQuery: url.Values(map[string][]string{"database": {os.Getenv("DB_NAME")}}).Encode(),
	}

	db, err := sql.Open("sqlserver", conProp.String())
	if err != nil {
		slog.Error(funcDetails + "Unable to initialize " + os.Getenv("DB_NAME") + " DB connection query " + err.Error())
	}

	query := `
		IF OBJECT_ID('dbo.Urls') IS NULL
		BEGIN
			CREATE TABLE dbo.Urls
			(
				ID INT PRIMARY KEY IDENTITY(1, 1),
				OriginalUrl NVARCHAR(MAX),
				ShortenedUrl NVARCHAR(MAX),
				Description NVARCHAR(MAX),
				CreatedAt DATETIME NOT NULL
			)
		END
	`
	tblStmt, err := db.PrepareContext(context.Background(), query)
	if err != nil {
		slog.Error(funcDetails + "Failed to prepare Urls table context " + err.Error())
	}

	if _, err := tblStmt.ExecContext(context.Background()); err != nil {
		slog.Error(funcDetails + "Failed to execute Urls table creation statement " + err.Error())
	}

	DB = db
}

func createDatabaseAndUser() {
	funcDetails := "Setup Database - createDatabaseAndUser "

	conProp := &url.URL{
		Scheme: os.Getenv("DB_SCHEME"),
		User:   url.UserPassword(os.Getenv("DB_SA_NAME"), os.Getenv("DB_SA_PASSWORD")),
		Host:   fmt.Sprintf("%s:%s", os.Getenv("HOST_NAME"), os.Getenv("DB_PORT")),
	}

	db, err := sql.Open(os.Getenv("DB_DRIVER_NAME"), conProp.String())
	if err != nil {
		slog.Error(funcDetails + "Unable to initialize DB connection query " + err.Error())
	}

	/*
	* SQL Server does not support parameters for create database query
	 */
	query := fmt.Sprintf("IF DB_ID('%s') IS NULL BEGIN CREATE DATABASE %s END;", os.Getenv("DB_NAME"), os.Getenv("DB_NAME"))
	if _, err := db.ExecContext(context.Background(), query); err != nil {
		slog.Error(funcDetails + "Failed to execute query " + err.Error())
	}

	db.Close()
}
