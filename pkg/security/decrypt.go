/*********************************************************************
 * Copyright (c) Intel Corporation 2025
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package security

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/config"
)

// Decrypt cipher text using AES-GCM with the provided key.
func (c Crypto) Decrypt(cipherText string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(c.EncryptionKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(data) < gcm.NonceSize() {
		return "", errors.New("cipher text too short")
	}

	nonce, cText := data[:gcm.NonceSize()], data[gcm.NonceSize():]

	plainText, err := gcm.Open(nil, nonce, cText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

// Read encrypted data from file and decrypt it.
func (c Crypto) ReadAndDecryptFile(filePath string) (config.Configuration, error) {
	encryptedData, err := os.ReadFile(filePath)
	if err != nil {
		return config.Configuration{}, err
	}

	decryptedData, err := c.Decrypt(string(encryptedData))
	if err != nil {
		return config.Configuration{}, err
	}

	var configuration config.Configuration

	err = yaml.Unmarshal([]byte(decryptedData), &configuration)
	if err != nil {
		return config.Configuration{}, err
	}

	return configuration, nil
}
