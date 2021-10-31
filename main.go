package main

import (
	"briscolone/api"
	"briscolone/database"
	"briscolone/partita"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sqlx.Open("sqlite3", "file:database.db?mode=rw")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	s := api.Server{}
	adapter := database.NewSqliteItemsAdapter(db)
	s.Initialize(adapter)
	partita.Init(adapter)
	s.SetupRoutes()
	s.ServeStatic()
	s.Start()
}
