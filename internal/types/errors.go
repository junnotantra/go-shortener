package types

import "errors"

var (
	// ErrServerError is general server error
	ErrServerError = errors.New("Server experiencing problem")
	// ErrNotFound is permission error
	ErrNotFound = errors.New("Could not found item")
	// ErrUnauthorized is permission error
	ErrUnauthorized = errors.New("Unauthorized action")
)
