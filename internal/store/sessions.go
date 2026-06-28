package store

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"database/sql"
	"time"
)

const lenght = 16
const expiryOffset = 168

type SessionStorage struct {
	db *sql.DB
}
