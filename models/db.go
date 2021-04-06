package models

import (
	"github.com/dgraph-io/badger/v3"
)

type DB struct {
	db *badger.DB
}

func NewDB(path string) (*DB, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}
