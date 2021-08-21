package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var database *sql.DB

// Db to create the singleton database connection
func Db() (*sql.DB, error) {
	if database != nil {
		return database, nil
	}
	if db, err := connect(); err != nil {
		return nil, err
	} else {
		database = db
		return database, nil
	}
}

// connect is to Connect to the database
func connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("connection to database failed: %w", err)
	}

	if err := db.Ping(); err != nil {
		defer func(db *sql.DB) {
			if err := db.Close(); err != nil {
				fmt.Printf("failed to close the connection: %v", err)
			}
		}(db)
		return nil, fmt.Errorf("can't sent ping to database: %w", err)
	}

	fmt.Println("Successfully connected to database!")
	return db, nil
}
