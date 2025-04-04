package security

import (
	"github.com/zalando/go-keyring"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/config"
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

type Storager interface {
	GetKeyValue(key string) (string, error)
	SetKeyValue(key, value string) error
	DeleteKeyValue(key string) error
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
