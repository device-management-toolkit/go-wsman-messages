/*********************************************************************
 * Copyright (c) Intel Corporation 2025
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package security

import (
	"crypto/aes"
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
	assert.NotEmpty(t, key)
}
