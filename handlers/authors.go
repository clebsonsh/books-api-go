package handlers

import (
	"net/http"
	"strconv"

	"github.com/clebsonsh/books-api-go/data"
	"github.com/clebsonsh/books-api-go/utils"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := data.GetAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, authors, http.StatusOK)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	author, err := data.GetAuthor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondWithJSON(w, author, http.StatusOK)
}
