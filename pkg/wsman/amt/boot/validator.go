package boot

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/url"
)

// ValidateTLVEntry validates a single TLV entry.
func ValidateTLVEntry(paramType ParameterType, length byte, value []byte) error {
	// Check if type is valid
	if _, exists := ParameterNames[paramType]; !exists {
		return fmt.Errorf("invalid parameter type: %d", paramType)
	}

	// Check length against max size
	maxSize := MaxSizes[paramType]
	if int(length) > maxSize {
		return fmt.Errorf("parameter %s exceeds maximum size: %d > %d",
			ParameterNames[paramType], length, maxSize)
	}

	// Type-specific validations
	switch paramType {
	case OCR_HTTPS_CERT_SYNC_ROOT_CA:
		// Validate boolean value (should be 0 or 1)
		if len(value) > 0 && value[0] != 0 && value[0] != 1 {
			return fmt.Errorf("invalid boolean value for %s: %d",
				ParameterNames[paramType], value[0])
		}

	case OCR_HTTPS_SERVER_NAME_VERIFY_METHOD:
		// Validate verify method (1=FullName, 2=DomainSuffix, 3=Other)
		if len(value) >= 2 {
			method := uint16(value[0]) | uint16(value[1])<<8 // Little-endian UINT16
			if method < 1 || method > 3 {
				return fmt.Errorf("invalid verification method: %d", method)
			}
		}

	case OCR_EFI_DEVICE_PATH_LEN:
		// Must be valid UINT16
		if int(length) != 2 {
			return fmt.Errorf("invalid length for UINT16 device path length: %d", length)
		}

	case OCR_EFI_NETWORK_DEVICE_PATH:
		// Validate URI
		uri := string(value)
		if _, err := url.ParseRequestURI(uri); err != nil {
			return fmt.Errorf("invalid URI for %s: %s", ParameterNames[paramType], uri)
		}

	case OCR_EFI_FILE_DEVICE_PATH:
		// Add validation if needed

	case OCR_BOOT_IMAGE_HASH_SHA256, OCR_BOOT_IMAGE_HASH_SHA384, OCR_BOOT_IMAGE_HASH_SHA512:
		// Validate hash length if needed

	case OCR_EFI_BOOT_OPTIONAL_DATA:
		// Add validation if needed

	case OCR_HTTPS_CERT_SERVER_NAME:
		// Add validation if needed

	case OCR_HTTPS_SERVER_CERT_HASH_SHA256, OCR_HTTPS_SERVER_CERT_HASH_SHA384, OCR_HTTPS_SERVER_CERT_HASH_SHA512:
		// Validate hash length if needed

	case OCR_HTTPS_REQUEST_TIMEOUT:
		// Add timeout validation if needed

	case OCR_HTTPS_USER_NAME, OCR_HTTPS_PASSWORD:
		return fmt.Errorf("validation is not added for %s", ParameterNames[paramType])
	}

	return nil
}

// ParseTLVBuffer parses and validates a TLV buffer.
func ParseTLVBuffer(buffer []byte) ValidationResult {
	result := ValidationResult{
		Valid:      true,
		Parameters: []TLVParameter{},
		Errors:     []string{},
	}

	offset := 0
	presentTypes := make(map[ParameterType]bool)

	for offset < len(buffer) {
		// Need at least 2 bytes for type and length.
		if offset+2 > len(buffer) {
			result.Valid = false
			result.Errors = append(result.Errors, "incomplete TLV entry at end of buffer")

			break
		}

		paramType := ParameterType(buffer[offset])
		length := buffer[offset+1]
		offset += 2

		// Check if we have enough bytes for the value
		if offset+int(length) > len(buffer) {
			result.Valid = false
			result.Errors = append(result.Errors,
				fmt.Sprintf("incomplete value for type %d, expected %d bytes", paramType, length))

			break
		}

		value := buffer[offset : offset+int(length)]
		offset += int(length)

		// Validate this TLV entry
		valid := true

		err := ValidateTLVEntry(paramType, length, value)
		if err != nil {
			result.Valid = false
			result.Errors = append(result.Errors, err.Error())
			valid = false
		}

		// Track which parameter types we've seen
		presentTypes[paramType] = true

		// Add to parsed parameters
		if _, exists := ParameterNames[paramType]; !exists {
			ParameterNames[paramType] = "Unknown parameter"
		}

		result.Parameters = append(result.Parameters, TLVParameter{
			Type:   paramType,
			Length: length,
			Value:  value,
			Valid:  valid,
		})
	}

	// Check for mandatory parameters
	if !presentTypes[OCR_EFI_NETWORK_DEVICE_PATH] {
		result.Valid = false
		result.Errors = append(result.Errors, "missing mandatory parameter: URI to HTTPS Server")
	}

	// Check for dependent parameters
	if presentTypes[OCR_EFI_FILE_DEVICE_PATH] && !presentTypes[OCR_EFI_DEVICE_PATH_LEN] {
		result.Valid = false
		result.Errors = append(result.Errors,
			"missing device path length which is mandatory when file device path is provided")
	}

	return result
}

