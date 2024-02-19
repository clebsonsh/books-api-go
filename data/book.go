package data

import (
	"time"

	"github.com/clebsonsh/books-api-go/database"
	"github.com/clebsonsh/books-api-go/utils"
)

func init() {
	utils.LoadEnv()
	database.Init()

}

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookWithAuthor struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    Author    `json:"author"`
}

func GetBooks() ([]BookWithAuthor, error) {
	var books []BookWithAuthor
	rows, err := database.Db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book BookWithAuthor
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}

		author, err := GetAuthor(book.AuthorID)
		if err != nil {
			return nil, err
		}

		book.Author = author

		books = append(books, book)
	}

	return books, nil
}

func GetBook(id int) (BookWithAuthor, error) {
	var book BookWithAuthor
	err := database.Db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.AuthorID, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return book, err
	}

	author, err := GetAuthor(book.AuthorID)
	if err != nil {
		return book, err
	}

	book.Author = author

	return book, nil
}

func GetBooksByAuthor(authorID int) ([]Book, error) {
	var books []Book
	rows, err := database.Db.Query("SELECT * FROM books WHERE author_id = ?", authorID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
