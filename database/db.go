package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Init() {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_DATABASE"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
}
