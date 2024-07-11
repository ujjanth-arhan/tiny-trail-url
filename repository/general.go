package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)

func SetupRepositories() {
	query := url.Values{}
	query.Add("app name", "TinyTrail")
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("sa", "<YourStrong@Passw0rd>"),
		Host:   fmt.Sprintf("%s:%d", "localhost", 1433),
	}

	slog.Info("Query String:" + u.String())
	db, initErr := sql.Open("sqlserver", u.String())
	if initErr != nil {
		slog.Error("Unable to initialize DB connection query ", initErr)
	}

	if pingErr := db.PingContext(context.Background()); pingErr != nil {
		slog.Error("Failed to establish connection to database ", pingErr)
	}

	dbStmt, dbStmtErr := db.PrepareContext(context.Background(), "IF DB_ID('TinyTrail') IS NULL BEGIN CREATE DATABASE TinyTrail END ELSE PRINT 'DATABASE ALREADY EXISTS';")
	if dbStmtErr != nil {
		slog.Error("Failed to create database ", dbStmtErr)
	}

	_, execErr := dbStmt.ExecContext(context.Background())
	if execErr != nil {
		slog.Error("Failed to execute query ", execErr)
	}
	db.Close()

	query.Add("database", "TinyTrail")
	u.RawQuery = query.Encode()

	db, initErr = sql.Open("sqlserver", u.String())
	if initErr != nil {
		slog.Error("Unable to initialize TT DB connection query ", initErr)
	}

	if pingErr := db.PingContext(context.Background()); pingErr != nil {
		slog.Error("Failed to establish connection to TT database ", pingErr)
	}

	tblStmt, tblStmtErr := db.PrepareContext(context.Background(), "IF OBJECT_ID('dbo.Urls') IS NULL BEGIN CREATE TABLE dbo.Urls (ID INT PRIMARY KEY, OriginalUrl NVARCHAR, ShortenedUrl NVARCHAR, Description NVARCHAR, CreatedAt DATETIME) END ELSE PRINT 'TABLE ALREADY EXISTS';")
	if tblStmtErr != nil {
		slog.Error("Failed to create table ", tblStmtErr)
	}

	_, tblErr := tblStmt.ExecContext(context.Background())
	if tblErr != nil {
		slog.Error("Failed to execute query ", tblErr)
	}
}
