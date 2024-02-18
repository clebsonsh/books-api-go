package main

import (
	"github.com/clebsonsh/books-api-go/database"
	"github.com/clebsonsh/books-api-go/utils"
)

func init() {
	utils.LoadEnv()
	database.Init()
}

func main() {

}
