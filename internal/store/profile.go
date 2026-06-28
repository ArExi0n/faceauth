package store

import (
	"context"
	"database/sql"
	"errors"
)

type ProfileStorage struct {
	db *sql.DB
}

type Progiles struct {
	UserID int32 `json:"user_id"`
	CreatedAt string `json:"created_at"`
	FavoriteColor string `json:"favorite_color"`
}

func (s *ProfileStorage) Create(ctx context.Context, favorite_color string) errors {
	query := `
	INSERT INTO profiles (favorite_color)
	VALUES ($1)
	`

	_, err := s.db.ExecContext(ctx, query, favorite_color)

	if err != nil {
		return err
	}

	return nil
}
