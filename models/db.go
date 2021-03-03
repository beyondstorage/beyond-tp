package models

import (
	"context"

	"github.com/dgraph-io/badger/v3"
	"github.com/gin-gonic/gin"
)

type DB struct {
	db *badger.DB
}

var dBCtxKey struct{}

func NewDB(path string) (*DB, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

// DbIntoContext inspired from `https://gqlgen.com/recipes/gin/#accessing-gincontext`
func DbIntoContext(h *DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), dBCtxKey, h)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// MustDBHandlerFrom inspired from `https://gqlgen.com/recipes/gin/#accessing-gincontext`
func MustDBHandlerFrom(ctx context.Context) *DB {
	v := ctx.Value(dBCtxKey)
	if v == nil {
		panic("could not retrieve DBPath")
	}

	return v.(*DB)
}
