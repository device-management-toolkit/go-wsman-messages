/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package auditlog facilitates communication with Intel® AMT devices to read the audit log records
package auditlog

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewAuditLogWithClient instantiates a new Audit Log service.
func NewAuditLogWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTAuditLog, client),
	}
}

// Get retrieves the representation of the instance.
func (service Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Get(nil),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (service Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Enumerate(),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (service Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Pull(enumerationContext),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// ReadRecords returns a list of consecutive audit log records in chronological order:
// The first record in the returned array is the oldest record stored in the log.
// startIndex Identifies the position of the first record to retrieve. An index of 1 indicates the first record in the log.
func (service Service) ReadRecords(startIndex int) (response Response, err error) {
	if startIndex < 1 {
		startIndex = 0
	}

	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuditLog, ReadRecords), AMTAuditLog, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(ReadRecords), AMTAuditLog, &ReadRecordsInput{StartIndex: startIndex})

	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	response.Body.DecodedRecordsResponse = convertToAuditLogResult(response.Body.ReadRecordsResponse.EventRecords)

	return
}
