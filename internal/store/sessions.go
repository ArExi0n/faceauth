package store

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"time"
)

const length = 16
const expiryOffset = 168

type SessionStorage struct {
	db *sql.DB
}

type Session struct {
	ID        int32     `json:"id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func createSessionToken() (string, error) {
	b := make([]byte, length)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func (ts *SessionStorage) Create(ctx context.Context) (*Session, error) {
	query := `
		INSERT INTO sessions (token, expires_at)
		VALUES ($1, $2)
	`

	sessionToken, err := createSessionToken()
	if err != nil {
		return nil, err
	}

	expiryDate := time.Now().Add(expiryOffset * time.Hour)

	err = ts.db.QueryRowContext(ctx, query, sessionToken, expiryDate).Err()
	if err != nil {
		return nil, err
	}

	return &Session{
		Token:     sessionToken,
		ExpiresAt: expiryDate,
	}, nil
}

func (ts *SessionStorage) FindBySessionToken(ctx context.Context, sessionToken string) (*Session, error) {
	query := `
		SELECT id, token, created_at, expires_at FROM sessions WHERE token = $1
	`

	var session Session

	err := ts.db.QueryRowContext(ctx, query, sessionToken).Scan(
		&session.ID,
		&session.Token,
		&session.CreatedAt,
		&session.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return &session, nil
}
