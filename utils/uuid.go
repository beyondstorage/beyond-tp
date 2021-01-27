package utils

import "github.com/satori/go.uuid"

// NewUUID generates an UUID string.
func NewUUID() string {
	return uuid.NewV4().String()
}
