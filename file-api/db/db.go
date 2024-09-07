package db

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	DB *sql.DB
}

func InitDB(username, password, dbname, port string) *Postgres {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", username, password, port, dbname))
	if err != nil {
		panic(err)
	}

	return &Postgres{
		DB: db,
	}
}
