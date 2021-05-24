package models

import (
	"errors"
	"fmt"
)

var (
	// ErrNotFound returns when request record not found in db
	ErrNotFound = errors.New("record not found")
	// ErrAlreadyExists returns when insert record with given key already exists
	ErrAlreadyExists = errors.New("record already exists")
)

type Error struct {
	op  string
	err error

	key string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s %s: %s", e.op, e.key, e.err)
}

func (e Error) Unwrap() error {
	return e.err
}

// NewNotFoundErr wrap not found error with specific msg
func NewNotFoundErr(key string) error {
	return Error{
		op:  "get",
		err: ErrNotFound,
		key: key,
	}
}
