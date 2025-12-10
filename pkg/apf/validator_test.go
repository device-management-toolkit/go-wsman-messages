package apf

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateProtocolVersion(t *testing.T) {
	testCases := []struct {
		name string
		len  int
		want bool
	}{
		{
			name: "length < 93",
			len:  92,
			want: false,
		},
		{
			name: "length = 93",
			len:  93,
			want: true,
		},
		{
			name: "length > 93",
			len:  100,
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := make([]byte, tc.len)
			got := ValidateProtocolVersion(data)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestValidateServiceRequest(t *testing.T) {
	testCases := []struct {
		name       string
		dataLen    int
		serviceLen uint32
		want       bool
	}{
		{
			name:    "length < 5",
			dataLen: 4,
			want:    false,
		},
		{
			name:       "exactly 5, serviceLen=0",
			dataLen:    5,
			serviceLen: 0,
			want:       true,
		},
		{
			name:       "exactly 5, serviceLen=1 (not enough)",
			dataLen:    5,
			serviceLen: 1,
			want:       false,
		},
		{
			name:       "length=6, serviceLen=1 (enough)",
			dataLen:    6,
			serviceLen: 1,
			want:       true,
		},
		{
			name:       "length=10, serviceLen=5 (enough)",
			dataLen:    10,
			serviceLen: 5,
			want:       true,
		},
		{
			name:       "length=9, serviceLen=5 (not enough)",
			dataLen:    9,
			serviceLen: 5,
			want:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := make([]byte, tc.dataLen)
			if tc.dataLen >= 5 {
				binary.BigEndian.PutUint32(data[1:5], tc.serviceLen)
			}

			got := ValidateServiceRequest(data)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestValidateGlobalRequest(t *testing.T) {
	testCases := []struct {
		name        string
		serviceName string
		addrLen     int
		totalLen    int
		dataSetup   func([]byte, int, string, int)
		want        bool
	}{
		{
			name:        "length < 5",
			serviceName: "",
			totalLen:    4, // Not enough length to even read the globalReqLen
			dataSetup: func(data []byte, totalLen int, service string, addrLen int) {
				// No writes, as we know this will fail due to length anyway.
			},
			want: false,
		},
		{
			name:        "serviceName empty",
			serviceName: "",
			totalLen:    5, // Enough to write [1:5]
			dataSetup: func(data []byte, totalLen int, service string, addrLen int) {
				binary.BigEndian.PutUint32(data[1:5], 0)
			},
			want: false,
		},
		{
			name:        "not enough for string plus one byte",
			serviceName: "test",
			totalLen:    9, // 5 bytes header + 4 bytes for "test" = 9, but need 5+4+1=10
			dataSetup: func(data []byte, totalLen int, service string, addrLen int) {
				globalReqLen := len(service)
				binary.BigEndian.PutUint32(data[1:5], uint32(globalReqLen))
				copy(data[5:5+globalReqLen], service)
			},
			want: false,
		},
		{
			name:        "non-matching service name",
			serviceName: "unknown",
			// globalReqLen = len("unknown") = 7
			// Minimum length for reading globalReqLen and serviceName and WantReply: 5 + 7 + 1 = 13
			totalLen: 13,
			dataSetup: func(data []byte, totalLen int, service string, addrLen int) {
				globalReqLen := len(service)
				binary.BigEndian.PutUint32(data[1:5], uint32(globalReqLen))
				copy(data[5:5+globalReqLen], service)
			},
			want: false,
		},
		{
			name:        "tcpip-forward insufficient for addrLen",
			serviceName: APF_GLOBAL_REQUEST_STR_TCP_FORWARD_REQUEST,
			// globalReqLen = 13
			// Need to pass line 30: len(data) >= 5+13+1 = 19
			// But fail at line 37: len(data) < 6+13+4 = 23
			// So we need 19 <= totalLen < 23
			totalLen: 22,
			dataSetup: func(data []byte, totalLen int, service string, addrLen int) {
				globalReqLen := len(service)
				binary.BigEndian.PutUint32(data[1:5], uint32(globalReqLen))
				copy(data[5:5+globalReqLen], service)
			},
			want: false,
		},
		{
			name:        "tcpip-forward with enough length",
			serviceName: APF_GLOBAL_REQUEST_STR_TCP_FORWARD_REQUEST,
			addrLen:     4,
			// For success:
			// globalReqLen = 13
			// Need at least: 14 + 13 + 4 = 31
			totalLen: 31,
			dataSetup: func(data []byte, totalLen int, service string, addrLen int) {
				globalReqLen := len(service)
				binary.BigEndian.PutUint32(data[1:5], uint32(globalReqLen))
				copy(data[5:5+globalReqLen], service)
				binary.BigEndian.PutUint32(data[6+globalReqLen:10+globalReqLen], uint32(addrLen))
				for i := 10 + globalReqLen; i < 10+globalReqLen+addrLen; i++ {
					data[i] = 0x01
				}
			},
			want: true,
		},
		{
			name:        "cancel-tcpip-forward with enough length",
			serviceName: APF_GLOBAL_REQUEST_STR_TCP_FORWARD_CANCEL_REQUEST,
			addrLen:     4,
			// globalReqLen = len("cancel-tcpip-forward")=22
			// needed: 14+22+4=40
			totalLen: 40,
			dataSetup: func(data []byte, totalLen int, service string, addrLen int) {
				globalReqLen := len(service)
				binary.BigEndian.PutUint32(data[1:5], uint32(globalReqLen))
				copy(data[5:5+globalReqLen], service)
				binary.BigEndian.PutUint32(data[6+globalReqLen:10+globalReqLen], uint32(addrLen))
				for i := 10 + globalReqLen; i < 10+globalReqLen+addrLen; i++ {
					data[i] = 0x01
				}
			},
			want: true,
		},
		{
			name:        "tcpip-forward insufficient for large addrLen",
			serviceName: APF_GLOBAL_REQUEST_STR_TCP_FORWARD_REQUEST,
			addrLen:     10,
			// For success: 14+11+10=35 required
			// Give 34 to fail
			totalLen: 34,
			dataSetup: func(data []byte, totalLen int, service string, addrLen int) {
				globalReqLen := len(service)
				binary.BigEndian.PutUint32(data[1:5], uint32(globalReqLen))
				copy(data[5:5+globalReqLen], service)
				binary.BigEndian.PutUint32(data[6+globalReqLen:10+globalReqLen], uint32(addrLen))
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := make([]byte, tc.totalLen)
			tc.dataSetup(data, tc.totalLen, tc.serviceName, tc.addrLen)
			got := ValidateGlobalRequest(data)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestValidateChannelOpenConfirmation(t *testing.T) {
	testCases := []struct {
		name string
		len  int
		want bool
	}{
		{
			name: "length < 17",
			len:  16,
			want: false,
		},
		{
			name: "length = 17",
			len:  17,
			want: true,
		},
		{
			name: "length > 17",
			len:  20,
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := make([]byte, tc.len)
			got := ValidateChannelOpenConfirmation(data)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestValidateChannelOpenFailure(t *testing.T) {
	testCases := []struct {
		name string
		len  int
		want bool
	}{
		{
			name: "length < 17",
			len:  16,
			want: false,
		},
		{
			name: "length = 17",
			len:  17,
			want: true,
		},
		{
			name: "length > 17",
			len:  20,
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := make([]byte, tc.len)
			got := ValidateChannelOpenFailure(data)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestValidateChannelClose(t *testing.T) {
	testCases := []struct {
		name string
		len  int
		want bool
	}{
		{
			name: "length < 5",
			len:  4,
			want: false,
		},
		{
			name: "length = 5",
			len:  5,
			want: true,
		},
		{
			name: "length > 5",
			len:  10,
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := make([]byte, tc.len)
			got := ValidateChannelClose(data)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestValidateChannelData(t *testing.T) {
	testCases := []struct {
		name      string
		dataLen   int
		dataField uint32
		want      bool
	}{
		{
			name:      "length < 9",
			dataLen:   8,
			dataField: 0,
			want:      false,
		},
		{
			name:      "length = 9, dataField=0",
			dataLen:   9,
			dataField: 0,
			want:      true,
		},
		{
			name:      "length=10, dataField=1",
			dataLen:   10,
			dataField: 1,
			want:      true,
		},
		{
			name:      "length=13, dataField=5 (needs 14)",
			dataLen:   13,
			dataField: 5,
			want:      false,
		},
		{
			name:      "length=14, dataField=5",
			dataLen:   14,
			dataField: 5,
			want:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := make([]byte, tc.dataLen)
			if tc.dataLen >= 9 {
				binary.BigEndian.PutUint32(data[5:9], tc.dataField)
			}

			got := ValidateChannelData(data)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestValidateChannelWindowAdjust(t *testing.T) {
	testCases := []struct {
		name string
		len  int
		want bool
	}{
		{
			name: "length < 9",
			len:  8,
			want: false,
		},
		{
			name: "length = 9",
			len:  9,
			want: true,
		},
		{
			name: "length > 9",
			len:  10,
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := make([]byte, tc.len)
			got := ValidateChannelWindowAdjust(data)
			assert.Equal(t, tc.want, got)
		})
	}
}
