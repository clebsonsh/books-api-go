package main

import (
	"github.com/clebsonsh/books-api-go/utils"
)

func init() {
	utils.LoadEnv()
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
