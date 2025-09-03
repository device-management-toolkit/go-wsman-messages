/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package messagelog facilitiates communication with Intel® AMT devices to provide access to the AMT event log.
// Client should read the record time parameter in order to determine which record is newer.
// In most implementations, log entries are stored backwards, i.e. the newest record is the first record and the oldest record is the last (index based).
//
// Records have no header and the record data is combined of 21 binary bytes which could be read as:
//
//	struct {
//	UINT32 TimeStamp; // little endian
//	UINT8 DeviceAddress;
//	UINT8 EventSensorType;
//	UINT8 EventType;
//	UINT8 EventOffset;
//	UINT8 EventSourceType;
//	UINT8 EventSeverity;
//	UINT8 SensorNumber;
//	UINT8 Entity;
//	UINT8 EntityInstance;
//	UINT8 EventData[8];
//	} EVENT_DATA;
package messagelog

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Service struct {
	base.WSManService[Response]
}

// NewMessageLogWithClient instantiates a new MessageLog.
func NewMessageLogWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base.NewService[Response](wsmanMessageCreator, AMTMessageLog, client),
	}
}

// GetRecords retrieves multiple records from event log.
// The IterationIdentifier input parameter is a numeric value (starting at 1) which is the position of the first record in the log that should be extracted.
// MaxReadRecords is set to 390.  If NoMoreRecords returns false, call this again setting the identifier to the start of the next IterationIdentifier.
func (messageLog Service) GetRecords(identifier, maxReadRecords int) (response Response, err error) {
	if identifier < 1 {
		identifier = 1
	}

	if maxReadRecords < 1 {
		maxReadRecords = DefaultRecords
	} else if maxReadRecords > MaxAMTRecords {
		maxReadRecords = MaxAMTRecords
	}

	header := messageLog.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTMessageLog, GetRecords), AMTMessageLog, nil, "", "")
	body := messageLog.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetRecords), AMTMessageLog, &GetRecords_INPUT{
		IterationIdentifier: identifier,
		MaxReadRecords:      maxReadRecords,
	})

	response = Response{
		Message: &client.Message{
			XMLInput: messageLog.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = messageLog.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	response.Body.GetRecordsResponse.RawEventData, err = parseEventLogResult(response.Body.GetRecordsResponse.RecordArray)
	if err != nil {
		return response, err
	}

	response.Body.GetRecordsResponse.RefinedEventData = decodeEventRecord(response.Body.GetRecordsResponse.RawEventData)

	return response, err
}

// Requests that an iteration of the MessageLog be established and that the iterator be set to the first entry in the Log. An identifier for the iterator is returned as an output parameter of the method. Regarding iteration, you have 2 choices: 1) Embed iteration data in the method call, and allow implementations to track/ store this data manually; or, 2) Iterate using a separate object (for example, class ActiveIterator) as an iteration agent. The first approach is used here for interoperability. The second requires an instance of the Iterator object for EACH iteration in progress. 2's functionality could be implemented underneath 1.
//
// Product Specific Usage: In current implementation this method doesn't have any affect. In order to get the events from the log user should just call GetRecord or GetRecords.
func (messageLog Service) PositionToFirstRecord() (response Response, err error) {
	header := messageLog.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTMessageLog, PositionToFirstRecord), AMTMessageLog, nil, "", "")
	body := messageLog.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(PositionToFirstRecord), AMTMessageLog, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: messageLog.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = messageLog.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
