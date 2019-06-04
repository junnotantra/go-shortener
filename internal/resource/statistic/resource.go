package statistic

import (
	bolt "go.etcd.io/bbolt"
)

type (
	// Resource struct act as function receiver for resource level on shortener
	Resource struct {
		DB *bolt.DB
	}
)

// New function return a pointer to shortener's resource struct
func New(db *bolt.DB) *Resource {
	r := &Resource{
		DB: db,
	}

	return r
}
