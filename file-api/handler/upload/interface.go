package upload

import (
	"context"

	"bustanil.com/file-api/dto"
)

type Handler interface {
	HandleUpload(ctx context.Context, req *dto.PostFileMetadataRequest) (*dto.PostFileMetadataResponse, error)
}
