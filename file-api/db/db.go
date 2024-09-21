package db

import (
	"context"
	"database/sql"
	"fmt"
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

func (p *Postgres) RunWithConn(ctx context.Context, f func(conn *sql.Conn) error) error {
	conn, err := p.DB.Conn(ctx)
	if err != nil {
		log.Printf("Error getting connection: %v", err)
		return err
	}
	defer func(conn *sql.Conn) {
		if conn != nil {
			err := conn.Close()
			if err != nil {
				log.Printf("Error closing connection: %v", err)
			}
		}
	}(conn)

	err = f(conn)
	if err != nil {
		return err
	}

	return nil
}
