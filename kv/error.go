package kv

import "errors"

// ErrKeyNotFound is returned when a requested key does not exist in the store.
var ErrKeyNotFound = errors.New("key not found")
