package kv

// KV is a simple key-value store interface.
type KV interface {
	// Get retrieves the value for the given key.
	// Returns `ErrNotFound` if the key does not exist.
	// Returns other errors for other failures.
	Get(key string) ([]byte, error)

	// Set sets the value for the given key.
	Set(key string, value []byte) error

	// Delete removes the key-value pair associated with the given key.
	// Does nothing if the key does not exist.
	Delete(key string) error

	// Clear removes all key-value pairs in the store.
	Clear() error

	// Close closes the key-value store and releases any resources.
	Close() error
}
