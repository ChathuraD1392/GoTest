package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ChathuraD1392/gotest/internal/models/sqlite"
	_ "modernc.org/sqlite"
)

type app struct {
	posts *sqlite.PostModels
}

func main() {

	db, err := sql.Open("sqlite", "./app.db")

	if err != nil {
		log.Fatal(err)
	}

	app := app{
		posts: &sqlite.PostModels{
			DB: db,
		},
	}

	// create a server
	server := http.Server{
		Addr:    ":8000",
		Handler: app.routes(),
	}

	log.Println("Listning through :8000")
	server.ListenAndServe()
}