// CreateTLVBuffer creates a TLV buffer from parameters.
func CreateTLVBuffer(parameters []TLVParameter) ([]byte, error) {
	buffer := []byte{}

	for _, param := range parameters {
		// default vendor Intel
		buffer = binary.LittleEndian.AppendUint16(buffer, 0x8086)

		// param type 1: OCR_EFI_NETWORK_DEVICE_PATH
		buffer = binary.LittleEndian.AppendUint16(buffer, uint16(param.Type))

		if param.Type == 20 {
			buffer = binary.LittleEndian.AppendUint32(buffer, 1)
			buffer = append(buffer, 1)
		} else {
			buffer = binary.LittleEndian.AppendUint32(buffer, uint32(len(param.Value)))
			buffer = append(buffer, param.Value...)
		}
	}

	return buffer, nil
}

// GetUint16Value retrieves a uint16 from a parameter value (little-endian).
func GetUint16Value(param TLVParameter) (uint16, error) {
	if len(param.Value) != 2 {
		return 0, fmt.Errorf("expected 2 bytes for uint16, got %d", len(param.Value))
	}

	return binary.LittleEndian.Uint16(param.Value), nil
}

// GetStringValue retrieves a string from a parameter value (UTF-8).
func GetStringValue(param TLVParameter) string {
	return string(bytes.TrimRight(param.Value, "\x00"))
}

// NewUint16Parameter creates a new parameter with a uint16 value.
func NewUint16Parameter(paramType ParameterType, value uint16) TLVParameter {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, value)

	return TLVParameter{
		Type:   paramType,
		Length: 2,
		Value:  buf,
	}
}

// NewStringParameter creates a new parameter with a string value.
func NewStringParameter(paramType ParameterType, value string) (TLVParameter, error) {
	bytes := []byte(value)
	if len(bytes) > MaxSizes[paramType] {
		return TLVParameter{}, fmt.Errorf("string value too long: %d bytes (max %d)",
			len(bytes), MaxSizes[paramType])
	}

	return TLVParameter{
		Type:   paramType,
		Length: byte(len(bytes)),
		Value:  bytes,
	}, nil
}

// NewBoolParameter creates a new parameter with a boolean value.
func NewBoolParameter(paramType ParameterType, value bool) TLVParameter {
	var boolValue byte = 0
	if value {
		boolValue = 1
	}

	return TLVParameter{
		Type:   paramType,
		Length: 1,
		Value:  []byte{boolValue},
	}
}

// ValidateParameters validates a set of parameters without creating a buffer.
func ValidateParameters(parameters []TLVParameter) (bool, []string) {
	valid := true
	errors := []string{}

	// Check each parameter individually
	presentTypes := make(map[ParameterType]bool)

	for _, param := range parameters {
		err := ValidateTLVEntry(param.Type, byte(len(param.Value)), param.Value)
		if err != nil {
			valid = false

			errors = append(errors, err.Error())
		}

		presentTypes[param.Type] = true
	}

	// Check for mandatory parameters
	if !presentTypes[OCR_EFI_NETWORK_DEVICE_PATH] {
		if !presentTypes[OCR_EFI_FILE_DEVICE_PATH] || !presentTypes[OCR_EFI_DEVICE_PATH_LEN] {
			valid = false

			errors = append(errors,
				"missing file device path or device path length which is mandatory")
		}
	}

	return valid, errors
}
