package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"bustanil.com/file-api/api/upload"
	"bustanil.com/file-api/db"
	panicmiddleware "bustanil.com/file-api/middleware/panic"
	"bustanil.com/file-api/router"
	"github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/lib/pq"
)

const awsProfileName = "sync"

var (
	pg *db.Postgres
)

func main() {
	fmt.Println("Starting server...")

	pg = db.InitDB("postgres", "postgres", "sync", "5432")

	if err := pg.DB.Ping(); err != nil {
		log.Panicf("Failed to connect to the database +%v", err)
		panic(err)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(awsProfileName))
	if err != nil {
		panic(err)
	}

	uploadAPI := upload.NewAPI(&cfg, pg)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: router.NewRouter().
			WithMiddlewares(
				&panicmiddleware.PanicLogger{},
			).
			POST("/api/file/metadata", uploadAPI.PostFileMetadata).
			Build(),
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
