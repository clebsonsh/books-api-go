package main

import (
	"os"

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

}
