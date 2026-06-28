package store

import (
	"context"
	"database/sql"
	"errors"
)

var(
	ErrNotFound = errors.New("resource not found")
)

type Storage struct {
	Session interface{
		Create(context.Context) (*Session, error)
		FindBySessionToken(context.Context, string) (*Session, error)
	}

	Profile interface{
		Create(context.Context, string) error
	}
}

func NewStorage(db * sql.DB) Storage  {
	return Storage{
		Session: &SessionStorage{db},
		Profile: &ProfileStorage{db},
	}
}
