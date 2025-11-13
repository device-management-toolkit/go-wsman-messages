package security

import (
	"context"

	"github.com/zalando/go-keyring"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/config"
)

type Cryptor interface {
	Decrypt(cipherText string) (string, error)
	Encrypt(plainText string) (string, error)
	EncryptWithKey(plainText, key string) (string, error)
	GenerateKey() string
	ReadAndDecryptFile(filePath string) (config.Configuration, error)
}

type Crypto struct {
	EncryptionKey string
}

// Storager interface for secret storage operations.
// Implementations can be simple (keyring) or complex (Vault, AWS Secrets Manager).
type Storager interface {
	// Simple key-value operations - works for both local keyring and remote storage
	GetKeyValue(key string) (string, error)
	SetKeyValue(key, value string) error
	DeleteKeyValue(key string) error

	// Remote secret storage operations - hierarchical path-based storage (Vault, AWS Secrets Manager)
	GetSecret(ctx context.Context, path string) (map[string]interface{}, error)
	SetSecret(ctx context.Context, path string, data map[string]interface{}) error
	DeleteSecret(ctx context.Context, path string) error
	GetSecretValue(ctx context.Context, path, key string) (string, error)
	SetSecretValue(ctx context.Context, path, key, value string) error
}

type Storage struct {
	ServiceName string
	Keyring     Keyring
}

// Keyring interface to abstract the keyring operations.
type Keyring interface {
	Set(serviceName, key, value string) error
	Get(serviceName, key string) (string, error)
	Delete(serviceName, key string) error
}

// RealKeyring struct to implement the Keyring interface using the real keyring package.
type RealKeyring struct{}

// Set method to set a key-value pair in the real keyring.
func (r RealKeyring) Set(serviceName, key, value string) error {
	return keyring.Set(serviceName, key, value)
}

// Get method to get a value from the real keyring by key.
func (r RealKeyring) Get(serviceName, key string) (string, error) {
	return keyring.Get(serviceName, key)
}

// Delete method to delete a key-value pair from the real keyring.
func (r RealKeyring) Delete(serviceName, key string) error {
	return keyring.Delete(serviceName, key)
}
