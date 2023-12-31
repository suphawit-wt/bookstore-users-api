package database

import (
	"fmt"
	"net/url"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

var DB *sqlx.DB

func InitDatabase() {
	var err error

	query := url.Values{}
	query.Add("database", os.Getenv("DB_DATABASE"))
	query.Add("connection+timeout", "30")

	dsn := &url.URL{
		Scheme:   os.Getenv("DB_SCHEME"),
		User:     url.UserPassword(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")),
		Host:     os.Getenv("DB_HOST"),
		Path:     os.Getenv("DB_INSTANCE"),
		RawQuery: query.Encode(),
	}

	fmt.Println(dsn.String())

	DB, err = sqlx.Connect("sqlserver", dsn.String())
	if err != nil {
		panic(err)
	}
}
