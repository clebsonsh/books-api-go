package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/clebsonsh/books-api-go/sqlc"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	loadEnv()

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
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func main() {
	ctx := context.Background()
	queries := sqlc.New(db)

	author := sqlc.CreateAuthorParams{
		Name: "Clebson Moura",
		Bio:  sql.NullString{String: "Software Engineer", Valid: true},
	}

	authorFound, err := queries.AuthorExists(ctx, author.Name)

	if err != nil {
		fmt.Print(err)
	}

	if !authorFound {
		err = queries.CreateAuthor(ctx, author)

		if err != nil {
			fmt.Print(err)
		}
	}

	authors, err := queries.GetAuthors(ctx)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(authors)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
