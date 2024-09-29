package dao

import (
	"context"
	"database/sql"
	"fmt"

	"bustanil.com/file-api/db"
	"bustanil.com/file-api/entity"
)

type FileMetadataDao interface {
	Save(ctx context.Context, md *entity.FileMetadata) error
}

type impl struct {
	pg *db.Postgres
}

func NewDao(pg *db.Postgres) FileMetadataDao {
	return &impl{
		pg: pg,
	}
}

func (i *impl) Save(ctx context.Context, m *entity.FileMetadata) error {
	err := i.pg.RunWithConn(ctx, func(conn *sql.Conn) error {
		stmt, err := conn.PrepareContext(ctx, "INSERT INTO file_metadata (uuid, path, mimetype, size, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, m.UUID, m.Path, m.Mimetype, m.Size, m.CreatedAt, m.UpdatedAt)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	return nil
}
