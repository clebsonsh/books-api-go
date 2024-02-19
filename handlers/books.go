package handlers

import (
	"net/http"
	"strconv"

	"github.com/clebsonsh/books-api-go/data"
	"github.com/clebsonsh/books-api-go/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := data.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, books, http.StatusOK)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := data.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondWithJSON(w, book, http.StatusOK)
}
