package data

import (
	"database/sql"
	"fmt"
	"time"

	// load but dont use postgres driver
	_ "github.com/lib/pq"
)

// OpenConnection creates a PostgreSQL connection with configured connection parameters
func OpenConnection(user, password, database, host string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(1 * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
