/*********************************************************************
 * Copyright (c) Intel Corporation 2025
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package common

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"
)

// ReadShort reads a short value from a string at position p.
func ReadShort(v string, p int) int {
	return int(v[p])<<8 + int(v[p+1])
}

// ReadShortX reads a short value from a string at position p in reverse order.
func ReadShortX(v string, p int) int {
	return int(v[p+1])<<8 + int(v[p])
}

// ReadInt reads an int value from a string at position p.
func ReadInt(v string, p int) int {
	return int(v[p])*0x1000000 + int(v[p+1])<<16 + int(v[p+2])<<8 + int(v[p+3])
}

// ReadIntX reads an int value from a string at position p in reverse order.
func ReadIntX(v string, p int) int {
	return int(v[p+3])*0x1000000 + int(v[p+2])<<16 + int(v[p+1])<<8 + int(v[p])
}

// ShortToStr converts a short value to a string.
func ShortToStr(v int) string {
	return string([]byte{byte((v >> 8) & 0xFF), byte(v & 0xFF)})
}

// ShortToStrX converts a short value to a string in reverse order.
func ShortToStrX(v int) string {
	return string([]byte{byte(v & 0xFF), byte((v >> 8) & 0xFF)})
}

// IntToStr converts an int value to a string.
func IntToStr(v int) string {
	return string([]byte{byte((v >> 24) & 0xFF), byte((v >> 16) & 0xFF), byte((v >> 8) & 0xFF), byte(v & 0xFF)})
}

// IntToStrX converts an int value to a string in reverse order.
func IntToStrX(v int) string {
	return string([]byte{byte(v & 0xFF), byte((v >> 8) & 0xFF), byte((v >> 16) & 0xFF), byte((v >> 24) & 0xFF)})
}

// MakeToArray attempts to convert an interface to a slice of interfaces.
func MakeToArray(v interface{}) []interface{} {
	return []interface{}{v}
}

// Rstr2hex converts a raw string to a hex string.
func Rstr2hex(input string) string {
	result := ""
	for i := 0; i < len(input); i++ {
		result += fmt.Sprintf("%02X", input[i])
	}

	return result
}

// Hex2rstr converts a hex string to a raw string.
func Hex2rstr(d string) string {
	bytes, err := hex.DecodeString(d)
	if err != nil {
		panic(err) // For simplicity, though you might want to handle errors more gracefully
	}

	return string(bytes)
}

// Char2hex converts a character code to hex.
func Char2hex(i int) string {
	return fmt.Sprintf("%02X", i)
}

// ComputeDigesthash computes the MD5 digest hash for a set of values.
func ComputeDigesthash(username, password, realm, method, path, qop, nonce, nc, cnonce string) string {
	ha1 := md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", username, realm, password)))
	ha2 := md5.Sum([]byte(fmt.Sprintf("%s:%s", method, path)))
	final := md5.Sum([]byte(fmt.Sprintf("%x:%s:%s:%s:%s:%x", ha1, nonce, nc, cnonce, qop, ha2)))

	return fmt.Sprintf("%x", final)
}

// RandomValueHex generates a random hex string of a given length.
func RandomValueHex(i int) string {
	b := make([]byte, (i+1)/2) // +1 to handle odd lengths
	if _, err := rand.Read(b); err != nil {
		panic(err) // For simplicity, though you might want to handle errors more gracefully
	}

	return hex.EncodeToString(b)[:i]
}

// GetSidString converts a byte array of SID into string.
// Note: This function assumes sid is provided as a string for simplicity but should be []byte in a real Go implementation.
func GetSidString(sid string) string {
	value := fmt.Sprintf("S-%d-%d", sid[0], sid[7])

	for i := 2; i < len(sid)/4; i++ {
		substr := sid[i*4 : (i+1)*4]

		intValue, err := strconv.ParseInt(strconv.Itoa(ReadIntX(substr, 0)), 10, 64)
		if err != nil {
			return ""
		}

		value += fmt.Sprintf("-%d", intValue)
	}

	return value
}
