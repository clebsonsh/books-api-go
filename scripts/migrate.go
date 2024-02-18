package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/clebsonsh/books-api-go/database"
	"github.com/clebsonsh/books-api-go/utils"
)

var db *sql.DB

func init() {
	utils.LoadEnv()
	database.Init()
	db = database.Db
	createMigrationsTable()
}

func main() {
	var shouldRefreshDatabase bool

	if len(os.Args) > 1 {
		shouldRefreshDatabase = os.Args[1] == "fresh"
	}

	if shouldRefreshDatabase {
		refreshDatabase()
	}

	files, err := os.ReadDir("./database/migrations")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {

		if migrationAlreadyExecuted(file.Name()) {
			continue
		}

		sqlFile, sqlFileErr := os.ReadFile("./database/migrations/" + file.Name())
		if sqlFileErr != nil {
			log.Fatal(sqlFileErr)
		}

		fmt.Println("Migrating: ", file.Name())

		_, err := db.Exec(string(sqlFile))

		if err != nil {
			panic(err)
		}

		saveMigration(file.Name())
	}

	println("Database migrated!")

}

func migrationAlreadyExecuted(name string) bool {
	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM migrations WHERE name = ?", name).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func saveMigration(name string) {
	_, err := db.Exec("INSERT INTO migrations (name) VALUES (?)", name)
	if err != nil {
		panic(err)
	}
}

func createMigrationsTable() {
	query := `CREATE TABLE IF NOT EXISTS migrations (
		id int primary key auto_increment,
		name varchar(255),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}
}

func refreshDatabase() {
	db_tables, dropErr := db.Query("SHOW TABLES")
	if dropErr != nil {
		panic(dropErr)
	}

	var tables []string

	for db_tables.Next() {
		var table string
		db_tables.Scan(&table)
		tables = append(tables, table)
	}

	for i := len(tables); i > 0; i-- {
		_, dropErr := db.Exec("DROP TABLE " + tables[i-1])
		if dropErr != nil {
			panic(dropErr)
		}
	}

	createMigrationsTable()
}
