/*********************************************************************
 * Copyright (c) Intel Corporation 2022
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package apf

// AuthRequest contains the decoded fields from an APF_USERAUTH_REQUEST message.
type AuthRequest struct {
	Username    string
	ServiceName string
	MethodName  string
	Password    string // Only populated when MethodName == "password"
}

// AuthResponse indicates whether authentication succeeded.
type AuthResponse struct {
	Authenticated bool
}

// GlobalRequest contains the decoded fields from an APF_GLOBAL_REQUEST message.
type GlobalRequest struct {
	RequestType string // "tcpip-forward" or "cancel-tcpip-forward"
	Address     string
	Port        uint32
}

// ProtocolVersionInfo contains the decoded fields from an APF_PROTOCOL_VERSION_MESSAGE.
type ProtocolVersionInfo struct {
	MajorVersion  uint32
	MinorVersion  uint32
	TriggerReason uint32
	UUID          string   // Formatted as GUID string (e.g., "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX")
	RawUUID       [16]byte // Raw UUID bytes
}

// Handler is the interface that applications must implement to receive
// APF protocol events and provide application-specific decisions.
//
// Methods are called synchronously during message processing.
// The library handles protocol encoding/decoding; the application
// handles business logic (authentication, device registration, etc.).
type Handler interface {
	// OnProtocolVersion is called when an APF_PROTOCOLVERSION message is received.
	// The application can use the UUID to identify the device.
	// Returns an error if the device should be rejected.
	OnProtocolVersion(info ProtocolVersionInfo) error

	// OnAuthRequest is called when an APF_USERAUTH_REQUEST message is received.
	// The application should validate the credentials and return an AuthResponse.
	// The library will generate the appropriate success/failure response message.
	OnAuthRequest(request AuthRequest) AuthResponse

	// OnGlobalRequest is called when an APF_GLOBAL_REQUEST message is received.
	// The application can use this to track TCP forwarding requests.
	// Returns true if the application wants to send a keep-alive options request
	// after the response (caller is responsible for sending it separately).
	OnGlobalRequest(request GlobalRequest) bool
}

// DefaultHandler provides a no-op implementation of Handler.
// This maintains backward compatibility - existing code that doesn't need
// callbacks can continue to work unchanged.
type DefaultHandler struct{}

// OnProtocolVersion accepts all protocol version messages.
func (h DefaultHandler) OnProtocolVersion(info ProtocolVersionInfo) error {
	return nil
}

// OnAuthRequest does not authenticate by default.
func (h DefaultHandler) OnAuthRequest(request AuthRequest) AuthResponse {
	return AuthResponse{Authenticated: false}
}

// OnGlobalRequest returns false (no keep-alive by default).
func (h DefaultHandler) OnGlobalRequest(request GlobalRequest) bool {
	return false
}
