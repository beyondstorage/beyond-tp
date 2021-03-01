package model

import (
	"errors"

	"github.com/dgraph-io/badger/v3"
)

var DB *badger.DB

var (
	// ErrNotFound returns when request record not found in db
	ErrNotFound = errors.New("record not found")
	// ErrKeyRequired returns when request with blank key
	ErrKeyRequired = errors.New("key is required")
)

func Init(path string) error {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return err
	}
	DB = db
	return nil
}
