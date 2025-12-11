/*********************************************************************
 * Copyright (c) Intel Corporation 2022
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package apf implements the APF (AMT Port Forwarding) Protocol
package apf

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// Processor handles APF protocol messages with optional callbacks.
type Processor struct {
	handler Handler
}

// NewProcessor creates a new Processor with the given handler.
// If handler is nil, a DefaultHandler is used (maintains backward compatibility).
func NewProcessor(handler Handler) *Processor {
	if handler == nil {
		handler = DefaultHandler{}
	}

	return &Processor{handler: handler}
}

// decodeProtocolVersion extracts fields from an APF_PROTOCOL_VERSION_MESSAGE.
func (p *Processor) decodeProtocolVersion(data []byte) ProtocolVersionInfo {
	message := APF_PROTOCOL_VERSION_MESSAGE{}
	dataBuffer := bytes.NewBuffer(data)

	err := binary.Read(dataBuffer, binary.BigEndian, &message)
	if err != nil {
		log.Error(err)
	}

	// Convert UUID to GUID string format
	uuidBytes := message.UUID[:]
	hexStr := bytesToHex(uuidBytes)
	guidStr := hexToGUID(hexStr)

	return ProtocolVersionInfo{
		MajorVersion:  message.MajorVersion,
		MinorVersion:  message.MinorVersion,
		TriggerReason: message.TriggerReason,
		UUID:          guidStr,
		RawUUID:       message.UUID,
	}
}

// decodeUserAuthRequest extracts fields from an APF_USERAUTH_REQUEST message.
func (p *Processor) decodeUserAuthRequest(data []byte) (AuthRequest, error) {
	dataBuffer := bytes.NewReader(data)

	var messageType byte

	if err := binary.Read(dataBuffer, binary.BigEndian, &messageType); err != nil {
		return AuthRequest{}, err
	}

	// Read username length
	var usernameLen uint32

	if err := binary.Read(dataBuffer, binary.BigEndian, &usernameLen); err != nil {
		return AuthRequest{}, err
	}

	if usernameLen > 2048 || uint32(dataBuffer.Len()) < usernameLen {
		return AuthRequest{}, errors.New("invalid username length")
	}

	usernameBytes := make([]byte, usernameLen)
	if _, err := dataBuffer.Read(usernameBytes); err != nil {
		return AuthRequest{}, err
	}

	// Read service name length
	var serviceNameLen uint32

	if err := binary.Read(dataBuffer, binary.BigEndian, &serviceNameLen); err != nil {
		return AuthRequest{}, err
	}

	if serviceNameLen > 2048 || uint32(dataBuffer.Len()) < serviceNameLen {
		return AuthRequest{}, errors.New("invalid service name length")
	}

	serviceNameBytes := make([]byte, serviceNameLen)
	if _, err := dataBuffer.Read(serviceNameBytes); err != nil {
		return AuthRequest{}, err
	}

	// Read method name length
	var methodNameLen uint32

	if err := binary.Read(dataBuffer, binary.BigEndian, &methodNameLen); err != nil {
		return AuthRequest{}, err
	}

	if methodNameLen > 2048 || uint32(dataBuffer.Len()) < methodNameLen {
		return AuthRequest{}, errors.New("invalid method name length")
	}

	methodNameBytes := make([]byte, methodNameLen)
	if _, err := dataBuffer.Read(methodNameBytes); err != nil {
		return AuthRequest{}, err
	}

	request := AuthRequest{
		Username:    string(usernameBytes),
		ServiceName: string(serviceNameBytes),
		MethodName:  string(methodNameBytes),
	}

	// If method is password, read the password
	if request.MethodName == APF_AUTH_PASSWORD {
		if dataBuffer.Len() < 1 {
			return AuthRequest{}, errors.New("not enough data for password FALSE byte")
		}

		var passwordFalse byte

		if err := binary.Read(dataBuffer, binary.BigEndian, &passwordFalse); err != nil {
			return AuthRequest{}, err
		}

		if passwordFalse != 0 {
			return AuthRequest{}, errors.New("password FALSE byte is not zero")
		}

		var passwordLen uint32

		if err := binary.Read(dataBuffer, binary.BigEndian, &passwordLen); err != nil {
			return AuthRequest{}, err
		}

		if passwordLen > 2048 || uint32(dataBuffer.Len()) < passwordLen {
			return AuthRequest{}, errors.New("invalid password length")
		}

		passwordBytes := make([]byte, passwordLen)
		if _, err := dataBuffer.Read(passwordBytes); err != nil {
			return AuthRequest{}, err
		}

		request.Password = string(passwordBytes)
	}

	return request, nil
}

// createAuthFailureMessage creates an APF_USERAUTH_FAILURE_MESSAGE.
func (p *Processor) createAuthFailureMessage() *APF_USERAUTH_FAILURE_MESSAGE {
	var authMethods [8]byte

	copy(authMethods[:], []byte(APF_AUTH_PASSWORD))

	return &APF_USERAUTH_FAILURE_MESSAGE{
		MessageType:                          APF_USERAUTH_FAILURE,
		AuthenticationsThatCanContinueLength: uint32(len(APF_AUTH_PASSWORD)),
		AuthenticationsThatCanContinue:       authMethods,
		PartialSuccess:                       0,
	}
}

// Process handles incoming APF data with handler callbacks and returns the response.
func (p *Processor) Process(data []byte, session *Session) bytes.Buffer {
	var bin_buf bytes.Buffer

	var dataToSend interface{}

	switch data[0] {
	case APF_KEEPALIVE_REQUEST:
		log.Debug("received APF_KEEPALIVE_REQUEST")

		dataToSend = ProcessKeepAliveRequest(data, session)
	case APF_KEEPALIVE_REPLY:
		log.Debug("received APF_KEEPALIVE_REPLY")

		ProcessKeepAliveReply(data, session)
	case APF_KEEPALIVE_OPTIONS_REPLY:
		log.Debug("received APF_KEEPALIVE_OPTIONS_REPLY")

		ProcessKeepAliveOptionsReply(data, session)
	case APF_GLOBAL_REQUEST: // 80
		log.Debug("received APF_GLOBAL_REQUEST")

		if ValidateGlobalRequest(data) {
			// Decode once - get both the request info and the reply
			request, reply := ProcessGlobalRequest(data)

			// Notify handler (returns true if keep-alive should be sent after)
			p.handler.OnGlobalRequest(request)

			dataToSend = reply
		}
	case APF_CHANNEL_OPEN: // (90) Sent by Intel AMT when a channel needs to be open from Intel AMT. This is not common, but WSMAN events are a good example of channel coming from AMT.
		log.Debug("received APF_CHANNEL_OPEN")
	case APF_DISCONNECT: // (1) Intel AMT wants to completely disconnect. Not sure when this happens.
		log.Debug("received APF_DISCONNECT")
	case APF_SERVICE_REQUEST: // (5)
		log.Debug("received APF_SERVICE_REQUEST")

		if ValidateServiceRequest(data) {
			dataToSend = ProcessServiceRequest(data)
		}
	case APF_CHANNEL_OPEN_CONFIRMATION: // (91) Intel AMT confirmation to an APF_CHANNEL_OPEN request.
		log.Debug("received APF_CHANNEL_OPEN_CONFIRMATION")

		if ValidateChannelOpenConfirmation(data) {
			ProcessChannelOpenConfirmation(data, session)
		}
	case APF_CHANNEL_OPEN_FAILURE: // (92) Intel AMT rejected our connection attempt.
		log.Debug("received APF_CHANNEL_OPEN_FAILURE")

		if ValidateChannelOpenFailure(data) {
			ProcessChannelOpenFailure(data, session)
		}
	case APF_CHANNEL_CLOSE: // (97) Intel AMT is closing this channel, we need to disconnect the LMS TCP connection
		log.Debug("received APF_CHANNEL_CLOSE")

		if ValidateChannelClose(data) {
			ProcessChannelClose(data, session)
		}
	case APF_CHANNEL_DATA: // (94) Intel AMT is sending data that we must relay into an LMS TCP connection.
		log.Debug("received APF_CHANNEL_DATA")

		if ValidateChannelData(data) {
			ProcessChannelData(data, session)
		}
	case APF_CHANNEL_WINDOW_ADJUST: // 93
		log.Debug("received APF_CHANNEL_WINDOW_ADJUST")

		if ValidateChannelWindowAdjust(data) {
			ProcessChannelWindowAdjust(data, session)
		}
	case APF_PROTOCOLVERSION: // 192
		log.Debug("received APF PROTOCOL VERSION")

		if ValidateProtocolVersion(data) {
			info := p.decodeProtocolVersion(data)

			if err := p.handler.OnProtocolVersion(info); err != nil {
				log.Errorf("Protocol version rejected: %v", err)

				return bin_buf
			}

			dataToSend = ProcessProtocolVersion(data)
		}
	case APF_USERAUTH_REQUEST: // 50
		log.Debug("received APF_USERAUTH_REQUEST")

		request, err := p.decodeUserAuthRequest(data)
		if err != nil {
			log.Errorf("Failed to decode auth request: %v", err)

			dataToSend = p.createAuthFailureMessage()
		} else {
			log.Debugf("username=%s serviceName=%s methodName=%s", request.Username, request.ServiceName, request.MethodName)

			response := p.handler.OnAuthRequest(request)
			if response.Authenticated {
				dataToSend = &APF_USERAUTH_SUCCESS_MESSAGE{MessageType: APF_USERAUTH_SUCCESS}
			} else {
				dataToSend = p.createAuthFailureMessage()
			}
		}
	default:
	}

	if dataToSend != nil {
		err := binary.Write(&bin_buf, binary.BigEndian, dataToSend)
		if err != nil {
			log.Error(err)
		}
	}

	return bin_buf
}

// Process is maintained for backward compatibility.
// It uses a DefaultHandler for non CIRA use cases.
func Process(data []byte, session *Session) bytes.Buffer {
	p := NewProcessor(nil)

	return p.Process(data, session)
}

func ProcessKeepAliveRequest(data []byte, session *Session) any {
	if len(data) < 5 {
		log.Warn("APF_KEEPALIVE_REQUEST message too short")

		return APF_KEEPALIVE_REPLY_MESSAGE{}
	}

	cookie := binary.BigEndian.Uint32(data[1:5])
	log.Debugf("received APF_KEEPALIVE_REQUEST with cookie: %d", cookie)

	reply := APF_KEEPALIVE_REPLY_MESSAGE{
		MessageType: APF_KEEPALIVE_REPLY,
		Cookie:      cookie,
	}

	return reply
}

func ProcessKeepAliveReply(data []byte, session *Session) {
	if len(data) < 5 {
		log.Warn("APF_KEEPALIVE_REPLY message too short")

		return
	}

	cookie := binary.BigEndian.Uint32(data[1:5])
	log.Debugf("received APF_KEEPALIVE_REPLY with cookie: %d", cookie)
}

func ProcessKeepAliveOptionsReply(data []byte, session *Session) {
	if len(data) < 9 {
		log.Warn("APF_KEEPALIVE_OPTIONS_REPLY message too short")

		return
	}

	keepaliveInterval := binary.BigEndian.Uint32(data[1:5])
	timeout := binary.BigEndian.Uint32(data[5:9])
	log.Debugf("KEEPALIVE_OPTIONS_REPLY, Keepalive Interval=%d Timeout=%d", keepaliveInterval, timeout)
}

func ProcessChannelWindowAdjust(data []byte, session *Session) {
	adjustMessage := APF_CHANNEL_WINDOW_ADJUST_MESSAGE{}
	dataBuffer := bytes.NewBuffer(data)

	err := binary.Read(dataBuffer, binary.BigEndian, &adjustMessage)
	if err != nil {
		log.Error(err)
	}

	session.TXWindow += adjustMessage.BytesToAdd
	log.Tracef("%+v", adjustMessage)
}

func ProcessChannelClose(data []byte, session *Session) APF_CHANNEL_CLOSE_MESSAGE {
	closeMessage := APF_CHANNEL_CLOSE_MESSAGE{}
	dataBuffer := bytes.NewBuffer(data)

	err := binary.Read(dataBuffer, binary.BigEndian, &closeMessage)
	if err != nil {
		log.Error(err)
	}

	log.Tracef("%+v", closeMessage)

	return ChannelClose(closeMessage.RecipientChannel)
}

// ProcessGlobalRequest decodes the global request and returns both the decoded info and the reply.
func ProcessGlobalRequest(data []byte) (GlobalRequest, interface{}) {
	genericHeader := APF_GENERIC_HEADER{}
	dataBuffer := bytes.NewBuffer(data)

	err := binary.Read(dataBuffer, binary.BigEndian, &genericHeader.MessageType)
	if err != nil {
		log.Error(err)
	}

	err = binary.Read(dataBuffer, binary.BigEndian, &genericHeader.StringLength)
	if err != nil {
		log.Error(err)
	}

	var reply interface{}

	request := GlobalRequest{}

	if int(genericHeader.StringLength) > 0 {
		stringBuffer := make([]byte, genericHeader.StringLength)
		tcpForwardRequest := APF_TCP_FORWARD_REQUEST{}

		err = binary.Read(dataBuffer, binary.BigEndian, &stringBuffer)
		if err != nil {
			log.Error(err)
		}

		genericHeader.String = string(stringBuffer[:int(genericHeader.StringLength)])
		request.RequestType = genericHeader.String

		err = binary.Read(dataBuffer, binary.BigEndian, &tcpForwardRequest.WantReply)
		if err != nil {
			log.Error(err)
		}

		err = binary.Read(dataBuffer, binary.BigEndian, &tcpForwardRequest.AddressLength)
		if err != nil {
			log.Error(err)
		}

		if int(tcpForwardRequest.AddressLength) > 0 {
			addressBuffer := make([]byte, tcpForwardRequest.AddressLength)

			err = binary.Read(dataBuffer, binary.BigEndian, &addressBuffer)
			if err != nil {
				log.Error(err)
			}

			tcpForwardRequest.Address = string(addressBuffer[:int(tcpForwardRequest.AddressLength)])
			request.Address = tcpForwardRequest.Address
		}

		err = binary.Read(dataBuffer, binary.BigEndian, &tcpForwardRequest.Port)
		if err != nil {
			log.Error(err)
		}

		request.Port = tcpForwardRequest.Port

		log.Tracef("%+v", genericHeader)
		log.Tracef("%+v", tcpForwardRequest)

		switch genericHeader.String {
		case APF_GLOBAL_REQUEST_STR_TCP_FORWARD_REQUEST:
			reply = TcpForwardReplySuccess(tcpForwardRequest.Port)
		case APF_GLOBAL_REQUEST_STR_TCP_FORWARD_CANCEL_REQUEST:
			reply = APF_REQUEST_SUCCESS
		}
	}

	return request, reply
}

func ProcessChannelData(data []byte, session *Session) {
	channelData := APF_CHANNEL_DATA_MESSAGE{}
	buf2 := bytes.NewBuffer(data)

	err := binary.Read(buf2, binary.BigEndian, &channelData.MessageType)
	if err != nil {
		log.Error(err)
	}

	err = binary.Read(buf2, binary.BigEndian, &channelData.RecipientChannel)
	if err != nil {
		log.Error(err)
	}

	err = binary.Read(buf2, binary.BigEndian, &channelData.DataLength)
	if err != nil {
		log.Error(err)
	}

	session.RXWindow = channelData.DataLength
	dataBuffer := make([]byte, channelData.DataLength)

	err = binary.Read(buf2, binary.BigEndian, &dataBuffer)
	if err != nil {
		log.Error(err)
	}

	// log.Debug("received APF_CHANNEL_DATA - " + fmt.Sprint(channelData.DataLength))
	// log.Tracef("%+v", channelData)

	session.Tempdata = append(session.Tempdata, dataBuffer[:channelData.DataLength]...)
	// var windowAdjust APF_CHANNEL_WINDOW_ADJUST_MESSAGE
	// if session.RXWindow > 1024 { // TODO: Check this
	// 	windowAdjust = ChannelWindowAdjust(channelData.RecipientChannel, session.RXWindow)
	// 	session.RXWindow = 0
	// }

	// var windowAdjust APF_CHANNEL_WINDOW_ADJUST_MESSAGE
	// if session.RXWindow > 1024 { // TODO: Check this
	// 	windowAdjust = ChannelWindowAdjust(channelData.RecipientChannel, session.RXWindow)
	// 	session.RXWindow = 0
	// }
	// // log.Tracef("%+v", session)
	// return windowAdjust
	// return windowAdjust
	session.Timer.Reset(3 * time.Second)
}

func ProcessServiceRequest(data []byte) APF_SERVICE_ACCEPT_MESSAGE {
	service := 0
	message := APF_SERVICE_REQUEST_MESSAGE{}
	dataBuffer := bytes.NewBuffer(data)

	err := binary.Read(dataBuffer, binary.BigEndian, &message.MessageType)
	if err != nil {
		log.Error(err)
	}

	err = binary.Read(dataBuffer, binary.BigEndian, &message.ServiceNameLength)
	if err != nil {
		log.Error(err)
	}

	if int(message.ServiceNameLength) > 0 {
		serviceNameBuffer := make([]byte, message.ServiceNameLength)

		err = binary.Read(dataBuffer, binary.BigEndian, &serviceNameBuffer)
		if err != nil {
			log.Error(err)
		}

		message.ServiceName = string(serviceNameBuffer[:int(message.ServiceNameLength)])
	}

	log.Tracef("%+v", message)

	if message.ServiceNameLength == 18 {
		switch message.ServiceName {
		case APF_SERVICE_PFWD:
			service = 1
		case APF_SERVICE_AUTH:
			service = 2
		}
	}

	var serviceAccept APF_SERVICE_ACCEPT_MESSAGE

	if service > 0 {
		serviceAccept = ServiceAccept(message.ServiceName)
	}

	return serviceAccept
}

func ProcessChannelOpenConfirmation(data []byte, session *Session) {
	confirmationMessage := APF_CHANNEL_OPEN_CONFIRMATION_MESSAGE{}
	dataBuffer := bytes.NewBuffer(data)

	err := binary.Read(dataBuffer, binary.BigEndian, &confirmationMessage)
	if err != nil {
		log.Error(err)
	}

	log.Tracef("%+v", confirmationMessage)
	// replySuccess := ChannelOpenReplySuccess(confirmationMessage.RecipientChannel, confirmationMessage.SenderChannel)

	log.Trace("our channel: "+fmt.Sprint(confirmationMessage.RecipientChannel), " AMT's channel: "+fmt.Sprint(confirmationMessage.SenderChannel))
	log.Trace("initial window: " + fmt.Sprint(confirmationMessage.InitialWindowSize))
	session.SenderChannel = confirmationMessage.SenderChannel
	session.RecipientChannel = confirmationMessage.RecipientChannel
	session.TXWindow = confirmationMessage.InitialWindowSize
	session.WaitGroup.Done()
}

func ProcessChannelOpenFailure(data []byte, session *Session) {
	channelOpenFailure := APF_CHANNEL_OPEN_FAILURE_MESSAGE{}
	dataBuffer := bytes.NewBuffer(data)

	err := binary.Read(dataBuffer, binary.BigEndian, &channelOpenFailure)
	if err != nil {
		log.Error(err)
	}

	log.Tracef("%+v", channelOpenFailure)
	session.Status <- false
	session.ErrorBuffer <- errors.New("error opening APF channel, reason code: " + fmt.Sprint(channelOpenFailure.ReasonCode))
}

func ProcessProtocolVersion(data []byte) APF_PROTOCOL_VERSION_MESSAGE {
	message := APF_PROTOCOL_VERSION_MESSAGE{}
	dataBuffer := bytes.NewBuffer(data)

	err := binary.Read(dataBuffer, binary.BigEndian, &message)
	if err != nil {
		log.Error(err)
	}

	// Convert UUID from raw bytes to proper GUID format and log it
	if len(data) >= 29 {
		uuidBytes := data[13:29]
		hexStr := bytesToHex(uuidBytes)
		guidStr := hexToGUID(hexStr)
		log.Debugf("SystemId UUID: %s", guidStr)
	}

	log.Tracef("%+v", message)
	version := ProtocolVersionWithUUID(message.MajorVersion, message.MinorVersion, message.TriggerReason, message.UUID)

	return version
}

// bytesToHex converts bytes to uppercase hex string.
func bytesToHex(data []byte) string {
	return strings.ToUpper(hex.EncodeToString(data))
}

// hexToGUID converts a hex string to GUID format with proper byte swapping
// Matches the TypeScript guidToStr function.
func hexToGUID(hexStr string) string {
	if len(hexStr) < 32 {
		log.Warnf("hexToGUID: input string too short (%d chars), expected 32", len(hexStr))

		return "00000000-0000-0000-0000-000000000000"
	}

	// Rearrange according to GUID format (little-endian for first 3 groups)
	guid := hexStr[6:8] + hexStr[4:6] + hexStr[2:4] + hexStr[0:2] + "-" +
		hexStr[10:12] + hexStr[8:10] + "-" +
		hexStr[14:16] + hexStr[12:14] + "-" +
		hexStr[16:20] + "-" +
		hexStr[20:]

	return guid
}

// Send the APF service accept message to the MEI.
func ServiceAccept(serviceName string) APF_SERVICE_ACCEPT_MESSAGE {
	log.Debug("sending APF_SERVICE_ACCEPT_MESSAGE")

	if len(serviceName) != 18 {
		serviceName = fmt.Sprintf("'%-18s'", serviceName)
	}

	var test [18]byte

	copy(test[:], []byte(serviceName)[:18])

	serviceAcceptMessage := APF_SERVICE_ACCEPT_MESSAGE{
		MessageType:       APF_SERVICE_ACCEPT,
		ServiceNameLength: 18,
		ServiceName:       test,
	}

	log.Tracef("%+v", serviceAcceptMessage)

	return serviceAcceptMessage
}

func ProtocolVersion(majorversion, minorversion, triggerreason uint32) APF_PROTOCOL_VERSION_MESSAGE {
	return ProtocolVersionWithUUID(majorversion, minorversion, triggerreason, [16]byte{})
}

func ProtocolVersionWithUUID(majorversion, minorversion, triggerreason uint32, uuid [16]byte) APF_PROTOCOL_VERSION_MESSAGE {
	log.Debug("sending APF_PROTOCOL_VERSION_MESSAGE")

	protVersion := APF_PROTOCOL_VERSION_MESSAGE{}
	protVersion.MessageType = APF_PROTOCOLVERSION
	protVersion.MajorVersion = majorversion
	protVersion.MinorVersion = minorversion
	protVersion.TriggerReason = triggerreason
	protVersion.UUID = uuid

	// Log the UUID as GUID format
	hexStr := bytesToHex(uuid[:])
	guidStr := hexToGUID(hexStr)
	log.Debugf("Sending SystemId UUID: %s", guidStr)

	log.Tracef("%+v", protVersion)

	return protVersion
}

func KeepAliveOptionsRequest(keepAliveTime, timeout uint32) APF_KEEPALIVE_OPTIONS_REQUEST_MESSAGE {
	log.Debug("sending APF_KEEPALIVE_OPTIONS_REQUEST_MESSAGE")

	message := APF_KEEPALIVE_OPTIONS_REQUEST_MESSAGE{
		MessageType:     APF_KEEPALIVE_OPTIONS_REQUEST,
		IntervalSeconds: keepAliveTime,
		TimeoutSeconds:  timeout,
	}

	return message
}

func TcpForwardReplySuccess(port uint32) APF_TCP_FORWARD_REPLY_MESSAGE {
	log.Debug("sending APF_TCP_FORWARD_REPLY_MESSAGE")

	message := APF_TCP_FORWARD_REPLY_MESSAGE{
		MessageType: APF_REQUEST_SUCCESS,
		PortBound:   port,
	}

	log.Tracef("%+v", message)

	return message
}

func ChannelOpen(senderChannel int) bytes.Buffer {
	var channelType [15]byte

	copy(channelType[:], []byte(APF_OPEN_CHANNEL_REQUEST_FORWARDED)[:15])

	var address [3]byte

	copy(address[:], []byte("::1")[:3])

	openMessage := APF_CHANNEL_OPEN_MESSAGE{
		MessageType:               APF_CHANNEL_OPEN,
		ChannelTypeLength:         15,
		ChannelType:               channelType,
		SenderChannel:             uint32(senderChannel), // hmm
		Reserved:                  0xFFFFFFFF,
		InitialWindowSize:         LME_RX_WINDOW_SIZE,
		ConnectedAddressLength:    3,
		ConnectedAddress:          address,
		ConnectedPort:             16992,
		OriginatorIPAddressLength: 3,
		OriginatorIPAddress:       address,
		OriginatorPort:            123,
	}

	log.Tracef("%+v", openMessage)

	var bin_buf bytes.Buffer

	err := binary.Write(&bin_buf, binary.BigEndian, openMessage)
	if err != nil {
		log.Error(err)
	}

	return bin_buf
}

func ChannelOpenReplySuccess(recipientChannel, senderChannel uint32) APF_CHANNEL_OPEN_CONFIRMATION_MESSAGE {
	log.Debug("sending APF_CHANNEL_OPEN_CONFIRMATION")

	message := APF_CHANNEL_OPEN_CONFIRMATION_MESSAGE{}
	message.MessageType = APF_CHANNEL_OPEN_CONFIRMATION
	message.RecipientChannel = recipientChannel
	message.SenderChannel = senderChannel
	message.InitialWindowSize = LME_RX_WINDOW_SIZE
	message.Reserved = 0xFFFFFFFF

	log.Tracef("%+v", message)

	return message
}

func ChannelOpenReplyFailure(recipientChannel, reason uint32) APF_CHANNEL_OPEN_FAILURE_MESSAGE {
	log.Debug("sending APF_CHANNEL_OPEN_FAILURE")

	message := APF_CHANNEL_OPEN_FAILURE_MESSAGE{}
	message.MessageType = APF_CHANNEL_OPEN_FAILURE
	message.RecipientChannel = recipientChannel
	message.ReasonCode = reason
	message.Reserved = 0x00000000
	message.Reserved2 = 0x00000000

	return message
}

func ChannelClose(recipientChannel uint32) APF_CHANNEL_CLOSE_MESSAGE {
	log.Debug("sending APF_CHANNEL_CLOSE_MESSAGE")

	message := APF_CHANNEL_CLOSE_MESSAGE{}
	message.MessageType = APF_CHANNEL_CLOSE
	message.RecipientChannel = recipientChannel

	return message
}

func ChannelData(recipientChannel uint32, buffer []byte) APF_CHANNEL_DATA_MESSAGE {
	log.Debug("sending APF_CHANNEL_DATA_MESSAGE")

	message := APF_CHANNEL_DATA_MESSAGE{}
	message.MessageType = APF_CHANNEL_DATA
	message.RecipientChannel = recipientChannel
	message.DataLength = uint32(len(buffer))
	message.Data = buffer

	return message
}

func ChannelWindowAdjust(recipientChannel, l uint32) APF_CHANNEL_WINDOW_ADJUST_MESSAGE {
	log.Debug("sending APF_CHANNEL_WINDOW_ADJUST_MESSAGE")

	message := APF_CHANNEL_WINDOW_ADJUST_MESSAGE{}
	message.MessageType = APF_CHANNEL_WINDOW_ADJUST
	message.RecipientChannel = recipientChannel
	message.BytesToAdd = l

	return message
}

// BuildChannelDataBytes serializes APF_CHANNEL_DATA for sending over the wire.
func BuildChannelDataBytes(recipientChannel uint32, data []byte) []byte {
	log.Debug("building APF_CHANNEL_DATA bytes")

	buf := make([]byte, 9+len(data))
	buf[0] = APF_CHANNEL_DATA
	binary.BigEndian.PutUint32(buf[1:5], recipientChannel)
	binary.BigEndian.PutUint32(buf[5:9], uint32(len(data)))
	copy(buf[9:], data)

	return buf
}

// BuildChannelCloseBytes serializes APF_CHANNEL_CLOSE for sending over the wire.
func BuildChannelCloseBytes(recipientChannel uint32) []byte {
	log.Debug("building APF_CHANNEL_CLOSE bytes")

	buf := make([]byte, 5)
	buf[0] = APF_CHANNEL_CLOSE
	binary.BigEndian.PutUint32(buf[1:5], recipientChannel)

	return buf
}

// BuildChannelWindowAdjustBytes serializes APF_CHANNEL_WINDOW_ADJUST for sending over the wire.
func BuildChannelWindowAdjustBytes(recipientChannel, bytesToAdd uint32) []byte {
	log.Debug("building APF_CHANNEL_WINDOW_ADJUST bytes")

	buf := make([]byte, 9)
	buf[0] = APF_CHANNEL_WINDOW_ADJUST
	binary.BigEndian.PutUint32(buf[1:5], recipientChannel)
	binary.BigEndian.PutUint32(buf[5:9], bytesToAdd)

	return buf
}
