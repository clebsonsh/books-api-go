package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/clebsonsh/books-api-go/data"
	"github.com/clebsonsh/books-api-go/database"
	"github.com/clebsonsh/books-api-go/scripts"
	"github.com/clebsonsh/books-api-go/utils"
)

func init() {
	utils.LoadEnv()
	database.Init()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate:fresh:seed":
			scripts.Migrate(true)
			scripts.Seed()
		case "migrate:fresh":
			scripts.Migrate(true)
		case "migrate":
			scripts.Migrate(false)
		case "seed":
			scripts.Seed()
		}
	}
}

func main() {
	serverUriAndPort := os.Getenv("HOST_URI") + ":" + os.Getenv("HOST_PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/books/", func(w http.ResponseWriter, r *http.Request) {
		books, err := data.GetBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.RespondWithJSON(w, books, http.StatusOK)
	})

	mux.HandleFunc("GET /api/v1/authors/", func(w http.ResponseWriter, r *http.Request) {
		authors, err := data.GetAuthors()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.RespondWithJSON(w, authors, http.StatusOK)
	})

	fmt.Printf("Server running at http://%s\n", serverUriAndPort)

	http.ListenAndServe(serverUriAndPort, mux)
}
