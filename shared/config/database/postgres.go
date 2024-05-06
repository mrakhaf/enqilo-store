package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	errConn := db.Ping()
	if errConn != nil {
		err = fmt.Errorf("failed to connect to db: %s", err)
		return nil, err
	}

	fmt.Println("successfully connected to db")
	return db, err

}
