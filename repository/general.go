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

var db *sql.DB

func SetupDatabase() {
	query := url.Values{}
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("sa", os.Getenv("DB_SA_PASSWORD")),
		Host:   fmt.Sprintf("%s:%s", os.Getenv("HOST_NAME"), os.Getenv("DB_PORT")),
	}

	slog.Info("Query String:" + u.String())
	dbConnection, initErr := sql.Open("sqlserver", u.String())
	if initErr != nil {
		slog.Error("Unable to initialize DB connection query ", initErr)
	}

	/*
	* SQL Server does not support parameters for create database query
	 */
	_, execErr := dbConnection.ExecContext(context.Background(), fmt.Sprintf("IF DB_ID('%s') IS NULL BEGIN CREATE DATABASE %s END ELSE PRINT 'DATABASE ALREADY EXISTS';", os.Getenv("DB_NAME"), os.Getenv("DB_NAME")))
	if execErr != nil {
		slog.Error("Failed to execute query ", execErr)
	}

	dbConnection.Close()

	query.Add("database", os.Getenv("DB_NAME"))
	u.RawQuery = query.Encode()

	dbConnection, initErr = sql.Open("sqlserver", u.String())
	if initErr != nil {
		slog.Error("Unable to initialize TT DB connection query ", initErr)
	}

	tblStmt, tblStmtErr := dbConnection.PrepareContext(context.Background(), "IF OBJECT_ID('dbo.Urls') IS NULL BEGIN CREATE TABLE dbo.Urls (ID INT PRIMARY KEY, OriginalUrl NVARCHAR, ShortenedUrl NVARCHAR, Description NVARCHAR, CreatedAt DATETIME) END ELSE PRINT 'TABLE ALREADY EXISTS';")
	if tblStmtErr != nil {
		slog.Error("Failed to create table ", tblStmtErr)
	}

	_, tblErr := tblStmt.ExecContext(context.Background())
	if tblErr != nil {
		slog.Error("Failed to execute query ", tblErr)
	}

	db = dbConnection
}

func PingDB() error {
	return db.PingContext(context.Background())
}
