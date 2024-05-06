package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbParams := os.Getenv("DB_PARAMS")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", dbUsername, dbPassword, dbHost, dbPort, dbName, dbParams)

	// Define connection pool parameters (adjust as needed)
	maxOpenConns := 20
	maxIdleConns := 10

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	// Create connection pool
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		err = fmt.Errorf("failed to connect to db: %s", err)
		return nil, err
	}

	fmt.Println("successfully connected to db")
	return db, err

}
