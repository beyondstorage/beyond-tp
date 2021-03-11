package models

import (
	"context"
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

type DB struct {
	db *badger.DB
}

var contextKey struct{}

func NewDB(path string) (*DB, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func DbIntoContext(ctx context.Context, db *DB) context.Context {
	if ctx == nil {
		panic(fmt.Errorf("ctx is nil"))
	}
	if db == nil {
		panic(fmt.Errorf("db is nil"))
	}
	// If context already has db, we can return directly.
	_, ok := ctx.Value(contextKey).(*DB)
	if ok {
		return ctx
	}
	return context.WithValue(ctx, contextKey, db)
}

func DBFromContext(ctx context.Context) *DB {
	db, ok := ctx.Value(contextKey).(*DB)
	if !ok {
		panic(fmt.Errorf("db is not set"))
	}
	return db
}
