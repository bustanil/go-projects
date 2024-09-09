package main

import (
	"bustanil.com/file-api/db"
	"bustanil.com/file-api/entity"
	"bustanil.com/file-api/middleware"
	"bustanil.com/file-api/router"
	"database/sql"
	"encoding/json"
	"fmt"
	uuid "github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	pg *db.Postgres
)

func main() {
	fmt.Println("Starting server...")

	pg = db.InitDB("postgres", "postgres", "sync", "5432")

	server := http.Server{
		Addr: "localhost:8080",
		Handler: router.NewRouter().
			WithMiddlewares(
				&middleware.HTTPMiddleware{Counter: 0},
				&middleware.HTTPMiddleware{Counter: 1},
				&middleware.HTTPMiddleware{Counter: 2},
			).
			POST("/api/file/metadata", PostFileMetadata).
			Build(),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func PostFileMetadata(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing body: %v", err)
		}
	}(r.Body)

	m := entity.FileMetadata{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Printf("Error unmarshalling body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	m.Mimetype = "unknown"
	m.UUID = uuid.NewString()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	ctx := r.Context()
	pg.RunWithConn(ctx, func(conn *sql.Conn) {
		stmt, err := conn.PrepareContext(ctx, "INSERT INTO file_metadata (uuid, path, mimetype, size, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			log.Printf("Error preparing statement: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		_, err = stmt.ExecContext(ctx, m.UUID, m.Path, m.Mimetype, m.Size, m.CreatedAt, m.UpdatedAt)
		if err != nil {
			log.Printf("Error inserting: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

}
