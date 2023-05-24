package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//From main func will call this function to connect the database
func Connect() (*sql.DB) {
	db, err := sql.Open("postgres", "postgresql://max:roach@localhost:26257/socialmedia?sslmode=require")
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}
