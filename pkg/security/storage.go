package security

import (
	"context"
	"errors"
)

// ErrNotSupportedByKeyring indicates the operation is not supported by keyring-based storage.
var ErrNotSupportedByKeyring = errors.New("operation not supported by keyring storage backend")

// NewStorage function to create a new Storage instance with a keyring interface (for testing).
func NewStorage(serviceName string, kr Keyring) Storage {
	return Storage{
		ServiceName: serviceName,
		Keyring:     kr,
	}
}

// NewKeyRingStorage function to create a new Storage instance with RealKeyring.
func NewKeyRingStorage(serviceName string) Storage {
	return Storage{
		ServiceName: serviceName,
		Keyring:     RealKeyring{},
	}
}

// SetKeyValue method to set a key-value pair in the keyring.
func (s Storage) SetKeyValue(key, value string) error {
	return s.Keyring.Set(s.ServiceName, key, value)
}

// GetKeyValue method to get a value from the keyring by key.
func (s Storage) GetKeyValue(key string) (string, error) {
	return s.Keyring.Get(s.ServiceName, key)
}

// DeleteKeyValue method to delete a key-value pair from the keyring.
func (s Storage) DeleteKeyValue(key string) error {
	return s.Keyring.Delete(s.ServiceName, key)
}

// Remote storage methods for Storager interface.
// For keyring backend, these map to local operations or return errors for unsupported features.
// Context is ignored for keyring since it's a local synchronous operation.

// GetSecret retrieves all key-value pairs for this service (not supported for keyring).
func (s Storage) GetSecret(ctx context.Context, path string) (map[string]interface{}, error) {
	// Keyring stores individual keys, not hierarchical secrets
	return nil, ErrNotSupportedByKeyring
}

// SetSecret sets multiple key-value pairs (not supported for keyring).
func (s Storage) SetSecret(ctx context.Context, path string, data map[string]interface{}) error {
	// Keyring stores individual keys, not hierarchical secrets
	// Use SetSecretValue for individual keys instead
	return ErrNotSupportedByKeyring
}

// DeleteSecret deletes a secret at the given path (maps to local key deletion for keyring).
func (s Storage) DeleteSecret(ctx context.Context, path string) error {
	// For keyring, path is treated as a simple key name
	return s.DeleteKeyValue(path)
}

// GetSecretValue retrieves a specific value from a remote secret path.
// For keyring: combines path+key into a single local key (e.g., "path/key").
func (s Storage) GetSecretValue(ctx context.Context, path, key string) (string, error) {
	// Keyring doesn't support hierarchical paths, so we create a composite key
	compositeKey := path + "/" + key

	return s.GetKeyValue(compositeKey)
}

// SetSecretValue sets a specific key-value pair in a remote secret path.
// For keyring: combines path+key into a single local key (e.g., "path/key").
func (s Storage) SetSecretValue(ctx context.Context, path, key, value string) error {
	// Keyring doesn't support hierarchical paths, so we create a composite key
	compositeKey := path + "/" + key

	return s.SetKeyValue(compositeKey, value)
}
