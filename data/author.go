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

type Author struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Books     []Book    `json:"books,omitempty"`
}

func GetAuthors() ([]Author, error) {
	var authors []Author
	rows, err := database.Db.Query("SELECT * FROM authors")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var author Author
		if err := rows.Scan(&author.ID, &author.Name, &author.Bio, &author.CreatedAt, &author.UpdatedAt); err != nil {
			return nil, err
		}

		books, err := GetBooksByAuthor(author.ID)
		if err != nil {
			return nil, err
		}

		author.Books = books

		authors = append(authors, author)
	}

	return authors, nil
}

func GetAuthor(id int) (Author, error) {
	var author Author
	err := database.Db.QueryRow("SELECT * FROM authors WHERE id = ?", id).Scan(&author.ID, &author.Name, &author.Bio, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return author, err
	}

	return author, nil
}
