/*********************************************************************
 * Copyright (c) Intel Corporation 2025
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

const (
	// generatedKeyLen is the raw byte length of keys from GenerateKey (256-bit).
	generatedKeyLen = 32
	// generatedKeyEncodedLen is the base64-encoded length of a generated key.
	generatedKeyEncodedLen = 44
)

// keyBytes returns the AES key material for a stored key string. Keys from
// GenerateKey are base64-encoded 32-byte values (44 characters); anything
// else is a legacy key whose raw string bytes are the key. A working legacy
// key is 16/24/32 characters (a valid AES key size), so no legacy key can be
// 44 characters, and 32 base64 characters decode to 24 bytes, never 32 — the
// two formats cannot be confused.
func keyBytes(key string) []byte {
	if len(key) == generatedKeyEncodedLen {
		if decoded, err := base64.StdEncoding.DecodeString(key); err == nil && len(decoded) == generatedKeyLen {
			return decoded
		}
	}

	return []byte(key)
}

// Encrypt encrypts a string.
func (c Crypto) Encrypt(plainText string) (string, error) {
	return c.EncryptWithKey(plainText, c.EncryptionKey)
}

// EncryptWithKey encrypts a string with the provided key.
func (c Crypto) EncryptWithKey(plainText, key string) (string, error) {
	block, err := aes.NewCipher(keyBytes(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// GenerateKey returns a new 256-bit AES key as a 44-character base64 string.
// Encrypt/Decrypt recognize this format and decode it before use; keys in any
// other format are treated as legacy raw-byte keys (see keyBytes).
func (c Crypto) GenerateKey() string {
	key := make([]byte, generatedKeyLen)

	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(key)
}
