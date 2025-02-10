/*********************************************************************
 * Copyright (c) Intel Corporation 2022
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package apf

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func Process(data []byte, session *Session) bytes.Buffer {
	var bin_buf bytes.Buffer

	var dataToSend interface{}

	switch data[0] {
	case APF_KEEPALIVE_REQUEST:
		log.Debug("received APF_KEEPALIVE_REQUEST")

		dataToSend = ProcessKeepAliveRequest(data, session)
	case APF_KEEPALIVE_REPLY:
		log.Debug("received APF_KEEPALIVE_REPLY")

		// 	dataToSend = ProcessKeepAliveReply(data, session)
	case APF_KEEPALIVE_OPTIONS_REPLY:
		log.Debug("received APF_KEEPALIVE_OPTIONS_REQUEST")

	// 	dataToSend = ProcessKeepAliveOptionsReply(data, session)
	case APF_GLOBAL_REQUEST: // 80
		log.Debug("received APF_GLOBAL_REQUEST")

		if ValidateGlobalRequest(data) {
			dataToSend = ProcessGlobalRequest(data)
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
			dataToSend = ProcessProtocolVersion(data)
		}
	case APF_USERAUTH_REQUEST: // 50
		log.Debug("received APF_USERAUTH_REQUEST")
		dataToSend = ProcessUserAuthRequest(data, session)
	default:
	}

	if dataToSend != nil {
		err := binary.Write(&bin_buf, binary.BigEndian, dataToSend)
		if err != nil {
			log.Error(err)
		}
	}
	fmt.Printf("bin_buf: %x\n", bin_buf.Bytes())
	return bin_buf
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
	// Update session state if necessary
}

func ProcessKeepAliveOptionsReply(data []byte, session *Session) {
	if len(data) < 9 {
		log.Warn("APF_KEEPALIVE_OPTIONS_REPLY message too short")
		return
	}
	keepaliveInterval := binary.BigEndian.Uint32(data[1:5])
	timeout := binary.BigEndian.Uint32(data[5:9])
	log.Debugf("KEEPALIVE_OPTIONS_REPLY, Keepalive Interval=%d Timeout=%d", keepaliveInterval, timeout)
	// Update session state or configurations as needed
}

func ProcessUserAuthRequest(data []byte, session *Session) interface{} {
	log.Debug("received APF_USERAUTH_REQUEST")

	dataBuffer := bytes.NewReader(data)

	var messageType byte
	err := binary.Read(dataBuffer, binary.BigEndian, &messageType)
	if err != nil {
		log.Error(err)
		return nil
	}

	// Read username length
	var usernameLen uint32
	err = binary.Read(dataBuffer, binary.BigEndian, &usernameLen)
	if err != nil {
		log.Error(err)
		return nil
	}

	if usernameLen > 2048 || uint32(dataBuffer.Len()) < usernameLen {
		log.Error("Invalid username length")
		return nil
	}

	usernameBytes := make([]byte, usernameLen)
	n, err := dataBuffer.Read(usernameBytes)
	if err != nil || n != int(usernameLen) {
		log.Error("Failed to read username")
		return nil
	}

	username := string(usernameBytes)

	// Read serviceName length
	var serviceNameLen uint32
	err = binary.Read(dataBuffer, binary.BigEndian, &serviceNameLen)
	if err != nil {
		log.Error(err)
		return nil
	}

	if serviceNameLen > 2048 || uint32(dataBuffer.Len()) < serviceNameLen {
		log.Error("Invalid serviceName length")
		return nil
	}

	serviceNameBytes := make([]byte, serviceNameLen)
	n, err = dataBuffer.Read(serviceNameBytes)
	if err != nil || n != int(serviceNameLen) {
		log.Error("Failed to read serviceName")
		return nil
	}

	serviceName := string(serviceNameBytes)

	// Read methodName length
	var methodNameLen uint32
	err = binary.Read(dataBuffer, binary.BigEndian, &methodNameLen)
	if err != nil {
		log.Error(err)
		return nil
	}

	if methodNameLen > 2048 || uint32(dataBuffer.Len()) < methodNameLen {
		log.Error("Invalid methodName length")
		return nil
	}

	methodNameBytes := make([]byte, methodNameLen)
	n, err = dataBuffer.Read(methodNameBytes)
	if err != nil || n != int(methodNameLen) {
		log.Error("Failed to read methodName")
		return nil
	}

	methodName := string(methodNameBytes)

	var password string
	if methodName == "password" {
		if dataBuffer.Len() < 1 {
			log.Error("Not enough data for password FALSE byte")
			return nil
		}
		// Read boolean FALSE
		var passwordFalse byte
		err = binary.Read(dataBuffer, binary.BigEndian, &passwordFalse)
		if err != nil {
			log.Error(err)
			return nil
		}

		if passwordFalse != 0 {
			log.Error("passwordFalse is not zero")
			return nil
		}

		// Read password length
		var passwordLen uint32
		err = binary.Read(dataBuffer, binary.BigEndian, &passwordLen)
		if err != nil {
			log.Error(err)
			return nil
		}

		if passwordLen > 2048 || uint32(dataBuffer.Len()) < passwordLen {
			log.Error("Invalid password length")
			return nil
		}

		passwordBytes := make([]byte, passwordLen)
		n, err = dataBuffer.Read(passwordBytes)
		if err != nil || n != int(passwordLen) {
			log.Error("Failed to read password")
			return nil
		}

		password = string(passwordBytes)
		fmt.Print(password)
	} else {
		// Unsupported method
		log.Warn("Unsupported authentication method: ", methodName)
		// Return failure
		// failureMessage := &APF_USERAUTH_FAILURE_MESSAGE{
		// 	MessageType:                          APF_USERAUTH_FAILURE,
		// 	AuthenticationsThatCanContinueLength: uint32(len("password")),
		// 	AuthenticationsThatCanContinue:       []byte("password"),
		// 	PartialSuccess:                       0,
		// }
		// return failureMessage
		return nil
	}

	log.Debugf("usernameLen=%d serviceNameLen=%d methodNameLen=%d", usernameLen, serviceNameLen, methodNameLen)
	log.Debugf("username=%s serviceName=%s methodName=%s", username, serviceName, methodName)

	// Now authenticate the user
	authenticated := true //session.AuthenticateUser(username, password)

	if authenticated {
		// Return success message
		message := &APF_USERAUTH_SUCCESS_MESSAGE{
			MessageType: APF_USERAUTH_SUCCESS,
		}
		return message
	} else {
		// Return failure message
		// failureMessage := &APF_USERAUTH_FAILURE_MESSAGE{
		// 	MessageType:                          APF_USERAUTH_FAILURE,
		// 	AuthenticationsThatCanContinueLength: uint32(len("password")),
		// 	AuthenticationsThatCanContinue:       []byte("password"),
		// 	PartialSuccess:                       0,
		// }
		// return failureMessage
		return nil
	}
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

func ProcessGlobalRequest(data []byte) interface{} {
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

	if int(genericHeader.StringLength) > 0 {
		stringBuffer := make([]byte, genericHeader.StringLength)
		tcpForwardRequest := APF_TCP_FORWARD_REQUEST{}

		err = binary.Read(dataBuffer, binary.BigEndian, &stringBuffer)
		if err != nil {
			log.Error(err)
		}

		genericHeader.String = string(stringBuffer[:int(genericHeader.StringLength)])

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
		}

		err = binary.Read(dataBuffer, binary.BigEndian, &tcpForwardRequest.Port)
		if err != nil {
			log.Error(err)
		}

		log.Tracef("%+v", genericHeader)
		log.Tracef("%+v", tcpForwardRequest)

		if genericHeader.String == APF_GLOBAL_REQUEST_STR_TCP_FORWARD_REQUEST {
			//if tcpForwardRequest.Port == 16992 || tcpForwardRequest.Port == 16993 {
			reply = TcpForwardReplySuccess(tcpForwardRequest.Port)
			if tcpForwardRequest.Port == 5900 {

			}
			// } else {
			// 	reply = APF_REQUEST_FAILURE
			// }
		} else if genericHeader.String == APF_GLOBAL_REQUEST_STR_TCP_FORWARD_CANCEL_REQUEST {
			reply = APF_REQUEST_SUCCESS
		}
	}

	return reply
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
		if message.ServiceName == "pfwd@amt.intel.com" {
			service = 1
		} else if message.ServiceName == "auth@amt.intel.com" {
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

	log.Tracef("%+v", message)
	version := ProtocolVersion(message.MajorVersion, message.MinorVersion, message.TriggerReason)

	return version
}

// Send the AFP service accept message to the MEI.
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
	log.Debug("sending APF_PROTOCOL_VERSION_MESSAGE")

	protVersion := APF_PROTOCOL_VERSION_MESSAGE{}
	protVersion.MessageType = APF_PROTOCOLVERSION
	protVersion.MajorVersion = majorversion
	protVersion.MinorVersion = minorversion
	protVersion.TriggerReason = triggerreason

	log.Tracef("%+v", protVersion)

	return protVersion
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
