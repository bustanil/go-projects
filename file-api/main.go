package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"bustanil.com/file-api/api/upload"
	"bustanil.com/file-api/dao"
	"bustanil.com/file-api/db"
	"bustanil.com/file-api/external/aws/s3"
	uploadhandler "bustanil.com/file-api/handler/upload"
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

	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(awsProfileName))
	if err != nil {
		panic(err)
	}

	s3Client := s3.NewClient(&awsCfg, "sync-bucket")
	fileMetadataDao := dao.NewDao(pg)

	uploadAPI := upload.NewAPI(uploadhandler.NewHandler(s3Client, fileMetadataDao))

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
