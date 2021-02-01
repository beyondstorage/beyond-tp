package utils

import "github.com/google/uuid"

// NewUUID generates an UUID string.
func NewUUID() string {
	return uuid.New().String()
}
