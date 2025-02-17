/*
Package Name: lib
File Name: database.go
Abstract: This file in the 'lib' package contains a method named 'GetDatabase'
that enables asynchronous connection to a PostgreSQL database. Using the 'pgxpool'
library, it establishes a database pool by constructing the connection URL from
environment variables. The method returns the database pool for seamless database interaction
and logs the connection status using a provided logger. It plays a crucial role in
facilitating database connectivity in the 'lib' package.
*/
package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ======== TYPES ========

// A type alias for the connection pool
type Database = pgxpool.Pool

// ======== METHODS ========

// GetDatabase returns a database pool to connect to the database asynchronously
func GetDatabase(logger Logger) *Database {

	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	// Create a connection pool to the database using pgxpool
	dbPool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		logger.Fatal("Unable to connect to database: ", err)
		os.Exit(1)
	}

	logger.Info("Connected to the database successfully.")
	// Closes the pool once the function goes out of scope.
	// defer dbPool.Close()

	return dbPool
}
