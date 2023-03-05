package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-gorp/gorp"

	_ "github.com/lib/pq"
)

type Comment struct {
	Id      int64     `db:"id,primarykey,autoincrement"`
	Name    string    `db:"name,notnull,default:'名無し',size:200"`
	Text    string    `db:"text,notnull,size:400"`
	Created time.Time `db:"created,notnull"`
	Updated time.Time `db:"updated,notnull"`
}

func (c *Comment) PreInsert(s gorp.SqlExecutor) error {
	c.Created = time.Now()
	c.Updated = c.Created
	return nil
}

func (c *Comment) PreUpdate(s gorp.SqlExecutor) error {
	c.Updated = time.Now()
	return nil
}

func main() {
	/*
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
	*/

	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Comment{}, "comments").
		SetKeys(true, "id").
		AddIndex("idx_comments", "Btree", []string{"id"}).
		SetUnique(true)
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}
	err = dbmap.Insert(&Comment{Text: "こんにちわ"})
	if err != nil {
		log.Fatal(err)
	}

	dbmap.Insert(&Comment{Name: "bob", Text: "こんにちわ"})
	var comment Comment
	dbmap.SelectOne(&comment, "SELECT * FROM comments WHERE id = 1")

	comment.Text = "こんばんわ"
	dbmap.Update(&comment)

	var comments []Comment
	dbmap.Select(&comments, "SELECT * FROM comments WHERE name = $1", "bob")
}
