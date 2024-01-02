package database

import (
	"net/url"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

func InitDatabase() *sqlx.DB {
	var err error

	query := url.Values{}
	query.Add("database", os.Getenv("DB_DATABASE"))
	query.Add("connection+timeout", "30")

	dsn := &url.URL{
		Scheme:   os.Getenv("DB_SCHEME"),
		User:     url.UserPassword(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")),
		Host:     os.Getenv("DB_HOST"),
		RawQuery: query.Encode(),
	}

	db, err := sqlx.Connect("sqlserver", dsn.String())
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
