package boot

import (
	"bytes"
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

	// Missing dependent parameter
	missingDependentParams := []TLVParameter{
		{
			Type:   OCR_EFI_NETWORK_DEVICE_PATH,
			Length: 22,
			Value:  []byte("https://example.com/boot"),
		},
		{
			Type:   OCR_EFI_FILE_DEVICE_PATH,
			Length: 5,
			Value:  []byte("/boot"),
		},
		// Missing OCR_EFI_DEVICE_PATH_LEN
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
