package models

import (
	"errors"
	"fmt"
)

var (
	// ErrNotFound returns when request record not found in db
	ErrNotFound = errors.New("record not found")
)

// NewNotFoundErr wrap not found error with specific msg
func NewNotFoundErr(key string) error {
	return fmt.Errorf("with key %s: %w", key, ErrNotFound)
}
