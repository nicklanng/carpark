package data

import (
	"database/sql"
	"fmt"
	"time"

	pq "github.com/lib/pq"

	"github.com/nicklanng/carpark/logging"
)

// OpenConnection creates a PostgreSQL connection with configured connection parameters
func OpenConnection(user, password, database, host string) (*sql.DB, *pq.Listener, error) {
	uri := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable", database, user, password, host)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(4)
	db.SetConnMaxLifetime(4 * time.Second)

	reportProblem := func(et pq.ListenerEventType, err error) {
		if err != nil {
			logging.Error(err.Error())
		}
	}
	listener := pq.NewListener(uri, 10*time.Second, time.Minute, reportProblem)
	if err = listener.Listen("new_event"); err != nil {
		return nil, nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, nil, err
	}

	if err = listener.Ping(); err != nil {
		return nil, nil, err
	}

	return db, listener, nil
}
