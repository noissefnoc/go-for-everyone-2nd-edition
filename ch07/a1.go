package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec(
		"INSERT INTO users(name, age) VALUES($name, $age)",
		sql.Named("name", "Bob"),
		sql.Named("age", 18))
	if err != nil {
		log.Fatal(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("affected: %d, lastInsertedID: %d\n", affected, lastInsertedID)
}
