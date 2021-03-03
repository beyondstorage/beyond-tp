package models

import (
	"errors"

	"github.com/dgraph-io/badger/v3"
)

type DBHandler struct {
	db *badger.DB
}

var DBCtxKey struct{}

var (
	// ErrNotFound returns when request record not found in db
	ErrNotFound = errors.New("record not found")
)

func NewDB(path string) (*DBHandler, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &DBHandler{db: db}, nil
}
