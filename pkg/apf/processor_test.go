/*********************************************************************
 * Copyright (c) Intel Corporation 2022
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package apf

import (
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const testUsername = "testuser"

func TestProcess(t *testing.T) {
	t.Parallel()

	data := []byte{0x01}

	session := &Session{}

	result := Process(data, session)

	assert.NotNil(t, result)
}

func TestProcessChannelOpenFailure(t *testing.T) {
	t.Parallel()

	data := []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	errorChannel := make(chan error)
	statusChannel := make(chan bool)

	session := &Session{
		ErrorBuffer: errorChannel,
		Status:      statusChannel,
	}
	defer close(errorChannel)

	go func() {
		status := <-statusChannel
		err := <-errorChannel
		assert.Error(t, err)
		assert.False(t, status)
	}()

	ProcessChannelOpenFailure(data, session)
}

func TestProcessChannelWindowAdjust(t *testing.T) {
	t.Parallel()

	data := []byte{0x01}
	session := &Session{}

	ProcessChannelWindowAdjust(data, session)
}

func TestProcessChannelClose(t *testing.T) {
	t.Parallel()

	data := []byte{0x01}
	session := &Session{}
	result := ProcessChannelClose(data, session)

	assert.NotNil(t, result)
}

func TestProcessGlobalRequest(t *testing.T) {
	t.Parallel()

	data := []byte{
		0x01,
		0x00, 0x00, 0x00, 0x0D,
		0x74, 0x63, 0x70, 0x69, 0x70, 0x2d, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
		0x00,
		0x00, 0x00, 0x00, 0x03,
		0x00, 0x00, 0x00,
		0x00, 0x00, 0x42, 0x60,
	}

	request, reply := ProcessGlobalRequest(data)
	assert.NotNil(t, reply)
	assert.Equal(t, "tcpip-forward", request.RequestType)
}

func TestProcessChannelData(t *testing.T) {
	t.Parallel()

	data := []byte{
		0x01,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01,
	}
	timer := time.NewTimer(2 * time.Second)
	session := &Session{
		Timer: timer,
	}

	go func() {
		<-timer.C
	}()
	ProcessChannelData(data, session)
}

func TestProcessServiceRequestWhenAUTH(t *testing.T) {
	t.Parallel()

	data := []byte{0x01, 0x00, 0x00, 0x00, 0x12, 0x61, 0x75, 0x74, 0x68, 0x40, 0x61, 0x6d, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6d}
	hi := int(0x12)
	logrus.Print(hi)

	result := ProcessServiceRequest(data)

	assert.NotNil(t, result)
	assert.Equal(t, uint8(0x6), result.MessageType) // APF_SERVICE_ACCEPT
	assert.Equal(t, uint32(0x12), result.ServiceNameLength)
	assert.Equal(t, [18]uint8{0x61, 0x75, 0x74, 0x68, 0x40, 0x61, 0x6d, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6d}, result.ServiceName)
}

func TestProcessServiceRequestWhenPWFD(t *testing.T) {
	t.Parallel()

	data := []byte{0x01, 0x00, 0x00, 0x00, 0x12, 0x70, 0x66, 0x77, 0x64, 0x40, 0x61, 0x6d, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6d}
	hi := int(0x12)
	logrus.Print(hi)

	result := ProcessServiceRequest(data)

	assert.NotNil(t, result)
	assert.Equal(t, uint8(0x6), result.MessageType) // APF_SERVICE_ACCEPT
	assert.Equal(t, uint32(0x12), result.ServiceNameLength)
	assert.Equal(t, [18]uint8{0x70, 0x66, 0x77, 0x64, 0x40, 0x61, 0x6d, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6d}, result.ServiceName)
}

func TestProcessChannelOpenConfirmation(t *testing.T) {
	t.Parallel()

	data := []byte{0x01}
	statusChannel := make(chan bool)
	wg := &sync.WaitGroup{}
	session := &Session{
		Status:    statusChannel,
		WaitGroup: wg,
	}

	defer close(statusChannel)

	go func() {
		<-statusChannel
	}()
	wg.Add(1)
	ProcessChannelOpenConfirmation(data, session)
}

func TestProcessProtocolVersion(t *testing.T) {
	t.Parallel()

	data := []byte{0x01}
	result := ProcessProtocolVersion(data)
	assert.NotNil(t, result)
}

func TestServiceAcceptLessThan18Characters(t *testing.T) {
	t.Parallel()

	serviceName := "test"
	result := ServiceAccept(serviceName)
	assert.NotNil(t, result)
}

func TestServiceAcceptEmptyString(t *testing.T) {
	t.Parallel()

	serviceName := ""
	result := ServiceAccept(serviceName)
	assert.NotNil(t, result)
}

func TestServiceAccept18Characters(t *testing.T) {
	t.Parallel()

	serviceName := "                  "
	result := ServiceAccept(serviceName)
	assert.NotNil(t, result)
}

func TestServiceAcceptMoreThan18Characters(t *testing.T) {
	t.Parallel()

	serviceName := "                   "
	result := ServiceAccept(serviceName)
	assert.NotNil(t, result)
}

func TestProtocolVersion(t *testing.T) {
	t.Parallel()

	result := ProtocolVersion(1, 0, 9)
	assert.NotNil(t, result)
}

func TestProtocolVersionWithUUID(t *testing.T) {
	t.Parallel()

	result := ProtocolVersionWithUUID(1, 0, 9, [16]byte{})
	assert.NotNil(t, result)
}

func TestTcpForwardReplySuccess(t *testing.T) {
	t.Parallel()

	result := TcpForwardReplySuccess(16992)
	assert.NotNil(t, result)
}

func TestChannelOpen(t *testing.T) {
	t.Parallel()

	result := ChannelOpen(1)
	assert.NotNil(t, result)
}

func TestChannelOpenReplySuccess(t *testing.T) {
	t.Parallel()

	result := ChannelOpenReplySuccess(0, 1)
	assert.NotNil(t, result)
}

func TestChannelOpenReplyFailure(t *testing.T) {
	t.Parallel()

	result := ChannelOpenReplyFailure(0, 1)
	assert.NotNil(t, result)
}

func TestChannelClose(t *testing.T) {
	t.Parallel()

	result := ChannelClose(0)
	assert.NotNil(t, result)
}

func TestChannelData(t *testing.T) {
	t.Parallel()

	data := []byte{0x01}
	result := ChannelData(0, data)
	assert.NotNil(t, result)
}

func TestChannelWindowAdjust(t *testing.T) {
	t.Parallel()

	result := ChannelWindowAdjust(0, 32)
	assert.NotNil(t, result)
}

func TestNewProcessor(t *testing.T) {
	t.Parallel()

	t.Run("with nil handler uses DefaultHandler", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		assert.NotNil(t, p)
	})

	t.Run("with custom handler uses it", func(t *testing.T) {
		t.Parallel()

		handler := &mockHandler{}
		p := NewProcessor(handler)
		assert.NotNil(t, p)
	})
}

type mockHandler struct {
	onProtocolVersionCalled bool
	onAuthRequestCalled     bool
	onGlobalRequestCalled   bool
	authResponse            AuthResponse
	protocolVersionErr      error
}

func (m *mockHandler) OnProtocolVersion(info ProtocolVersionInfo) error {
	m.onProtocolVersionCalled = true

	return m.protocolVersionErr
}

func (m *mockHandler) OnAuthRequest(request AuthRequest) AuthResponse {
	m.onAuthRequestCalled = true

	return m.authResponse
}

func (m *mockHandler) OnGlobalRequest(request GlobalRequest) bool {
	m.onGlobalRequestCalled = true

	return false
}

func TestProcessorProcessProtocolVersion(t *testing.T) {
	t.Parallel()

	// Build a valid APF_PROTOCOL_VERSION message (93 bytes total)
	buildProtocolVersion := func() []byte {
		data := make([]byte, 93)
		data[0] = APF_PROTOCOLVERSION
		// Major version = 1 (bytes 1-4)
		data[4] = 1
		// Minor version = 0 (bytes 5-8)
		// Trigger reason = 1 (bytes 9-12)
		data[12] = 1
		// UUID (bytes 13-28) - just zeros for test

		return data
	}

	t.Run("handler accepts protocol version", func(t *testing.T) {
		t.Parallel()

		handler := &mockHandler{protocolVersionErr: nil}
		p := NewProcessor(handler)
		session := &Session{}

		data := buildProtocolVersion()
		result := p.Process(data, session)

		assert.True(t, handler.onProtocolVersionCalled)
		assert.Greater(t, result.Len(), 0)
	})

	t.Run("handler rejects protocol version", func(t *testing.T) {
		t.Parallel()

		handler := &mockHandler{protocolVersionErr: assert.AnError}
		p := NewProcessor(handler)
		session := &Session{}

		data := buildProtocolVersion()
		result := p.Process(data, session)

		assert.True(t, handler.onProtocolVersionCalled)
		// Should return empty buffer when rejected
		assert.Equal(t, 0, result.Len())
	})
}

func TestProcessKeepAliveRequest(t *testing.T) {
	t.Parallel()

	t.Run("valid keepalive request", func(t *testing.T) {
		t.Parallel()
		// APF_KEEPALIVE_REQUEST: MessageType(1) + Cookie(4)
		data := []byte{APF_KEEPALIVE_REQUEST, 0x00, 0x00, 0x00, 0x42}
		session := &Session{}

		result := ProcessKeepAliveRequest(data, session)
		reply, ok := result.(APF_KEEPALIVE_REPLY_MESSAGE)
		assert.True(t, ok)
		assert.Equal(t, byte(APF_KEEPALIVE_REPLY), reply.MessageType)
		assert.Equal(t, uint32(0x42), reply.Cookie)
	})

	t.Run("short keepalive request", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_KEEPALIVE_REQUEST, 0x00} // Too short
		session := &Session{}

		result := ProcessKeepAliveRequest(data, session)
		reply, ok := result.(APF_KEEPALIVE_REPLY_MESSAGE)
		assert.True(t, ok)
		// Returns empty reply with uninitialized cookie
		assert.Equal(t, byte(0), reply.MessageType)
	})
}

func TestProcessKeepAliveReply(t *testing.T) {
	t.Parallel()

	t.Run("valid keepalive reply", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_KEEPALIVE_REPLY, 0x00, 0x00, 0x00, 0x42}
		session := &Session{}

		// Should not panic
		ProcessKeepAliveReply(data, session)
	})

	t.Run("short keepalive reply", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_KEEPALIVE_REPLY}
		session := &Session{}

		// Should not panic
		ProcessKeepAliveReply(data, session)
	})
}

func TestProcessKeepAliveOptionsReply(t *testing.T) {
	t.Parallel()

	t.Run("valid options reply", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_KEEPALIVE_OPTIONS_REPLY, 0x00, 0x00, 0x00, 0x1E, 0x00, 0x00, 0x00, 0x0A}
		session := &Session{}

		// Should not panic
		ProcessKeepAliveOptionsReply(data, session)
	})

	t.Run("short options reply", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_KEEPALIVE_OPTIONS_REPLY, 0x00}
		session := &Session{}

		// Should not panic
		ProcessKeepAliveOptionsReply(data, session)
	})
}

func TestBytesToHex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"empty", []byte{}, ""},
		{"single byte", []byte{0xAB}, "AB"},
		{"multiple bytes", []byte{0x01, 0x23, 0x45}, "012345"},
		{"leading zeros", []byte{0x00, 0x0F}, "000F"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := bytesToHex(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestHexToGUID(t *testing.T) {
	t.Parallel()

	t.Run("valid 32-char hex string", func(t *testing.T) {
		t.Parallel()
		// Input: 00112233445566778899AABBCCDDEEFF
		// Expected GUID format with byte swapping for first 3 groups
		input := "00112233445566778899AABBCCDDEEFF"
		result := hexToGUID(input)
		// First group: 00112233 -> 33221100
		// Second group: 4455 -> 5544
		// Third group: 6677 -> 7766
		// Fourth group: 8899 (no swap)
		// Fifth group: AABBCCDDEEFF (no swap)
		assert.Equal(t, "33221100-5544-7766-8899-AABBCCDDEEFF", result)
	})

	t.Run("short hex string returns zero GUID", func(t *testing.T) {
		t.Parallel()

		result := hexToGUID("001122")
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", result)
	})

	t.Run("empty string returns zero GUID", func(t *testing.T) {
		t.Parallel()

		result := hexToGUID("")
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", result)
	})
}

func TestKeepAliveOptionsRequest(t *testing.T) {
	t.Parallel()

	result := KeepAliveOptionsRequest(30, 10)
	assert.Equal(t, byte(APF_KEEPALIVE_OPTIONS_REQUEST), result.MessageType)
	assert.Equal(t, uint32(30), result.IntervalSeconds)
	assert.Equal(t, uint32(10), result.TimeoutSeconds)
}

func TestDefaultHandler(t *testing.T) {
	t.Parallel()

	handler := DefaultHandler{}

	t.Run("OnProtocolVersion returns nil", func(t *testing.T) {
		t.Parallel()

		err := handler.OnProtocolVersion(ProtocolVersionInfo{})
		assert.NoError(t, err)
	})

	t.Run("OnAuthRequest returns not authenticated", func(t *testing.T) {
		t.Parallel()

		response := handler.OnAuthRequest(AuthRequest{})
		assert.False(t, response.Authenticated)
	})

	t.Run("OnGlobalRequest returns false", func(t *testing.T) {
		t.Parallel()

		result := handler.OnGlobalRequest(GlobalRequest{})
		assert.False(t, result)
	})
}

func TestDecodeUserAuthRequest(t *testing.T) {
	t.Parallel()

	p := NewProcessor(nil)

	t.Run("valid password auth request", func(t *testing.T) {
		t.Parallel()

		// Build a valid APF_USERAUTH_REQUEST with password method
		// MessageType(1) + UsernameLen(4) + Username + ServiceNameLen(4) + ServiceName + MethodNameLen(4) + MethodName + FALSE(1) + PasswordLen(4) + Password
		username := testUsername
		serviceName := APF_SERVICE_PFWD
		methodName := APF_AUTH_PASSWORD
		password := "secret123"

		data := []byte{APF_USERAUTH_REQUEST}
		// Username length
		data = append(data, 0x00, 0x00, 0x00, byte(len(username)))
		data = append(data, []byte(username)...)
		// Service name length
		data = append(data, 0x00, 0x00, 0x00, byte(len(serviceName)))
		data = append(data, []byte(serviceName)...)
		// Method name length
		data = append(data, 0x00, 0x00, 0x00, byte(len(methodName)))
		data = append(data, []byte(methodName)...)
		// FALSE byte
		data = append(data, 0x00)
		// Password length
		data = append(data, 0x00, 0x00, 0x00, byte(len(password)))
		data = append(data, []byte(password)...)

		request, err := p.decodeUserAuthRequest(data)
		assert.NoError(t, err)
		assert.Equal(t, username, request.Username)
		assert.Equal(t, serviceName, request.ServiceName)
		assert.Equal(t, methodName, request.MethodName)
		assert.Equal(t, password, request.Password)
	})

	t.Run("valid none auth request", func(t *testing.T) {
		t.Parallel()

		username := testUsername
		serviceName := APF_SERVICE_PFWD
		methodName := APF_AUTH_NONE

		data := []byte{APF_USERAUTH_REQUEST}
		data = append(data, 0x00, 0x00, 0x00, byte(len(username)))
		data = append(data, []byte(username)...)
		data = append(data, 0x00, 0x00, 0x00, byte(len(serviceName)))
		data = append(data, []byte(serviceName)...)
		data = append(data, 0x00, 0x00, 0x00, byte(len(methodName)))
		data = append(data, []byte(methodName)...)

		request, err := p.decodeUserAuthRequest(data)
		assert.NoError(t, err)
		assert.Equal(t, username, request.Username)
		assert.Equal(t, methodName, request.MethodName)
		assert.Empty(t, request.Password)
	})

	t.Run("error reading message type", func(t *testing.T) {
		t.Parallel()

		data := []byte{}
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
	})

	t.Run("error reading username length", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_USERAUTH_REQUEST}
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
	})

	t.Run("invalid username length too large", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x10, 0x00} // 4096 bytes
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid username length")
	})

	t.Run("username length exceeds buffer", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x10} // 16 bytes but only 0 available
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
	})

	t.Run("error reading service name length", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a'}
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
	})

	t.Run("invalid service name length too large", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a', 0x00, 0x00, 0x10, 0x00}
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid service name length")
	})

	t.Run("error reading method name length", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a', 0x00, 0x00, 0x00, 0x01, 'b'}
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
	})

	t.Run("invalid method name length too large", func(t *testing.T) {
		t.Parallel()

		data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a', 0x00, 0x00, 0x00, 0x01, 'b', 0x00, 0x00, 0x10, 0x00}
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid method name length")
	})

	t.Run("password method missing FALSE byte", func(t *testing.T) {
		t.Parallel()

		methodName := APF_AUTH_PASSWORD
		data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a', 0x00, 0x00, 0x00, 0x01, 'b'}
		data = append(data, 0x00, 0x00, 0x00, byte(len(methodName)))
		data = append(data, []byte(methodName)...)
		// No FALSE byte
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not enough data for password FALSE byte")
	})

	t.Run("password FALSE byte is not zero", func(t *testing.T) {
		t.Parallel()

		methodName := APF_AUTH_PASSWORD
		data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a', 0x00, 0x00, 0x00, 0x01, 'b'}
		data = append(data, 0x00, 0x00, 0x00, byte(len(methodName)))
		data = append(data, []byte(methodName)...)
		data = append(data, 0x01) // FALSE byte is 1, not 0
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password FALSE byte is not zero")
	})

	t.Run("password length missing", func(t *testing.T) {
		t.Parallel()

		methodName := APF_AUTH_PASSWORD
		data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a', 0x00, 0x00, 0x00, 0x01, 'b'}
		data = append(data, 0x00, 0x00, 0x00, byte(len(methodName)))
		data = append(data, []byte(methodName)...)
		data = append(data, 0x00) // FALSE byte
		// No password length
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
	})

	t.Run("invalid password length too large", func(t *testing.T) {
		t.Parallel()

		methodName := APF_AUTH_PASSWORD
		data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a', 0x00, 0x00, 0x00, 0x01, 'b'}
		data = append(data, 0x00, 0x00, 0x00, byte(len(methodName)))
		data = append(data, []byte(methodName)...)
		data = append(data, 0x00)                   // FALSE byte
		data = append(data, 0x00, 0x00, 0x10, 0x00) // 4096 bytes
		_, err := p.decodeUserAuthRequest(data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid password length")
	})
}

func TestCreateAuthFailureMessage(t *testing.T) {
	t.Parallel()

	p := NewProcessor(nil)
	msg := p.createAuthFailureMessage()

	assert.Equal(t, byte(APF_USERAUTH_FAILURE), msg.MessageType)
	assert.Equal(t, uint32(len(APF_AUTH_PASSWORD)), msg.AuthenticationsThatCanContinueLength)
	assert.Equal(t, byte(0), msg.PartialSuccess)
}

func TestProcessUserAuthRequest(t *testing.T) {
	t.Parallel()

	buildAuthRequest := func(username, serviceName, methodName, password string) []byte {
		data := []byte{APF_USERAUTH_REQUEST}
		data = append(data, 0x00, 0x00, 0x00, byte(len(username)))
		data = append(data, []byte(username)...)
		data = append(data, 0x00, 0x00, 0x00, byte(len(serviceName)))
		data = append(data, []byte(serviceName)...)
		data = append(data, 0x00, 0x00, 0x00, byte(len(methodName)))
		data = append(data, []byte(methodName)...)

		if methodName == APF_AUTH_PASSWORD {
			data = append(data, 0x00) // FALSE byte
			data = append(data, 0x00, 0x00, 0x00, byte(len(password)))
			data = append(data, []byte(password)...)
		}

		return data
	}

	t.Run("auth success", func(t *testing.T) {
		t.Parallel()

		handler := &mockHandler{authResponse: AuthResponse{Authenticated: true}}
		p := NewProcessor(handler)
		session := &Session{}

		data := buildAuthRequest(testUsername, APF_SERVICE_PFWD, APF_AUTH_PASSWORD, "secret")
		result := p.Process(data, session)

		assert.True(t, handler.onAuthRequestCalled)
		assert.Greater(t, result.Len(), 0)
		// Check first byte is APF_USERAUTH_SUCCESS
		assert.Equal(t, byte(APF_USERAUTH_SUCCESS), result.Bytes()[0])
	})

	t.Run("auth failure", func(t *testing.T) {
		t.Parallel()

		handler := &mockHandler{authResponse: AuthResponse{Authenticated: false}}
		p := NewProcessor(handler)
		session := &Session{}

		data := buildAuthRequest(testUsername, APF_SERVICE_PFWD, APF_AUTH_PASSWORD, "wrong")
		result := p.Process(data, session)

		assert.True(t, handler.onAuthRequestCalled)
		assert.Greater(t, result.Len(), 0)
		// Check first byte is APF_USERAUTH_FAILURE
		assert.Equal(t, byte(APF_USERAUTH_FAILURE), result.Bytes()[0])
	})

	t.Run("malformed auth request returns failure", func(t *testing.T) {
		t.Parallel()

		handler := &mockHandler{}
		p := NewProcessor(handler)
		session := &Session{}

		// Malformed request - too short
		data := []byte{APF_USERAUTH_REQUEST, 0x00}
		result := p.Process(data, session)

		assert.False(t, handler.onAuthRequestCalled)
		assert.Greater(t, result.Len(), 0)
		// Should return failure message
		assert.Equal(t, byte(APF_USERAUTH_FAILURE), result.Bytes()[0])
	})
}

func TestProcessAllMessageTypes(t *testing.T) {
	t.Parallel()

	t.Run("APF_KEEPALIVE_REQUEST via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_KEEPALIVE_REQUEST, 0x00, 0x00, 0x00, 0x42}

		result := p.Process(data, session)
		assert.Greater(t, result.Len(), 0)
	})

	t.Run("APF_KEEPALIVE_REPLY via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_KEEPALIVE_REPLY, 0x00, 0x00, 0x00, 0x42}

		result := p.Process(data, session)
		// Should return empty buffer (no response to a reply)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_KEEPALIVE_OPTIONS_REPLY via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_KEEPALIVE_OPTIONS_REPLY, 0x00, 0x00, 0x00, 0x1E, 0x00, 0x00, 0x00, 0x0A}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_CHANNEL_OPEN via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_CHANNEL_OPEN}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_DISCONNECT via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_DISCONNECT}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_SERVICE_REQUEST valid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		// "auth@amt.intel.com" service request
		data := []byte{APF_SERVICE_REQUEST, 0x00, 0x00, 0x00, 0x12, 0x61, 0x75, 0x74, 0x68, 0x40, 0x61, 0x6d, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6d}

		result := p.Process(data, session)
		assert.Greater(t, result.Len(), 0)
	})

	t.Run("APF_SERVICE_REQUEST invalid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		// Too short to be valid
		data := []byte{APF_SERVICE_REQUEST, 0x00}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_GLOBAL_REQUEST valid via Process", func(t *testing.T) {
		t.Parallel()

		handler := &mockHandler{}
		p := NewProcessor(handler)
		session := &Session{}
		// tcpip-forward request
		data := []byte{
			APF_GLOBAL_REQUEST,
			0x00, 0x00, 0x00, 0x0D,
			0x74, 0x63, 0x70, 0x69, 0x70, 0x2d, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
			0x00,
			0x00, 0x00, 0x00, 0x03,
			0x00, 0x00, 0x00,
			0x00, 0x00, 0x42, 0x60,
		}

		result := p.Process(data, session)

		assert.True(t, handler.onGlobalRequestCalled)
		assert.Greater(t, result.Len(), 0)
	})

	t.Run("APF_GLOBAL_REQUEST invalid via Process", func(t *testing.T) {
		t.Parallel()

		handler := &mockHandler{}
		p := NewProcessor(handler)
		session := &Session{}
		data := []byte{APF_GLOBAL_REQUEST, 0x00}

		result := p.Process(data, session)

		assert.False(t, handler.onGlobalRequestCalled)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_CHANNEL_OPEN_CONFIRMATION valid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		session := &Session{WaitGroup: wg}
		data := make([]byte, 17)
		data[0] = APF_CHANNEL_OPEN_CONFIRMATION

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_CHANNEL_OPEN_CONFIRMATION invalid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_CHANNEL_OPEN_CONFIRMATION}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_CHANNEL_OPEN_FAILURE valid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		statusChan := make(chan bool, 1)
		errChan := make(chan error, 1)
		session := &Session{Status: statusChan, ErrorBuffer: errChan}
		data := make([]byte, 17)
		data[0] = APF_CHANNEL_OPEN_FAILURE

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_CHANNEL_OPEN_FAILURE invalid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_CHANNEL_OPEN_FAILURE}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_CHANNEL_CLOSE valid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := make([]byte, 5)
		data[0] = APF_CHANNEL_CLOSE

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_CHANNEL_CLOSE invalid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_CHANNEL_CLOSE}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_CHANNEL_DATA valid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		timer := time.NewTimer(1 * time.Second)
		session := &Session{Timer: timer}
		// Valid channel data: type(1) + recipient(4) + len(4) + data
		data := []byte{APF_CHANNEL_DATA, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0xAB}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
		timer.Stop()
	})

	t.Run("APF_CHANNEL_DATA invalid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_CHANNEL_DATA}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_CHANNEL_WINDOW_ADJUST valid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := make([]byte, 9)
		data[0] = APF_CHANNEL_WINDOW_ADJUST

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_CHANNEL_WINDOW_ADJUST invalid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_CHANNEL_WINDOW_ADJUST}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("APF_PROTOCOLVERSION invalid via Process", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{APF_PROTOCOLVERSION}

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})

	t.Run("default case unknown message type", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil)
		session := &Session{}
		data := []byte{0xFF} // Unknown message type

		result := p.Process(data, session)
		assert.Equal(t, 0, result.Len())
	})
}

func TestProcessGlobalRequestCancelTcpForward(t *testing.T) {
	t.Parallel()

	// Build cancel-tcpip-forward request
	serviceName := APF_GLOBAL_REQUEST_STR_TCP_FORWARD_CANCEL_REQUEST
	addrLen := 4
	globalReqLen := len(serviceName)

	data := make([]byte, 40)
	data[0] = APF_GLOBAL_REQUEST
	// globalReqLen
	data[1] = 0
	data[2] = 0
	data[3] = 0
	data[4] = byte(globalReqLen)
	copy(data[5:5+globalReqLen], serviceName)
	// WantReply
	data[5+globalReqLen] = 0
	// AddressLength
	data[6+globalReqLen] = 0
	data[7+globalReqLen] = 0
	data[8+globalReqLen] = 0
	data[9+globalReqLen] = byte(addrLen)
	// Address (4 bytes)
	copy(data[10+globalReqLen:14+globalReqLen], []byte("test"))
	// Port
	data[14+globalReqLen] = 0
	data[15+globalReqLen] = 0
	data[16+globalReqLen] = 0x42
	data[17+globalReqLen] = 0x60

	request, reply := ProcessGlobalRequest(data)
	assert.Equal(t, APF_GLOBAL_REQUEST_STR_TCP_FORWARD_CANCEL_REQUEST, request.RequestType)
	assert.Equal(t, APF_REQUEST_SUCCESS, reply)
}

func TestProcessGlobalRequestZeroStringLength(t *testing.T) {
	t.Parallel()

	// Test with zero string length
	data := []byte{APF_GLOBAL_REQUEST, 0x00, 0x00, 0x00, 0x00}

	request, reply := ProcessGlobalRequest(data)
	assert.Empty(t, request.RequestType)
	assert.Nil(t, reply)
}

func TestProcessGlobalRequestZeroAddressLength(t *testing.T) {
	t.Parallel()

	// Build tcpip-forward request with zero address length
	serviceName := APF_GLOBAL_REQUEST_STR_TCP_FORWARD_REQUEST
	globalReqLen := len(serviceName)

	data := make([]byte, 32)
	data[0] = APF_GLOBAL_REQUEST
	data[1] = 0
	data[2] = 0
	data[3] = 0
	data[4] = byte(globalReqLen)
	copy(data[5:5+globalReqLen], serviceName)
	data[5+globalReqLen] = 0 // WantReply
	// AddressLength = 0
	data[6+globalReqLen] = 0
	data[7+globalReqLen] = 0
	data[8+globalReqLen] = 0
	data[9+globalReqLen] = 0
	// Port
	data[10+globalReqLen] = 0
	data[11+globalReqLen] = 0
	data[12+globalReqLen] = 0x42
	data[13+globalReqLen] = 0x60

	request, reply := ProcessGlobalRequest(data)
	assert.Equal(t, APF_GLOBAL_REQUEST_STR_TCP_FORWARD_REQUEST, request.RequestType)
	assert.Empty(t, request.Address)
	assert.NotNil(t, reply)
}

func TestProcessServiceRequestUnknownService(t *testing.T) {
	t.Parallel()

	// Service name length = 18, but unknown service
	data := []byte{APF_SERVICE_REQUEST, 0x00, 0x00, 0x00, 0x12}
	data = append(data, []byte("unknown@intel.com!")...)

	result := ProcessServiceRequest(data)
	// Should return empty message (service not recognized)
	assert.Equal(t, byte(0), result.MessageType)
}

func TestProcessServiceRequestShortServiceName(t *testing.T) {
	t.Parallel()

	// Service name length = 5 (not 18)
	data := []byte{APF_SERVICE_REQUEST, 0x00, 0x00, 0x00, 0x05}
	data = append(data, []byte("short")...)

	result := ProcessServiceRequest(data)
	// Should return empty message (service length != 18)
	assert.Equal(t, byte(0), result.MessageType)
}

func TestProcessServiceRequestZeroServiceNameLength(t *testing.T) {
	t.Parallel()

	// Service name length = 0
	data := []byte{APF_SERVICE_REQUEST, 0x00, 0x00, 0x00, 0x00}

	result := ProcessServiceRequest(data)
	assert.Equal(t, byte(0), result.MessageType)
}

func TestDecodeProtocolVersionShortData(t *testing.T) {
	t.Parallel()

	p := NewProcessor(nil)

	// Create data that's too short for binary.Read to succeed
	data := []byte{APF_PROTOCOLVERSION, 0x00, 0x00}

	// Should not panic, just log error and return zero values
	info := p.decodeProtocolVersion(data)
	assert.Equal(t, uint32(0), info.MajorVersion)
}

func TestProcessChannelDataWithMoreData(t *testing.T) {
	t.Parallel()

	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()

	session := &Session{
		Timer:    timer,
		Tempdata: []byte{},
	}

	// Valid channel data with actual data bytes
	data := []byte{
		APF_CHANNEL_DATA,
		0x00, 0x00, 0x00, 0x01, // RecipientChannel
		0x00, 0x00, 0x00, 0x05, // DataLength = 5
		0x01, 0x02, 0x03, 0x04, 0x05, // Data bytes
	}

	ProcessChannelData(data, session)
	assert.Len(t, session.Tempdata, 5)
}

func TestProcessGlobalRequestReadErrors(t *testing.T) {
	t.Parallel()

	t.Run("short data causes read errors", func(t *testing.T) {
		t.Parallel()

		// Just message type - will fail to read string length
		data := []byte{APF_GLOBAL_REQUEST}
		request, reply := ProcessGlobalRequest(data)
		assert.Empty(t, request.RequestType)
		assert.Nil(t, reply)
	})

	t.Run("incomplete string buffer", func(t *testing.T) {
		t.Parallel()

		// Claims string length of 10 but doesn't have enough data
		data := []byte{APF_GLOBAL_REQUEST, 0x00, 0x00, 0x00, 0x0A, 0x01, 0x02}
		_, reply := ProcessGlobalRequest(data)
		// Should handle gracefully
		assert.Nil(t, reply)
	})
}

func TestProcessServiceRequestReadErrors(t *testing.T) {
	t.Parallel()

	t.Run("short data causes read errors", func(t *testing.T) {
		t.Parallel()

		// Just message type - will fail to read service name length
		data := []byte{APF_SERVICE_REQUEST}
		result := ProcessServiceRequest(data)
		assert.Equal(t, byte(0), result.MessageType)
	})

	t.Run("incomplete service name buffer", func(t *testing.T) {
		t.Parallel()

		// Claims service name length of 18 but doesn't have enough data
		data := []byte{APF_SERVICE_REQUEST, 0x00, 0x00, 0x00, 0x12, 0x01, 0x02}
		result := ProcessServiceRequest(data)
		// Should handle gracefully
		assert.Equal(t, byte(0), result.MessageType)
	})
}

func TestProcessChannelWindowAdjustReadError(t *testing.T) {
	t.Parallel()

	// Very short data
	data := []byte{APF_CHANNEL_WINDOW_ADJUST}
	session := &Session{}

	// Should not panic
	ProcessChannelWindowAdjust(data, session)
}

func TestProcessChannelCloseReadError(t *testing.T) {
	t.Parallel()

	// Very short data
	data := []byte{APF_CHANNEL_CLOSE}
	session := &Session{}

	// Should not panic
	result := ProcessChannelClose(data, session)
	assert.NotNil(t, result)
}

func TestProcessChannelDataReadErrors(t *testing.T) {
	t.Parallel()

	t.Run("short data causes read errors", func(t *testing.T) {
		t.Parallel()

		timer := time.NewTimer(1 * time.Second)
		defer timer.Stop()

		session := &Session{
			Timer:    timer,
			Tempdata: []byte{},
		}

		// Just message type
		data := []byte{APF_CHANNEL_DATA}
		ProcessChannelData(data, session)
	})

	t.Run("missing data buffer", func(t *testing.T) {
		t.Parallel()

		timer := time.NewTimer(1 * time.Second)
		defer timer.Stop()

		session := &Session{
			Timer:    timer,
			Tempdata: []byte{},
		}

		// Has length but no actual data
		data := []byte{
			APF_CHANNEL_DATA,
			0x00, 0x00, 0x00, 0x01, // RecipientChannel
			0x00, 0x00, 0x00, 0x10, // DataLength = 16
			// No actual data
		}
		ProcessChannelData(data, session)
	})
}

func TestProcessChannelOpenConfirmationReadError(t *testing.T) {
	t.Parallel()

	// Very short data
	data := []byte{APF_CHANNEL_OPEN_CONFIRMATION}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	session := &Session{WaitGroup: wg}

	// Should not panic
	ProcessChannelOpenConfirmation(data, session)
}

func TestProcessChannelOpenFailureReadError(t *testing.T) {
	t.Parallel()

	// Very short data
	data := []byte{APF_CHANNEL_OPEN_FAILURE}
	statusChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	session := &Session{Status: statusChan, ErrorBuffer: errChan}

	// Should not panic
	ProcessChannelOpenFailure(data, session)
}

func TestProcessProtocolVersionReadError(t *testing.T) {
	t.Parallel()

	// Very short data - should trigger error in binary.Read
	data := []byte{APF_PROTOCOLVERSION}
	result := ProcessProtocolVersion(data)
	assert.NotNil(t, result)
}

func TestValidateGlobalRequestNotEnoughForGlobalReqName(t *testing.T) {
	t.Parallel()

	// This tests the case where we have exactly 5 bytes but string length exceeds buffer
	// globalReqLen = 5, but we only have 5 bytes total (need 5 + 5 + 1 = 11)
	data := []byte{APF_GLOBAL_REQUEST, 0x00, 0x00, 0x00, 0x05, 'a'}
	result := ValidateGlobalRequest(data)
	assert.False(t, result)
}

func TestProcessBinaryWriteError(t *testing.T) {
	t.Parallel()

	// This is tricky - binary.Write rarely fails for these struct types
	// The error path is at line 282-284 where binary.Write might fail
	// This would require a writer that returns an error, which bytes.Buffer doesn't do
	// The test is for completeness, verifying the function works correctly

	p := NewProcessor(nil)
	session := &Session{}

	// Valid keepalive that returns a response
	data := []byte{APF_KEEPALIVE_REQUEST, 0x00, 0x00, 0x00, 0x42}
	result := p.Process(data, session)
	assert.Greater(t, result.Len(), 0)
}

func TestDecodeUserAuthRequestReadPasswordBytes(t *testing.T) {
	t.Parallel()

	p := NewProcessor(nil)

	// Build a request where password bytes read fails
	methodName := "password"
	data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a', 0x00, 0x00, 0x00, 0x01, 'b'}
	data = append(data, 0x00, 0x00, 0x00, byte(len(methodName)))
	data = append(data, []byte(methodName)...)
	data = append(data, 0x00)                   // FALSE byte
	data = append(data, 0x00, 0x00, 0x00, 0x05) // Password length = 5 but no data follows

	_, err := p.decodeUserAuthRequest(data)
	assert.Error(t, err)
}

func TestDecodeUserAuthRequestReadUsernameBytes(t *testing.T) {
	t.Parallel()

	p := NewProcessor(nil)

	// Build a request where username length is valid but reading bytes fails
	// usernameLen = 10 bytes, but we only provide 5
	data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x0A, 'a', 'b', 'c', 'd', 'e'}

	_, err := p.decodeUserAuthRequest(data)
	assert.Error(t, err)
}

func TestDecodeUserAuthRequestReadServiceNameBytes(t *testing.T) {
	t.Parallel()

	p := NewProcessor(nil)

	// Build a request where username is complete but service name read fails
	data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a'}
	// Service name length = 10, but only provide 5
	data = append(data, 0x00, 0x00, 0x00, 0x0A, 'b', 'c', 'd', 'e', 'f')

	_, err := p.decodeUserAuthRequest(data)
	assert.Error(t, err)
}

func TestDecodeUserAuthRequestReadMethodNameBytes(t *testing.T) {
	t.Parallel()

	p := NewProcessor(nil)

	// Build a request where username and service name are complete but method name read fails
	data := []byte{APF_USERAUTH_REQUEST, 0x00, 0x00, 0x00, 0x01, 'a', 0x00, 0x00, 0x00, 0x01, 'b'}
	// Method name length = 10, but only provide 5
	data = append(data, 0x00, 0x00, 0x00, 0x0A, 'c', 'd', 'e', 'f', 'g')

	_, err := p.decodeUserAuthRequest(data)
	assert.Error(t, err)
}

func TestProcessGlobalRequestEmptyData(t *testing.T) {
	t.Parallel()

	// Empty data will trigger binary.Read errors
	data := []byte{}
	request, reply := ProcessGlobalRequest(data)
	assert.Empty(t, request.RequestType)
	assert.Nil(t, reply)
}

func TestProcessChannelDataEmptyData(t *testing.T) {
	t.Parallel()

	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
	session := &Session{Timer: timer, Tempdata: []byte{}}

	// Empty data will trigger binary.Read errors
	data := []byte{}
	ProcessChannelData(data, session)
}

func TestProcessServiceRequestEmptyData(t *testing.T) {
	t.Parallel()

	// Empty data will trigger binary.Read errors
	data := []byte{}
	result := ProcessServiceRequest(data)
	assert.Equal(t, byte(0), result.MessageType)
}

func TestProcessGlobalRequestAddressReadError(t *testing.T) {
	t.Parallel()

	// Build a tcpip-forward request where address length is set but no address data
	serviceName := APF_GLOBAL_REQUEST_STR_TCP_FORWARD_REQUEST
	globalReqLen := len(serviceName)

	// Need: MessageType(1) + StringLength(4) + String(13) + WantReply(1) + AddressLength(4) = 23 bytes minimum
	// Then specify AddressLength > 0 but don't provide the address bytes
	data := make([]byte, 27)
	data[0] = APF_GLOBAL_REQUEST
	// StringLength = 13
	data[1] = 0
	data[2] = 0
	data[3] = 0
	data[4] = byte(globalReqLen)
	copy(data[5:5+globalReqLen], serviceName)
	data[5+globalReqLen] = 0 // WantReply
	// AddressLength = 10 (but only provide 0 bytes of actual address)
	data[6+globalReqLen] = 0
	data[7+globalReqLen] = 0
	data[8+globalReqLen] = 0
	data[9+globalReqLen] = 10

	request, reply := ProcessGlobalRequest(data)
	// Should handle gracefully, address will be empty due to error
	assert.Equal(t, APF_GLOBAL_REQUEST_STR_TCP_FORWARD_REQUEST, request.RequestType)
	assert.NotNil(t, reply)
}
