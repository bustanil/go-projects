package entity

import "time"

type FileMetadata struct {
	ID        int64     `json:"id"`
	UUID      string    `json:"uuid"`
	Path      string    `json:"path"`
	Mimetype  string    `json:"mimetype"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
