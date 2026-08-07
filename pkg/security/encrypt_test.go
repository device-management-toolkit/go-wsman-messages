/*********************************************************************
 * Copyright (c) Intel Corporation 2025
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package security

import (
	"crypto/aes"
	"encoding/base64"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		message       string
		key           string
		expectedError expectedError
		errorMsg      error
		expected      interface{}
	}{
		{
			name:          "successful encryption",
			message:       "test message",
			key:           validKey,
			expectedError: expectedError{},
			errorMsg:      nil,
			expected:      "test message",
		},
		{
			name:          "successful encryption with legacy 16-char key",
			message:       "test message",
			key:           "0123456789abcdef",
			expectedError: expectedError{},
			errorMsg:      nil,
			expected:      "test message",
		},
		{
			name:          "successful encryption with legacy 24-char key",
			message:       "test message",
			key:           "0123456789abcdef01234567",
			expectedError: expectedError{},
			errorMsg:      nil,
			expected:      "test message",
		},
		{
			name:          "key too short",
			message:       "test message",
			key:           shortKey,
			expectedError: expectedError{InvalidKeySizeError: true},
			errorMsg:      aes.KeySizeError(8),
			expected:      "",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err error

			var encryptedString string

			cryptor := Crypto{
				EncryptionKey: tc.key,
			}

			if !tc.expectedError.Base64Error && !tc.expectedError.NewCipherError && !tc.expectedError.AuthenticationError && !tc.expectedError.FileReadError && !tc.expectedError.InvalidKeySizeError {
				encryptedString, err = cryptor.Encrypt(tc.message)
				assert.NoError(t, err)
				assert.NotEmpty(t, encryptedString)
				decryptedMessage, err := cryptor.Decrypt(encryptedString)
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, decryptedMessage)
			}

			if tc.expectedError.InvalidKeySizeError {
				_, err = cryptor.Encrypt(tc.message)
				assert.Equal(t, tc.errorMsg, err)
				assert.Equal(t, tc.expected, encryptedString)
			}
		})
	}
}

func TestGenerateKey(t *testing.T) {
	t.Parallel()

	cryptor := Crypto{}
	key := cryptor.GenerateKey()
	assert.Len(t, key, 44)

	decoded, err := base64.StdEncoding.DecodeString(key)
	assert.NoError(t, err)
	assert.Len(t, decoded, 32)

	assert.NotEqual(t, key, cryptor.GenerateKey())
}

func TestKeyBytes(t *testing.T) {
	t.Parallel()

	generated := Crypto{}.GenerateKey()
	decoded, err := base64.StdEncoding.DecodeString(generated)
	assert.NoError(t, err)

	tests := []struct {
		name     string
		key      string
		expected []byte
	}{
		{"generated 44-char key decodes to 32 raw bytes", generated, decoded},
		{"legacy 16-char key uses raw string bytes", "0123456789abcdef", []byte("0123456789abcdef")},
		{"legacy 24-char key uses raw string bytes", "0123456789abcdef01234567", []byte("0123456789abcdef01234567")},
		{"legacy 32-char key uses raw string bytes", validKey, []byte(validKey)},
		{"44-char non-base64 key uses raw string bytes", strings.Repeat("!", 44), []byte(strings.Repeat("!", 44))},
		{"44-char base64 decoding to 31 bytes uses raw string bytes", strings.Repeat("A", 42) + "==", []byte(strings.Repeat("A", 42) + "==")},
		{"empty key uses raw string bytes", "", []byte("")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.expected, keyBytes(tc.key))
		})
	}
}

func TestEncryptDecryptWithGeneratedKey(t *testing.T) {
	t.Parallel()

	cryptor := Crypto{EncryptionKey: Crypto{}.GenerateKey()}

	encryptedString, err := cryptor.Encrypt("test message")
	assert.NoError(t, err)

	decryptedMessage, err := cryptor.Decrypt(encryptedString)
	assert.NoError(t, err)
	assert.Equal(t, "test message", decryptedMessage)
}

func TestEncryptTreatsNonBase64KeyAsLegacy(t *testing.T) {
	t.Parallel()

	// 44 characters but not valid base64: must fall back to the legacy
	// raw-byte interpretation and be rejected as an invalid AES key size.
	cryptor := Crypto{EncryptionKey: strings.Repeat("!", 44)}

	_, err := cryptor.Encrypt("test message")
	assert.Equal(t, aes.KeySizeError(44), err)
}
