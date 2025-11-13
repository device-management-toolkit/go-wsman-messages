package apf

import (
	"encoding/binary"
)

// ValidateProtocolVersion checks if the data length is at least 93 bytes for APF_PROTOCOLVERSION.
func ValidateProtocolVersion(data []byte) bool {
	return len(data) >= 93
}

// ValidateServiceRequest checks if the data length is sufficient for APF_SERVICE_REQUEST.
func ValidateServiceRequest(data []byte) bool {
	if len(data) < 5 {
		return false
	}

	serviceLen := int(binary.BigEndian.Uint32(data[1:5]))

	return len(data) >= 5+serviceLen
}

// ValidateGlobalRequest checks if the data length is sufficient for APF_GLOBAL_REQUEST.
func ValidateGlobalRequest(data []byte) bool {
	if len(data) < 5 {
		return false
	}

	globalReqLen := int(binary.BigEndian.Uint32(data[1:5]))
	if len(data) < 5+globalReqLen+1 {
		return false
	}

	serviceName := string(data[5 : 5+globalReqLen])

	if serviceName == APF_GLOBAL_REQUEST_STR_TCP_FORWARD_REQUEST || serviceName == APF_GLOBAL_REQUEST_STR_TCP_FORWARD_CANCEL_REQUEST {
		if len(data) < 6+globalReqLen+4 {
			return false
		}

		addrLen := int(binary.BigEndian.Uint32(data[6+globalReqLen : 10+globalReqLen]))

		return len(data) >= 14+globalReqLen+addrLen
	}

	return false
}

// ValidateChannelOpenConfirmation checks if the data length is at least 17 bytes for APF_CHANNEL_OPEN_CONFIRMATION.
func ValidateChannelOpenConfirmation(data []byte) bool {
	return len(data) >= 17
}

// ValidateChannelOpenFailure checks if the data length is at least 17 bytes for APF_CHANNEL_OPEN_FAILURE.
func ValidateChannelOpenFailure(data []byte) bool {
	return len(data) >= 17
}

// ValidateChannelClose checks if the data length is at least 5 bytes for APF_CHANNEL_CLOSE.
func ValidateChannelClose(data []byte) bool {
	return len(data) >= 5
}

// ValidateChannelData checks if the data length is sufficient for APF_CHANNEL_DATA.
func ValidateChannelData(data []byte) bool {
	result := len(data) >= 9 && len(data) >= 9+int(binary.BigEndian.Uint32(data[5:9]))

	return result
}

// ValidateChannelWindowAdjust checks if the data length is at least 9 bytes for APF_CHANNEL_WINDOW_ADJUST.
func ValidateChannelWindowAdjust(data []byte) bool {
	return len(data) >= 9
}
