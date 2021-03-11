package models

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/gin-gonic/gin"
)

type DB struct {
	db *badger.DB
}

const dbGinCtxKey = "db_in_gin"

func NewDB(path string) (*DB, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

// DbIntoGin inspired from `https://gqlgen.com/recipes/gin/#accessing-gincontext`
func DbIntoGin(db *DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(dbGinCtxKey, db)
		c.Next()
	}
}

// DBFromGin inspired from `https://gqlgen.com/recipes/gin/#accessing-gincontext`
func DBFromGin(c *gin.Context) *DB {
	return c.MustGet(dbGinCtxKey).(*DB)
}
