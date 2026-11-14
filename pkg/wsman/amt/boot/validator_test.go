/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestValidateTLVEntry(t *testing.T) {
	tests := []struct {
		name        string
		paramType   ParameterType
		length      byte
		value       []byte
		expectError bool
	}{
		{
			name:        "Invalid parameter type",
			paramType:   255, // Assuming this doesn't exist in ParameterNames
			length:      1,
			value:       []byte{0},
			expectError: true,
		},
		{
			name:        "Parameter exceeds maximum size",
			paramType:   OCR_HTTPS_CERT_SYNC_ROOT_CA, // Using a valid type
			length:      byte(MaxSizes[OCR_HTTPS_CERT_SYNC_ROOT_CA] + 1),
			value:       make([]byte, MaxSizes[OCR_HTTPS_CERT_SYNC_ROOT_CA]+1),
			expectError: true,
		},
		{
			name:        "Valid boolean parameter - value 0",
			paramType:   OCR_HTTPS_CERT_SYNC_ROOT_CA,
			length:      1,
			value:       []byte{0},
			expectError: false,
		},
		{
			name:        "Valid boolean parameter - value 1",
			paramType:   OCR_HTTPS_CERT_SYNC_ROOT_CA,
			length:      1,
			value:       []byte{1},
			expectError: false,
		},
		{
			name:        "Invalid boolean parameter",
			paramType:   OCR_HTTPS_CERT_SYNC_ROOT_CA,
			length:      1,
			value:       []byte{2}, // Invalid boolean value
			expectError: true,
		},
		{
			name:        "Valid verify method - FullName",
			paramType:   OCR_HTTPS_SERVER_NAME_VERIFY_METHOD,
			length:      2,
			value:       []byte{1, 0}, // Little-endian 1
			expectError: false,
		},
		{
			name:        "Valid verify method - DomainSuffix",
			paramType:   OCR_HTTPS_SERVER_NAME_VERIFY_METHOD,
			length:      2,
			value:       []byte{2, 0}, // Little-endian 2
			expectError: false,
		},
		{
			name:        "Valid verify method - Other",
			paramType:   OCR_HTTPS_SERVER_NAME_VERIFY_METHOD,
			length:      2,
			value:       []byte{3, 0}, // Little-endian 3
			expectError: false,
		},
		{
			name:        "Invalid verify method",
			paramType:   OCR_HTTPS_SERVER_NAME_VERIFY_METHOD,
			length:      2,
			value:       []byte{4, 0}, // Little-endian 4, which is invalid
			expectError: true,
		},
		{
			name:        "Valid device path length",
			paramType:   OCR_EFI_DEVICE_PATH_LEN,
			length:      2,
			value:       []byte{10, 0}, // Any 2-byte value
			expectError: false,
		},
		{
			name:        "Invalid device path length - wrong size",
			paramType:   OCR_EFI_DEVICE_PATH_LEN,
			length:      1, // Should be 2 for UINT16
			value:       []byte{10},
			expectError: true,
		},
		{
			name:        "Valid URI",
			paramType:   OCR_EFI_NETWORK_DEVICE_PATH,
			length:      22,
			value:       []byte("https://example.com/boot"),
			expectError: false,
		},
		{
			name:        "Invalid URI",
			paramType:   OCR_EFI_NETWORK_DEVICE_PATH,
			length:      11,
			value:       []byte("not a valid uri"),
			expectError: true,
		},
		{
			name:        "Unimplemented validation",
			paramType:   OCR_HTTPS_USER_NAME,
			length:      5,
			value:       []byte("admin"),
			expectError: true, // The function returns an error for these types
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTLVEntry(tt.paramType, tt.length, tt.value)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateTLVEntry() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestParseTLVBuffer(t *testing.T) {
	// Create a valid TLV buffer with a URI
	validBuffer := []byte{
		byte(OCR_EFI_NETWORK_DEVICE_PATH), 22, // Type, Length
		'h', 't', 't', 'p', 's', ':', '/', '/', 'e', 'x', 'a', 'm', 'p', 'l', 'e', '.', 'c', 'o', 'm', '/', 'b', 'o', // Value
	}

	// Create a buffer with an invalid TLV entry (wrong boolean value)
	invalidEntryBuffer := []byte{
		byte(OCR_EFI_NETWORK_DEVICE_PATH), 22, // Type, Length
		'h', 't', 't', 'p', 's', ':', '/', '/', 'e', 'x', 'a', 'm', 'p', 'l', 'e', '.', 'c', 'o', 'm', '/', 'b', 'o', // Value
		byte(OCR_HTTPS_CERT_SYNC_ROOT_CA), 1, // Type, Length
		2, // Invalid boolean value
	}

	// Create a buffer with missing mandatory parameter
	missingMandatoryBuffer := []byte{
		byte(OCR_HTTPS_CERT_SYNC_ROOT_CA), 1, // Type, Length
		1, // Value
	}

	// Create a buffer with dependent parameter missing
	missingDependentBuffer := []byte{
		byte(OCR_EFI_NETWORK_DEVICE_PATH), 22, // Type, Length
		'h', 't', 't', 'p', 's', ':', '/', '/', 'e', 'x', 'a', 'm', 'p', 'l', 'e', '.', 'c', 'o', 'm', '/', 'b', 'o', // Value
		byte(OCR_EFI_FILE_DEVICE_PATH), 5, // Type, Length
		'/', 'b', 'o', 'o', 't', // Value, but missing OCR_EFI_DEVICE_PATH_LEN
	}

	// Create a buffer with incomplete TLV (missing length)
	incompleteBuffer := []byte{
		byte(OCR_EFI_NETWORK_DEVICE_PATH), // Type only, missing length
	}

	// Create a buffer with insufficient value bytes
	insufficientValueBuffer := []byte{
		byte(OCR_EFI_NETWORK_DEVICE_PATH), 22, // Type, Length
		'h', 't', 't', 'p', 's', // Only 5 bytes instead of 22
	}

	tests := []struct {
		name        string
		buffer      []byte
		expectValid bool
		paramCount  int
	}{
		{
			name:        "Valid buffer",
			buffer:      validBuffer,
			expectValid: true,
			paramCount:  1,
		},
		{
			name:        "Invalid entry in buffer",
			buffer:      invalidEntryBuffer,
			expectValid: false,
			paramCount:  2, // Still parsed both entries
		},
		{
			name:        "Missing mandatory parameter",
			buffer:      missingMandatoryBuffer,
			expectValid: false,
			paramCount:  1,
		},
		{
			name:        "Missing dependent parameter",
			buffer:      missingDependentBuffer,
			expectValid: false,
			paramCount:  2,
		},
		{
			name:        "Incomplete TLV entry",
			buffer:      incompleteBuffer,
			expectValid: false,
			paramCount:  0,
		},
		{
			name:        "Insufficient value bytes",
			buffer:      insufficientValueBuffer,
			expectValid: false,
			paramCount:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseTLVBuffer(tt.buffer)
			if result.Valid != tt.expectValid {
				t.Errorf("ParseTLVBuffer() valid = %v, expectValid %v, errors: %v",
					result.Valid, tt.expectValid, result.Errors)
			}

			if len(result.Parameters) != tt.paramCount {
				t.Errorf("ParseTLVBuffer() param count = %v, expected %v",
					len(result.Parameters), tt.paramCount)
			}
		})
	}
}

func TestCreateTLVBuffer(t *testing.T) {
	// Create test parameters
	params := []TLVParameter{
		{
			Type:   OCR_EFI_NETWORK_DEVICE_PATH,
			Length: 22,
			Value:  []byte("https://example.com/boot"),
		},
		{
			Type:   OCR_HTTPS_CERT_SYNC_ROOT_CA,
			Length: 1,
			Value:  []byte{1},
		},
	}

	// Create buffer
	buffer, err := CreateTLVBuffer(params)
	if err != nil {
		t.Fatalf("CreateTLVBuffer() error = %v", err)
	}

	// Check buffer content
	// For each parameter, we expect:
	// - 2 bytes for vendor ID (0x8086)
	// - 2 bytes for parameter type
	// - 4 bytes for length
	// - n bytes for value

	// Due to the custom format in CreateTLVBuffer, we need to verify the structure accordingly
	if len(buffer) == 0 {
		t.Errorf("CreateTLVBuffer() returned empty buffer")
	}
}

func TestBuildRPEParameters(t *testing.T) {
	tests := []struct {
		name        string
		bitmask     uint32
		credentials map[ParameterType]string
		expected    []RPEParameter
		expectError bool
	}{
		{
			name:    "bitmask and nonempty credentials",
			bitmask: 0x12345678,
			credentials: map[ParameterType]string{
				RPE_PSID:                "psid",
				RPE_SSD_MASTER_PASSWORD: "password",
				RPE_OEM_PARAMETER:       "oem",
			},
			expected: []RPEParameter{
				{Type: RPE_DEVICE_BITMASK, Value: []byte{0x78, 0x56, 0x34, 0x12}},
				{Type: RPE_PSID, Value: []byte("psid")},
				{Type: RPE_SSD_MASTER_PASSWORD, Value: []byte("password")},
				{Type: RPE_OEM_PARAMETER, Value: []byte("oem")},
			},
		},
		{
			name:    "empty credentials are omitted",
			bitmask: 1,
			credentials: map[ParameterType]string{
				RPE_PSID: "",
			},
			expected: []RPEParameter{
				{Type: RPE_DEVICE_BITMASK, Value: []byte{1, 0, 0, 0}},
			},
		},
		{
			name: "credential exceeds maximum size",
			credentials: map[ParameterType]string{
				RPE_SSD_MASTER_PASSWORD: string(bytes.Repeat([]byte{'x'}, RPEMaxSizes[RPE_SSD_MASTER_PASSWORD]+1)),
			},
			expectError: true,
		},
		{
			name: "unsupported credential type",
			credentials: map[ParameterType]string{
				OCR_HTTPS_PASSWORD: "password",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameters, err := BuildRPEParameters(tt.bitmask, tt.credentials)
			if (err != nil) != tt.expectError {
				t.Fatalf("BuildRPEParameters() error = %v, want error %v", err, tt.expectError)
			}

			if tt.expectError {
				return
			}

			if len(parameters) != len(tt.expected) {
				t.Fatalf("BuildRPEParameters() returned %d parameters, want %d", len(parameters), len(tt.expected))
			}

			for i, parameter := range parameters {
				expected := tt.expected[i]
				if parameter.Type != expected.Type || !bytes.Equal(parameter.Value, expected.Value) {
					t.Errorf("BuildRPEParameters()[%d] = %+v, want %+v", i, parameter, expected)
				}
			}
		})
	}
}

func TestCreateRPETLVBuffer(t *testing.T) {
	params := []RPEParameter{
		{Type: RPE_DEVICE_BITMASK, Value: []byte{0x78, 0x56, 0x34, 0x12}},
		{Type: RPE_SSD_MASTER_PASSWORD, Value: []byte("secret")},
		{Type: RPE_OEM_PARAMETER, Value: bytes.Repeat([]byte{'x'}, RPEMaxSizes[RPE_OEM_PARAMETER])},
	}

	buffer, err := CreateRPETLVBuffer(params)
	if err != nil {
		t.Fatalf("CreateRPETLVBuffer() error = %v", err)
	}

	expectedPrefix := []byte{
		0x86, 0x80, 0x01, 0x00, 0x04, 0x00, 0x00, 0x00,
		0x78, 0x56, 0x34, 0x12,
		0x86, 0x80, 0x14, 0x00, 0x06, 0x00, 0x00, 0x00,
		's', 'e', 'c', 'r', 'e', 't',
	}
	if !bytes.Equal(buffer[:len(expectedPrefix)], expectedPrefix) {
		t.Errorf("CreateRPETLVBuffer() bitmask and SSD password entries = %v, want %v",
			buffer[:len(expectedPrefix)], expectedPrefix)
	}

	oemLengthOffset := len(expectedPrefix) + 4
	oemLength := binary.LittleEndian.Uint32(buffer[oemLengthOffset : oemLengthOffset+4])

	if oemLength != uint32(RPEMaxSizes[RPE_OEM_PARAMETER]) {
		t.Errorf("CreateRPETLVBuffer() OEM parameter length = %d, want %d",
			oemLength, RPEMaxSizes[RPE_OEM_PARAMETER])
	}
}

func TestCreateRPETLVBufferInvalidBitmaskLength(t *testing.T) {
	tests := []struct {
		name  string
		value []byte
	}{
		{name: "nil"},
		{name: "empty", value: []byte{}},
		{name: "one byte", value: []byte{1}},
		{name: "two bytes", value: []byte{1, 2}},
		{name: "three bytes", value: []byte{1, 2, 3}},
		{name: "five bytes", value: []byte{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer, err := CreateRPETLVBuffer([]RPEParameter{
				{Type: RPE_PSID, Value: []byte("psid")},
				{Type: RPE_DEVICE_BITMASK, Value: tt.value},
			})
			if err == nil {
				t.Fatal("CreateRPETLVBuffer() expected an error for invalid bitmask length")
			}

			if buffer != nil {
				t.Errorf("CreateRPETLVBuffer() buffer = %v, want nil", buffer)
			}
		})
	}
}

func TestGetUint16Value(t *testing.T) {
	tests := []struct {
		name        string
		param       TLVParameter
		expectedVal uint16
		expectError bool
	}{
		{
			name: "Valid uint16",
			param: TLVParameter{
				Type:   OCR_EFI_DEVICE_PATH_LEN,
				Length: 2,
				Value:  []byte{0x34, 0x12}, // Little-endian 0x1234
			},
			expectedVal: 0x1234,
			expectError: false,
		},
		{
			name: "Invalid length",
			param: TLVParameter{
				Type:   OCR_EFI_DEVICE_PATH_LEN,
				Length: 1,
				Value:  []byte{0x34}, // Only 1 byte
			},
			expectedVal: 0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := GetUint16Value(tt.param)
			if (err != nil) != tt.expectError {
				t.Errorf("GetUint16Value() error = %v, expectError %v", err, tt.expectError)
			}

			if err == nil && val != tt.expectedVal {
				t.Errorf("GetUint16Value() = %v, want %v", val, tt.expectedVal)
			}
		})
	}
}

func TestGetStringValue(t *testing.T) {
	tests := []struct {
		name     string
		param    TLVParameter
		expected string
	}{
		{
			name: "Simple string",
			param: TLVParameter{
				Type:   OCR_HTTPS_SERVER_CERT_HASH_SHA256,
				Length: 5,
				Value:  []byte("hello"),
			},
			expected: "hello",
		},
		{
			name: "String with null terminator",
			param: TLVParameter{
				Type:   OCR_HTTPS_SERVER_CERT_HASH_SHA256,
				Length: 6,
				Value:  []byte{'h', 'e', 'l', 'l', 'o', 0},
			},
			expected: "hello",
		},
		{
			name: "Empty string",
			param: TLVParameter{
				Type:   OCR_HTTPS_SERVER_CERT_HASH_SHA256,
				Length: 0,
				Value:  []byte{},
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetStringValue(tt.param)
			if result != tt.expected {
				t.Errorf("GetStringValue() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNewUint16Parameter(t *testing.T) {
	tests := []struct {
		name      string
		paramType ParameterType
		value     uint16
	}{
		{
			name:      "Create uint16 parameter",
			paramType: OCR_EFI_DEVICE_PATH_LEN,
			value:     0x1234,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			param := NewUint16Parameter(tt.paramType, tt.value)
			if param.Type != tt.paramType {
				t.Errorf("NewUint16Parameter() type = %v, want %v", param.Type, tt.paramType)
			}

			if param.Length != 2 {
				t.Errorf("NewUint16Parameter() length = %v, want 2", param.Length)
			}

			// Check if the value is correctly stored in little-endian
			expected := make([]byte, 2)
			expected[0] = byte(tt.value)      // Low byte
			expected[1] = byte(tt.value >> 8) // High byte

			if !bytes.Equal(param.Value, expected) {
				t.Errorf("NewUint16Parameter() value = %v, want %v", param.Value, expected)
			}

			// Double-check by parsing it back
			parsedValue, err := GetUint16Value(param)
			if err != nil {
				t.Errorf("Failed to parse back uint16: %v", err)
			}

			if parsedValue != tt.value {
				t.Errorf("Round-trip failed: got %v, want %v", parsedValue, tt.value)
			}
		})
	}
}

func TestNewBoolParameter(t *testing.T) {
	tests := []struct {
		name      string
		paramType ParameterType
		value     bool
		expected  byte
	}{
		{
			name:      "Boolean true",
			paramType: OCR_HTTPS_CERT_SYNC_ROOT_CA,
			value:     true,
			expected:  1,
		},
		{
			name:      "Boolean false",
			paramType: OCR_HTTPS_CERT_SYNC_ROOT_CA,
			value:     false,
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			param := NewBoolParameter(tt.paramType, tt.value)
			if param.Type != tt.paramType {
				t.Errorf("NewBoolParameter() type = %v, want %v", param.Type, tt.paramType)
			}

			if param.Length != 1 {
				t.Errorf("NewBoolParameter() length = %v, want 1", param.Length)
			}

			if len(param.Value) != 1 || param.Value[0] != tt.expected {
				t.Errorf("NewBoolParameter() value = %v, want [%v]", param.Value, tt.expected)
			}
		})
	}
}

func TestNewUint32Parameter(t *testing.T) {
	tests := []struct {
		name      string
		paramType ParameterType
		value     uint32
	}{
		{
			name:      "Zero value",
			paramType: RPE_DEVICE_BITMASK,
			value:     0,
		},
		{
			name:      "Device bitmask value",
			paramType: RPE_DEVICE_BITMASK,
			value:     0x00000003,
		},
		{
			name:      "Max uint32 value",
			paramType: RPE_DEVICE_BITMASK,
			value:     0xFFFFFFFF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			param := NewUint32Parameter(tt.paramType, tt.value)

			if param.Type != tt.paramType {
				t.Errorf("NewUint32Parameter() type = %v, want %v", param.Type, tt.paramType)
			}

			if param.Length != 4 {
				t.Errorf("NewUint32Parameter() length = %v, want 4", param.Length)
			}

			if len(param.Value) != 4 {
				t.Fatalf("NewUint32Parameter() value length = %v, want 4", len(param.Value))
			}

			// Verify little-endian encoding
			expected := []byte{
				byte(tt.value),
				byte(tt.value >> 8),
				byte(tt.value >> 16),
				byte(tt.value >> 24),
			}

			if !bytes.Equal(param.Value, expected) {
				t.Errorf("NewUint32Parameter() value = %v, want %v", param.Value, expected)
			}
		})
	}
}

func TestValidateParameters(t *testing.T) {
	// Valid parameters with network device path
	validParams := []TLVParameter{
		{
			Type:   OCR_EFI_NETWORK_DEVICE_PATH,
			Length: 22,
			Value:  []byte("https://example.com/boot"),
		},
		{
			Type:   OCR_HTTPS_CERT_SYNC_ROOT_CA,
			Length: 1,
			Value:  []byte{1},
		},
	}

	// Invalid parameter type
	invalidTypeParams := []TLVParameter{
		{
			Type:   OCR_EFI_NETWORK_DEVICE_PATH,
			Length: 22,
			Value:  []byte("https://example.com/boot"),
		},
		{
			Type:   255, // Invalid type
			Length: 1,
			Value:  []byte{1},
		},
	}

	// Missing mandatory parameter
	missingMandatoryParams := []TLVParameter{
		{
			Type:   OCR_HTTPS_CERT_SYNC_ROOT_CA,
			Length: 1,
			Value:  []byte{1},
		},
	}

	// Valid parameters with device file path
	deviceFileParams := []TLVParameter{
		{
			Type:   OCR_EFI_DEVICE_PATH_LEN,
			Length: 2,
			Value:  []byte{5, 0}, // Length of device path
		},
		{
			Type:   OCR_EFI_FILE_DEVICE_PATH,
			Length: 5,
			Value:  []byte("/boot"),
		},
	}

	// Missing OCR_EFI_DEVICE_PATH_LEN
	missingDependentParams := []TLVParameter{
		{
			Type:   OCR_EFI_FILE_DEVICE_PATH,
			Length: 5,
			Value:  []byte("/boot"),
		},
	}

	// Invalid value for a parameter
	invalidValueParams := []TLVParameter{
		{
			Type:   OCR_EFI_NETWORK_DEVICE_PATH,
			Length: 22,
			Value:  []byte("https://example.com/boot"),
		},
		{
			Type:   OCR_HTTPS_CERT_SYNC_ROOT_CA,
			Length: 1,
			Value:  []byte{2}, // Invalid boolean, should be 0 or 1
		},
	}

	tests := []struct {
		name        string
		params      []TLVParameter
		expectValid bool
		errorCount  int
	}{
		{
			name:        "Valid parameters",
			params:      validParams,
			expectValid: true,
			errorCount:  0,
		},
		{
			name:        "Invalid parameter type",
			params:      invalidTypeParams,
			expectValid: false,
			errorCount:  1,
		},
		{
			name:        "Missing mandatory parameter",
			params:      missingMandatoryParams,
			expectValid: false,
			errorCount:  1,
		},
		{
			name:        "Valid device file parameters",
			params:      deviceFileParams,
			expectValid: true,
			errorCount:  0,
		},
		{
			name:        "Missing dependent parameter",
			params:      missingDependentParams,
			expectValid: false,
			errorCount:  1,
		},
		{
			name:        "Invalid value in parameter",
			params:      invalidValueParams,
			expectValid: false,
			errorCount:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, errors := ValidateParameters(tt.params)
			if valid != tt.expectValid {
				t.Errorf("ValidateParameters() valid = %v, want %v", valid, tt.expectValid)
			}

			if len(errors) != tt.errorCount {
				t.Errorf("ValidateParameters() error count = %v, want %v, errors: %v",
					len(errors), tt.errorCount, errors)
			}
		})
	}
}
