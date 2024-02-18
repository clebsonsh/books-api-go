package scripts

import (
	"fmt"
	"log"
	"os"

	"github.com/clebsonsh/books-api-go/database"
	"github.com/clebsonsh/books-api-go/utils"
)

func Seed() {
	utils.LoadEnv()
	database.Init()

	createSeedersTable()

	files, err := os.ReadDir("./database/seeders")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if seederAlreadyExecuted(file.Name()) {
			continue
		}

		sqlFile, sqlFileErr := os.ReadFile("./database/seeders/" + file.Name())
		if sqlFileErr != nil {
			log.Fatal(sqlFileErr)
		}

		fmt.Println("Seeding: ", file.Name())

		_, err := database.Db.Exec(string(sqlFile))
		if err != nil {
			panic(err)
		}

		saveSeeder(file.Name())
	}

	println("Database Seeded!")
}

func seederAlreadyExecuted(name string) bool {
	var count int

	err := database.Db.QueryRow("SELECT COUNT(*) FROM seeders WHERE name = ?", name).Scan(&count)
	if err != nil {
		panic(err)
	}

	return count > 0
}

func saveSeeder(name string) {
	_, err := database.Db.Exec("INSERT INTO seeders (name) VALUES (?)", name)
	if err != nil {
		panic(err)
	}
}

func createSeedersTable() {
	query := `CREATE TABLE IF NOT EXISTS seeders (
		id int primary key auto_increment,
		name varchar(255),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)`

	_, err := database.Db.Exec(query)
	if err != nil {
		panic(err)
	}
}
