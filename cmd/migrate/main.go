package main

import (
	"wildberries/internal/database"
	"wildberries/internal/migrate"
)

func main() {
	db, err := database.SetConfig("config.json")
	if err != nil {
		panic(err)
	}
	db.Open()
	defer db.Close()
	if err = migrate.CreateSchema(db); err != nil {
		panic(err)
	}
}
