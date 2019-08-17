package main

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type Validator struct {
	validator *validator.Validate
}

type Comment struct {
	Id int64 `json:"id" db:"id,primarykey,autoincrement"`
	Name string `json:"name" db:"name,notnull,default:'名無し',size:200"`
	Text string `json:"text" db:"text,notnull,size:399"`
	Created time.Time `json:"created" db:"created,notnull"`
	Updated time.Time `json:"updated" db:"updated,notnull"`
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

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {
	db, err := sql.Open("postgres", "sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Comment{}, "comments")

	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}

	e.Static("/", "./static")

	e.GET("/api/comments", func(c echo.Context) error {
		var comments []Comment
		_, err := dbmap.Select(&comments,
			"SELECT * FROM comments ORDER BY created desc LIMIT 10")
		if err != nil {
			c.Logger().Error("Select: ", err)
			return c.String(http.StatusBadRequest, "Select: " + err.Error())
		}
		return c.JSON(http.StatusOK, comments)
	})

	e.POST("/api/comments", func(c echo.Context) error {
		var comment Comment
		if err := c.Bind(&comment); err != nil {
			c.Logger().Error("Bind: ", err)
			return c.String(http.StatusBadRequest, "Bind: " + err.Error())
		}
		if err := c.Validate(&comment); err != nil {
			c.Logger().Error("Validator: ", err)
			return c.String(http.StatusBadRequest, "Validate: " + err.Error())
		}
		if err := dbmap.Insert(&comment); err != nil {
			c.Logger().Error("Insert: ", err)
			return c.String(http.StatusBadRequest, "Insert: " + err.Error())
		}
		c.Logger().Info("ADDED: %v", comment.Id)
		return c.JSON(http.StatusCreated, "")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
