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
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Service struct {
	base.WSManService[Response]
}

// NewAuditLogWithClient instantiates a new Audit Log service.
func NewAuditLogWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base.NewService[Response](wsmanMessageCreator, AMTAuditLog, client),
	}
}

// ReadRecords returns a list of consecutive audit log records in chronological order:
// The first record in the returned array is the oldest record stored in the log.
// startIndex Identifies the position of the first record to retrieve. An index of 1 indicates the first record in the log.
func (service Service) ReadRecords(startIndex int) (response Response, err error) {
	if startIndex < 1 {
		startIndex = 0
	}

	header := service.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuditLog, ReadRecords), AMTAuditLog, nil, "", "")
	body := service.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(ReadRecords), AMTAuditLog, &ReadRecordsInput{StartIndex: startIndex})

	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = service.Base.Execute(response.Message)
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
