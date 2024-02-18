package scripts

import (
	"fmt"
	"log"
	"os"

	"github.com/clebsonsh/books-api-go/database"
	"github.com/clebsonsh/books-api-go/utils"
)

func Migrate(fresh bool) {
	utils.LoadEnv()
	database.Init()

	createMigrationsTable()

	if fresh {
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

		_, err := database.Db.Exec(string(sqlFile))
		if err != nil {
			panic(err)
		}

		saveMigration(file.Name())
	}

	println("Database migrated!")

}

func migrationAlreadyExecuted(name string) bool {
	var count int

	err := database.Db.QueryRow("SELECT COUNT(*) FROM migrations WHERE name = ?", name).Scan(&count)
	if err != nil {
		panic(err)
	}

	return count > 0
}

func saveMigration(name string) {
	_, err := database.Db.Exec("INSERT INTO migrations (name) VALUES (?)", name)
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

	_, err := database.Db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func refreshDatabase() {
	db_tables, dropErr := database.Db.Query("SHOW TABLES")
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
		_, dropErr := database.Db.Exec("DROP TABLE " + tables[i-1])
		if dropErr != nil {
			panic(dropErr)
		}
	}

	createMigrationsTable()
}
