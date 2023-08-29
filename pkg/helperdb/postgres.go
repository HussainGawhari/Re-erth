package helperdb

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var DB *sql.DB

func ConnectToPostgres() (*sql.DB, error) {
	if err := gotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to PostgreSQL
	db, err := sql.Open("postgres", os.Getenv("POSTGRES"))
	if err != nil {
		log.Fatal(err)
	}

	// _, err = db.Exec(t)
	if err != nil {
		panic(err)
	}

	DB = db

	return db, nil
}
