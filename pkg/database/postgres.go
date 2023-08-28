package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

func ConnectToPostgres() (*sql.DB, error) {
	if err := gotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to PostgreSQL
	db, err := sql.Open("postgres", os.Getenv("POSTGRES"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Defer closing the database connection

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createAddressTableQuery := `
		CREATE TABLE IF NOT EXISTS addresses (
			id SERIAL PRIMARY KEY,
			street VARCHAR(25),
			postal_code VARCHAR(20),
			city VARCHAR(25),
			country VARCHAR(25)
		);
	`

	_, err = db.Exec(createAddressTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	createClientTableQuery := `
		CREATE TABLE IF NOT EXISTS clients (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(25),
			last_name VARCHAR(25),
			telephone VARCHAR(20),
			email VARCHAR(55),
			address_id INT REFERENCES addresses(id),
			is_active BOOLEAN,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		);
	`

	_, err = db.Exec(createClientTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Address and Clients tables created successfully!")

	// fmt.Println("Connected to PostgreSQL!")

	return db, nil
}
