package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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

func (p *Postgres) RunWithConn(ctx context.Context, f func(conn *sql.Conn)) {
	conn, err := p.DB.Conn(ctx)
	if err != nil {
		log.Printf("Error getting connection: %v", err)
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}(conn)

	f(conn)
}