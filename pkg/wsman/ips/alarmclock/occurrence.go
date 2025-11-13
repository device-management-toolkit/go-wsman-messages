/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package alarmclock facilitates communication with Intel® AMT devices to represent a single alarm clock setting.
package alarmclock

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Occurrence struct {
	base.WSManService[Response]
}

// NewAlarmClockOccurrence returns a new instance of the AlarmClockOccurrence struct.
func NewAlarmClockOccurrenceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Occurrence {
	return Occurrence{
		base.NewService[Response](wsmanMessageCreator, IPSAlarmClockOccurrence, client),
	}
}

// Delete removes a the specified instance.
func (occurrence Occurrence) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "InstanceID", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: occurrence.Base.Delete(selector),
		},
	}

	err = occurrence.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
