package upload

import (
	"context"
	"database/sql"
	"time"

	"bustanil.com/file-api/db"
	"bustanil.com/file-api/dto"
	"bustanil.com/file-api/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type impl struct {
	awsConfig *aws.Config
	pg        *db.Postgres
}

var (
	bucketName = "sync-bucket"
)

func NewHandler(cfg *aws.Config, pg *db.Postgres) Handler {
	return &impl{
		awsConfig: cfg,
		pg:        pg,
	}
}

func (i *impl) HandleUpload(ctx context.Context, req *dto.PostFileMetadataRequest) (*dto.PostFileMetadataResponse, error) {
	m := entity.FileMetadata{
		UUID:      uuid.NewString(),
		Path:      req.Path,
		Mimetype:  "unknown",
		Size:      req.Size,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := i.pg.RunWithConn(ctx, func(conn *sql.Conn) error {
		stmt, err := conn.PrepareContext(ctx, "INSERT INTO file_metadata (uuid, path, mimetype, size, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			return err
		}

		_, err = stmt.ExecContext(ctx, m.UUID, m.Path, m.Mimetype, m.Size, m.CreatedAt, m.UpdatedAt)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	c := s3.NewFromConfig(*i.awsConfig)
	presignClient := s3.NewPresignClient(c)
	signedURL, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &m.Path,
	})
	if err != nil {
		return nil, err
	}

	mdResp := &dto.PostFileMetadataResponse{
		PresignedURL:     signedURL.URL,
		PresignedHeaders: signedURL.SignedHeader,
	}

	return mdResp, nil
}
