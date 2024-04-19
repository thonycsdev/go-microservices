package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func MakeConn() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=local_password dbname=local_db sslmode=disable")
	if err != nil {
		return nil, err
	}

	return conn, nil

}
