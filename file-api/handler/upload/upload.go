package upload

import (
	"context"
	"time"

	"bustanil.com/file-api/dao"
	"bustanil.com/file-api/dto"
	"bustanil.com/file-api/entity"
	"bustanil.com/file-api/external/aws/s3"
	"github.com/google/uuid"
)

type impl struct {
	s3Client *s3.Client
	dao      dao.FileMetadataDao
}

var (
	bucketName = "sync-bucket"
)

func NewHandler(s3Client *s3.Client, dao dao.FileMetadataDao) Handler {
	return &impl{
		s3Client: s3Client,
		dao:      dao,
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

	err := i.dao.Save(ctx, &m)
	if err != nil {
		return nil, err
	}

	signedURL, signedHeader, err := i.s3Client.PresignPutObject(ctx, req.Path)
	if err != nil {
		return nil, err
	}

	mdResp := &dto.PostFileMetadataResponse{
		PresignedURL:     signedURL,
		PresignedHeaders: signedHeader,
	}

	return mdResp, nil
}
