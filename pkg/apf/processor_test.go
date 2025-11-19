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

	result := ProcessGlobalRequest(data)
	assert.NotNil(t, result)
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

	result := ProtocolVersion(1, 0, 9, [16]byte{})
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
