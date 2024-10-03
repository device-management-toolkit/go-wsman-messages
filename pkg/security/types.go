package security

import (
	"github.com/99designs/keyring"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/config"
)

type Cryptor interface {
	Decrypt(cipherText string, key []byte) ([]byte, error)
	Encrypt(plainText []byte, key string) (string, error)
	GenerateKey() string
	ReadAndDecryptFile(filePath string, key []byte) (config.Configuration, error)
}

type Crypto struct{}

type Storager interface {
	GetKeyValue(key string) (string, error)
	SetKeyValue(key, value string) error
}

type Storage struct {
	ServiceName string
	Keyring     Keyring
}

// Keyring interface to abstract the keyring operations.
type Keyring interface {
	Set(serviceName string, item keyring.Item) error
	Get(serviceName, key string) (keyring.Item, error)
}

// RealKeyring struct to implement the Keyring interface using the real keyring package.
type RealKeyring struct{}
