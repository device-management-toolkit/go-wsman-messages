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

func TestProcessorProcessUserAuthRequest(t *testing.T) {
	t.Parallel()

	// Build a valid APF_USERAUTH_REQUEST message
	// Format: MessageType(1) + UsernameLen(4) + Username + ServiceNameLen(4) + ServiceName + MethodNameLen(4) + MethodName
	buildAuthRequest := func(username, serviceName, methodName string) []byte {
		data := []byte{APF_USERAUTH_REQUEST}
		// Username
		data = append(data, byte(len(username)>>24), byte(len(username)>>16), byte(len(username)>>8), byte(len(username)))
		data = append(data, []byte(username)...)
		// ServiceName
		data = append(data, byte(len(serviceName)>>24), byte(len(serviceName)>>16), byte(len(serviceName)>>8), byte(len(serviceName)))
		data = append(data, []byte(serviceName)...)
		// MethodName
		data = append(data, byte(len(methodName)>>24), byte(len(methodName)>>16), byte(len(methodName)>>8), byte(len(methodName)))
		data = append(data, []byte(methodName)...)

		return data
	}

	t.Run("successful authentication", func(t *testing.T) {
		t.Parallel()

		handler := &mockHandler{authResponse: AuthResponse{Authenticated: true}}
		p := NewProcessor(handler)
		session := &Session{}

		data := buildAuthRequest("testuser", "auth@amt.intel.com", "none")
		result := p.Process(data, session)

		assert.True(t, handler.onAuthRequestCalled)
		assert.Greater(t, result.Len(), 0)
		// First byte should be APF_USERAUTH_SUCCESS
		assert.Equal(t, byte(APF_USERAUTH_SUCCESS), result.Bytes()[0])
	})

	t.Run("failed authentication", func(t *testing.T) {
		t.Parallel()

		handler := &mockHandler{authResponse: AuthResponse{Authenticated: false}}
		p := NewProcessor(handler)
		session := &Session{}

		data := buildAuthRequest("testuser", "auth@amt.intel.com", "none")
		result := p.Process(data, session)

		assert.True(t, handler.onAuthRequestCalled)
		assert.Greater(t, result.Len(), 0)
		// First byte should be APF_USERAUTH_FAILURE
		assert.Equal(t, byte(APF_USERAUTH_FAILURE), result.Bytes()[0])
	})

	t.Run("default handler rejects authentication", func(t *testing.T) {
		t.Parallel()

		p := NewProcessor(nil) // Uses DefaultHandler
		session := &Session{}

		data := buildAuthRequest("testuser", "auth@amt.intel.com", "none")
		result := p.Process(data, session)

		assert.Greater(t, result.Len(), 0)
		// Default handler returns Authenticated: false, so caller must provide custom handler
		assert.Equal(t, byte(APF_USERAUTH_FAILURE), result.Bytes()[0])
	})
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
