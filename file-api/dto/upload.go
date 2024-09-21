package dto

type PostFileMetadataRequest struct {
	Path     string `json:"path"`
	Mimetype string `json:"mimetype"`
	Size     int64  `json:"size"`
}

type PostFileMetadataResponse struct {
	PresignedURL     string              `json:"presigned_url"`
	PresignedHeaders map[string][]string `json:"presigned_header"`
}
